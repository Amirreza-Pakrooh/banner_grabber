package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

func CreateIPRange(startIP string, endIP string) []string {
	var startIpParts []string = strings.Split(startIP, ".")
	var endIpParts []string = strings.Split(endIP, ".")

	startYPart, _ := strconv.Atoi(startIpParts[2])
	endYPart, _ := strconv.Atoi(endIpParts[2])
	startXPart, _ := strconv.Atoi(startIpParts[3])
	endXPart, _ := strconv.Atoi(endIpParts[3])

	var ipRange []string = []string{}

	for i := startYPart; i <= endYPart; i++ {
		for j := startXPart; j <= endXPart; j++ {
			var IP string = fmt.Sprintf("%s.%s.%d.%d", startIpParts[0], startIpParts[1], i, j)
			ipRange = append(ipRange, IP)
		}
	}

	return ipRange
}

func ScanPort(IP string, port int, timeout time.Duration, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	var portAddress string = fmt.Sprintf("%s:%d", IP, port)
	var connection, error = net.DialTimeout("tcp", portAddress, timeout)

	if error != nil {
		fmt.Printf("[IP: %s: Port: %d] is close\n", IP, port)
		return
	}
	defer connection.Close()

	fmt.Printf("[IP: %s: Port: %d] is open\n", IP, port)

	if port == 21 || port == 25 || port == 587 {
		var banner string = BannerGrabbing(connection, port, timeout)
		fmt.Printf("[IP: %s: Port:%d] Bannar:\n%s\n", IP, port, banner)
	}
}

func BannerGrabbing(conection net.Conn, port int, timeout time.Duration) string {
	conection.SetReadDeadline(time.Now().Add(timeout))

	var buffer []byte = make([]byte, 4096)
	var banner string

	for {
		response, error := conection.Read(buffer)

		if error != nil {
			break
		}

		banner += string(buffer[:response])
	}

	if port == 21 {
		conection.Write([]byte("USER anonymous\r\n"))
	} else if port == 25 || port == 587 {
		conection.Write([]byte("EHLO example.com\r\n"))
	}

	for {
		response, error := conection.Read(buffer)

		if error != nil {
			break
		}

		banner += string(buffer[:response])
	}

	if banner == "" {
		return "No banner!!!"
	}

	return banner
}

func main() {
	var startIP string
	var endIP string
	var timeoutInt int

	var ports []int = []int{21, 22, 25, 80, 443, 587}

	fmt.Print("Enter start IP: ")
	fmt.Scan(&startIP)
	fmt.Print("Enter end IP: ")
	fmt.Scan(&endIP)
	fmt.Print("Enter timeout (seconds): ")
	fmt.Scan(&timeoutInt)

	var timeout time.Duration = time.Duration(timeoutInt) * time.Second

	var IPRange []string = CreateIPRange(startIP, endIP)

	var waitGroup sync.WaitGroup
	for i := 0; i < len(IPRange); i++ {
		for j := 0; j < len(ports); j++ {
			waitGroup.Add(1)
			go ScanPort(IPRange[i], ports[j], timeout, &waitGroup)
		}
	}

	waitGroup.Wait()

	fmt.Println("End!")

	for {

	}
}

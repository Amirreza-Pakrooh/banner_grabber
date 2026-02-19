Network Scanner & Banner Grabber (Golang)

A simple TCP network scanner written in Go that scans a range of IP addresses for open ports and performs banner grabbing on specific services.

📌 Project Overview

This project is a concurrent TCP port scanner that:

Accepts a start and end IP address

Scans predefined ports on all IPs within the range

Detects whether ports are open or closed

Performs Banner Grabbing on FTP and SMTP ports

Uses Goroutines for concurrent scanning

⚙️ Technologies Used

Go (Golang)

net package for TCP connections

sync.WaitGroup for concurrency control

time for timeout handling

🧠 Core Concepts
🔹 Ports

Ports are communication endpoints in networking, numbered from 0 to 65535.

Ports scanned in this project:

21 → FTP

22 → SSH

25 → SMTP

80 → HTTP

443 → HTTPS

587 → SMTP (Submission)

🔹 TCP Connection

The program uses:

net.DialTimeout("tcp", address, timeout)


If the connection is not established within the specified timeout, it is aborted.

🔹 Banner Grabbing

Some services send identifying information immediately after connection, such as:

Server software name

Version

Supported capabilities

This project performs banner grabbing on:

FTP (21) → sends USER anonymous

SMTP (25, 587) → sends EHLO example.com

🏗️ Function Structure
1️⃣ CreateIPRange(startIP, endIP)

Splits start and end IP addresses

Generates all IPs within the specified range

Returns a slice of strings

2️⃣ ScanPort(IP, port, timeout, wg)

Attempts TCP connection

Prints whether the port is open or closed

Calls BannerGrabbing if applicable

3️⃣ BannerGrabbing(connection, port, timeout)

Sets read deadline

Reads initial server response

Sends protocol-specific command

Reads additional response

Returns the collected banner

🚀 How to Run
1️⃣ Clone the repository
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name

2️⃣ Run the program
go run main.go

3️⃣ Provide input
Enter start IP:
Enter end IP:
Enter timeout (seconds):

🔄 Concurrency Model

The scanner uses Goroutines to scan multiple ports simultaneously:

go ScanPort(...)


Concurrency is controlled using:

sync.WaitGroup


Add(1) before launching each Goroutine

Done() inside each Goroutine

Wait() to block until all scans complete

📊 Testing & Verification

Results can be compared with professional network scanning tools such as:

Nmap

⚠️ Disclaimer

This tool is intended for educational purposes and authorized network testing only.
Unauthorized scanning of networks may violate laws or policies.

👨‍💻 Author

Developed as a networking and Go concurrency practice project.

# Network Scanner & Banner Grabber (Golang)

A concurrent TCP port scanner written in Go that scans a range of IP addresses for open ports and performs banner grabbing on selected services.

---

## 📌 Features

- Scan a range of IPv4 addresses
- Detect open and closed TCP ports
- Perform banner grabbing on FTP and SMTP services
- Configurable timeout
- Concurrent scanning using Goroutines

---

## ⚙️ Technologies Used

- Go (Golang)
- net package
- sync.WaitGroup
- time package

---

## 🧠 Scanned Ports

The following ports are scanned by default:

- 21  → FTP  
- 22  → SSH  
- 25  → SMTP  
- 80  → HTTP  
- 443 → HTTPS  
- 587 → SMTP Submission  

You can modify the `ports` slice in `main.go` to scan additional ports.

---

## 🏗 Project Structure

### CreateIPRange

Generates a list of IP addresses between the provided start and end IP.

### ScanPort

- Attempts a TCP connection using `net.DialTimeout`
- Detects whether the port is open or closed
- Calls banner grabbing when applicable

### BannerGrabbing

- Reads the initial server response
- Sends protocol-specific commands:
  - FTP → `USER anonymous`
  - SMTP → `EHLO example.com`
- Collects and returns the banner

---

## 🚀 How to Run

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
```

### 2️⃣ Run the Program

```bash
go run main.go
```

### 3️⃣ Provide Input

```
Enter start IP:
Enter end IP:
Enter timeout (seconds):
```

---

## 🔄 Concurrency

This project uses Goroutines for parallel port scanning:

- Each port scan runs in a separate Goroutine  
- `sync.WaitGroup` ensures all scans complete before program exit  

---

## 📊 Testing

Results can be compared with tools like Nmap for validation.

---

## ⚠️ Usage Notice

This project is intended for educational and authorized testing purposes only.

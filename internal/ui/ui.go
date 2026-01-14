package ui

import (
	"fmt"
	"os"
	"time"
)

var startTime time.Time

func Banner() {
	fmt.Println("┌─────────────────────────────────────────────┐")
	fmt.Println("│ Subhunt v0.1.0                              │")
	fmt.Println("│ Active Subdomain Enumeration (DoH)          │")
	fmt.Println("│ Author: Vikash Choudhary                    │")
	fmt.Println("└─────────────────────────────────────────────┘")
}

func Info(msg string) {
	fmt.Fprintf(os.Stderr, "[*] %s\n", msg)
}

func Found(sub string) {
	fmt.Fprintf(os.Stderr, "[+] %s\n", sub)
}

func Warn(msg string) {
	fmt.Fprintf(os.Stderr, "[!] %s\n", msg)
}

func Done(msg string) {
	fmt.Fprintf(os.Stderr, "[✓] %s\n", msg)
}

func StartTimer() {
	startTime = time.Now()
}

func Duration() string {
	return time.Since(startTime).Truncate(time.Second).String()
}

package ui

import (
	"fmt"
	"os"
	"time"
)

var start time.Time

// ANSI colors
const (
	reset  = "\033[0m"
	bold   = "\033[1m"

	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
	gray   = "\033[90m"
)

func StartTimer() {
	start = time.Now()
}

func Duration() string {
	if start.IsZero() {
		return "0s"
	}
	return time.Since(start).Round(time.Millisecond).String()
}

func Banner() {
	fmt.Fprintln(os.Stderr, cyan+"┌─────────────────────────────────────────────┐"+reset)
	fmt.Fprintln(os.Stderr, cyan+"│ "+bold+"Subhunt v1.1.0"+reset+cyan+"                              │"+reset)
	fmt.Fprintln(os.Stderr, cyan+"│ "+gray+"Active Subdomain Enumeration (DoH)"+reset+cyan+"          │"+reset)
	fmt.Fprintln(os.Stderr, cyan+"│ "+gray+"Author: Vikash Choudhary"+reset+cyan+"                    │"+reset)
	fmt.Fprintln(os.Stderr, cyan+"└─────────────────────────────────────────────┘"+reset)
}

func Info(msg string) {
	fmt.Fprintf(os.Stderr, blue+"[*] "+reset+"%s\n", msg)
}

func Warn(msg string) {
	fmt.Fprintf(os.Stderr, yellow+"[!] "+reset+"%s\n", msg)
}

func Error(msg string) {
	fmt.Fprintf(os.Stderr, red+"[x] "+reset+"%s\n", msg)
}

func Done(msg string) {
	fmt.Fprintf(os.Stderr, green+"[✓] "+reset+"%s\n", msg)
}

func Found(sub string) {
	// Results MUST always remain on stdout
	fmt.Fprintf(os.Stdout, green+"[+] "+reset+"%s\n", sub)
}

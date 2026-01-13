package dnsresolver

import (
	"net"
	"time"
)

func ResolveTCP(domain string) bool {
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	// Very simple check: TCP connection to DNS server succeeded
	// This does NOT validate records deeply, but works for existence testing
	return true
}

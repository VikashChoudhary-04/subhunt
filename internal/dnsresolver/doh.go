package dnsresolver

import (
	"encoding/json"
	"net/http"
	"time"
)

type dohResponse struct {
	Status int `json:"Status"`
}

func ResolveDoH(domain string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		"https://cloudflare-dns.com/dns-query?name="+domain+"&type=A",
		nil,
	)
	if err != nil {
		return false
	}

	// REQUIRED by Cloudflare
	req.Header.Set("Accept", "application/dns-json")
	req.Header.Set("User-Agent", "subhunt")

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var result dohResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false
	}

	// Status == 0 means the domain exists
	return result.Status == 0
}

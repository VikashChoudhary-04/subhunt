package dnsresolver

import (
	"encoding/json"
	"net/http"
	"time"
)

type dohResponse struct {
	Status int `json:"Status"`
}

// List of DoH endpoints (failover order)
var dohEndpoints = []string{
	"https://cloudflare-dns.com/dns-query",
	"https://dns.google/resolve",
	"https://dns.quad9.net/dns-query",
}

// ResolveDoH tries all configured DoH providers and
// returns true as soon as one confirms the domain exists.
func ResolveDoH(domain string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, endpoint := range dohEndpoints {
		if resolveWithEndpoint(client, endpoint, domain) {
			return true
		}
	}

	return false
}

func resolveWithEndpoint(client *http.Client, endpoint, domain string) bool {
	req, err := http.NewRequest(
		"GET",
		endpoint+"?name="+domain+"&type=A",
		nil,
	)
	if err != nil {
		return false
	}

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

	return result.Status == 0
}

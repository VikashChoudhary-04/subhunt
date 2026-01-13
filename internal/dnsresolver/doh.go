package dnsresolver

import (
	"encoding/json"
	"net/http"
	"time"
)

type dohResponse struct {
	Answer []struct {
		Data string `json:"data"`
	} `json:"Answer"`
}

func ResolveDoH(domain string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		"https://dns.google/resolve?name="+domain+"&type=A",
		nil,
	)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var result dohResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false
	}

	return len(result.Answer) > 0
}

package passive

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type crtEntry struct {
	NameValue string `json:"name_value"`
}

func CRTSH(domain string) ([]string, error) {
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	// Force IPv4
	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, "tcp4", addr)
		},
	}

	client := &http.Client{
		Timeout:   20 * time.Second,
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entries []crtEntry
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, err
	}

	results := []string{}
	for _, e := range entries {
		names := strings.Split(e.NameValue, "\n")
		for _, n := range names {
			n = strings.TrimSpace(n)
			if strings.HasSuffix(n, domain) {
				results = append(results, n)
			}
		}
	}

	return results, nil
}

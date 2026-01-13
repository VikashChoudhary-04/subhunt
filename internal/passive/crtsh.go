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

	dialer := &net.Dialer{
		Timeout:   8 * time.Second,
		KeepAlive: 8 * time.Second,
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, "tcp4", addr)
		},
	}

	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: transport,
	}

	var lastErr error

	for attempt := 1; attempt <= 3; attempt++ {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "subhunt/1.0")

		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("crt.sh returned %d", resp.StatusCode)
			resp.Body.Close()
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
			continue
		}

		var entries []crtEntry
		if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		seen := make(map[string]struct{})
		results := []string{}

		for _, e := range entries {
			for _, n := range strings.Split(e.NameValue, "\n") {
				n = strings.ToLower(strings.TrimSpace(n))

				// remove wildcard
				n = strings.TrimPrefix(n, "*.")

				// must end with .domain (real subdomain)
				if !strings.HasSuffix(n, "."+domain) {
					continue
				}

				// exclude root domain itself
				if n == domain {
					continue
				}

				// avoid duplicates
				if _, ok := seen[n]; ok {
					continue
				}
				seen[n] = struct{}{}

				results = append(results, n)
			}
		}

		return results, nil
	}

	return nil, lastErr
}

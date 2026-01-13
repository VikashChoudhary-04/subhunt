package passive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type crtEntry struct {
	NameValue string `json:"name_value"`
}

func CRTSH(domain string) ([]string, error) {
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	resp, err := http.Get(url)
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

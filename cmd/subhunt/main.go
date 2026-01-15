package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/VikashChoudhary-04/subhunt/internal/bruteforce"
	"github.com/VikashChoudhary-04/subhunt/internal/ui"
)

func main() {
	domain := flag.String("d", "", "Target domain (example.com)")
	wordlist := flag.String("bruteforce", "", "Path to wordlist")
	threads := flag.Int("threads", 10, "Number of concurrent workers")
	quiet := flag.Bool("quiet", false, "Show only results")
	jsonOut := flag.Bool("json", false, "Output results as JSON")

	flag.Parse()

	if *domain == "" || *wordlist == "" {
		fmt.Fprintf(os.Stderr,
			"Usage: subhunt -d example.com --bruteforce wordlist.txt [--threads 50]\n")
		os.Exit(1)
	}

	// JSON implies quiet mode
	if *jsonOut {
		*quiet = true
	}

	if !*quiet {
		ui.Banner()
		ui.StartTimer()

		ui.Info(fmt.Sprintf("Target      : %s", *domain))
		ui.Info(fmt.Sprintf("Wordlist    : %s", *wordlist))
		ui.Info(fmt.Sprintf("Threads     : %d", *threads))
		ui.Info("Resolver    : DNS over HTTPS (DoH)")
		ui.Info("Mode        : Active Bruteforce")
		fmt.Fprintln(os.Stderr, "------------------------------------------------")
	} else {
		// Still start timer for JSON duration
		ui.StartTimer()
	}

	results, stats := bruteforce.Brute(
		*domain,
		*wordlist,
		*threads,
		*quiet,
	)

	// JSON OUTPUT MODE
	if *jsonOut {
		_ = json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
			"domain":       *domain,
			"found":        results,
			"tested":       stats.Tested,
			"found_count":  stats.Found,
			"duration":     ui.Duration(),
		})

		if stats.Found > 0 {
			os.Exit(0)
		}
		os.Exit(1)
	}

	// NORMAL OUTPUT MODE
	for _, sub := range results {
		ui.Found(sub)
	}

	if !*quiet {
		fmt.Fprintln(os.Stderr)
		ui.Done("Scan Finished")

		fmt.Fprintf(os.Stderr, `
Target        : %s
Total Tested  : %d
Total Found   : %d
Duration      : %s
Resolver      : DoH
------------------------------------------------
`,
			*domain,
			stats.Tested,
			stats.Found,
			ui.Duration(),
		)
	}

	if stats.Found > 0 {
		os.Exit(0)
	}
	os.Exit(1)
}

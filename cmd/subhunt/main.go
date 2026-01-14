package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VikashChoudhary-04/subhunt/internal/bruteforce"
	"github.com/VikashChoudhary-04/subhunt/internal/ui"
)

func main() {
	// -------------------------------
	// CLI FLAGS
	// -------------------------------
	domain := flag.String("d", "", "Target domain (example.com)")
	wordlist := flag.String("bruteforce", "", "Path to wordlist")
	threads := flag.Int("threads", 10, "Number of concurrent workers")
	quiet := flag.Bool("quiet", false, "Show only results")
	verbose := flag.Bool("verbose", false, "Verbose output")

	flag.Parse()

	// -------------------------------
	// VALIDATION
	// -------------------------------
	if *domain == "" || *wordlist == "" {
		fmt.Fprintf(os.Stderr,
			"Usage: subhunt -d example.com --bruteforce wordlist.txt [--threads 50]\n")
		os.Exit(1)
	}

	// -------------------------------
	// UI STARTUP
	// -------------------------------
	if !*quiet {
		ui.Banner()
		ui.StartTimer()

		ui.Info(fmt.Sprintf("Target      : %s", *domain))
		ui.Info(fmt.Sprintf("Wordlist    : %s", *wordlist))
		ui.Info(fmt.Sprintf("Threads     : %d", *threads))
		ui.Info("Resolver    : DNS over HTTPS (Cloudflare)")
		ui.Info("Mode        : Active Bruteforce")
		fmt.Fprintln(os.Stderr, "------------------------------------------------")
	}

	if *verbose {
		ui.Info("Verbose mode enabled")
	}

	// -------------------------------
	// RUN BRUTEFORCE
	// -------------------------------
	results, stats := bruteforce.Brute(
		*domain,
		*wordlist,
		*threads,
		*quiet,
	)

	// -------------------------------
	// OUTPUT RESULTS (STDOUT)
	// -------------------------------
	for _, sub := range results {
		ui.Found(sub)
	}

	// -------------------------------
	// FINAL SUMMARY
	// -------------------------------
	if !*quiet {
		fmt.Fprintln(os.Stderr)
		ui.Done("Scan Finished")

		fmt.Fprintf(os.Stderr, `
Target        : %s
Total Tested  : %d
Total Found   : %d
Duration      : %s
Resolver      : DoH (Cloudflare)
------------------------------------------------
`,
			*domain,
			stats.Tested,
			stats.Found,
			ui.Duration(),
		)
	}

	// -------------------------------
	// EXIT CODE
	// -------------------------------
	if stats.Found > 0 {
		os.Exit(0)
	}
	os.Exit(1)
}

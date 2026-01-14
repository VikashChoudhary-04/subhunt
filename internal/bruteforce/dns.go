package bruteforce

import (
	"bufio"
	"fmt"
	"github.com/VikashChoudhary-04/subhunt/internal/dnsresolver"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

func Brute(domain, wordlist string, workers int) []string {
	file, err := os.Open(wordlist)
	if err != nil {
		return nil
	}
	defer file.Close()

	if workers < 1 {
		workers = 10
	}

	jobs := make(chan string)
	results := make(chan string)

	var tested uint64
	var wg sync.WaitGroup
	done := make(chan struct{})

	// 🔵 LIVE PROGRESS DISPLAY (stderr)
	rate := atomic.LoadUint64(&tested) / uint64(time.Since(start).Seconds())
	fmt.Fprintf(os.Stderr,
		"\r[RUNNING] Tested: %d | Found: %d | Rate: %d/s",
		atomic.LoadUint64(&tested),
		atomic.LoadUint64(&found),
		rate,
	)


	// 🔵 Worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sub := range jobs {
				if dnsresolver.ResolveDoH(sub) {
				// ✅ MATCH FOUND — send immediately
					results <- sub
				}

				atomic.AddUint64(&tested, 1)
			}
		}()
	}

	// 🔵 Feed jobs
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text() + "." + domain
		}
		close(jobs)
	}()

	// 🔵 Close channels
	go func() {
		wg.Wait()
		close(results)
		close(done)
	}()

	// 🔵 Print results immediately (stdout)
	var found []string
	for r := range results {
		fmt.Println(r) // 👈 PRINTS AS SOON AS FOUND
		found = append(found, r)
	}

	return found
	var found uint64
	atomic.AddUint64(&found, 1)
	results <- sub
	
}
ui.Done("Scan Finished")

fmt.Fprintf(os.Stderr, `
Target        : %s
Total Tested  : %d
Total Found   : %d
Duration      : %s
Resolver      : DoH (Cloudflare)
------------------------------------------------
`,
domain,
tested,
found,
ui.Duration(),
)

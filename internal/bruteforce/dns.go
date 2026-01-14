package bruteforce

import (
	"bufio"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/VikashChoudhary-04/subhunt/internal/dnsresolver"
)

type Stats struct {
	Tested uint64
	Found  uint64
}

func Brute(domain, wordlist string, workers int, quiet bool) ([]string, Stats) {
	file, err := os.Open(wordlist)
	if err != nil {
		return nil, Stats{}
	}
	defer file.Close()

	if workers < 1 {
		workers = 10
	}

	jobs := make(chan string)
	results := make(chan string)

	var stats Stats
	var wg sync.WaitGroup
	done := make(chan struct{})
	start := time.Now()

	// ------------------------------------------------
	// STATUS LINE (ONLY goroutine that touches stderr)
	// ------------------------------------------------
	if !quiet {
		ticker := time.NewTicker(500 * time.Millisecond)

		go func() {
			for {
				select {
				case <-ticker.C:
					elapsed := time.Since(start).Seconds()

					var rate uint64
					if elapsed >= 1 {
						rate = atomic.LoadUint64(&stats.Tested) / uint64(elapsed)
					}

					os.Stderr.WriteString(
						"\r\033[K[RUNNING] Tested: " +
							strconv.FormatUint(atomic.LoadUint64(&stats.Tested), 10) +
							" | Found: " +
							strconv.FormatUint(atomic.LoadUint64(&stats.Found), 10) +
							" | Rate: " +
							strconv.FormatUint(rate, 10) + "/s",
					)

				case <-done:
					ticker.Stop()
					os.Stderr.WriteString("\n")
					return
				}
			}
		}()
	}

	// ------------------------------------------------
	// WORKER POOL (NO UI OUTPUT HERE)
	// ------------------------------------------------
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for sub := range jobs {
				if dnsresolver.ResolveDoH(sub) {
					atomic.AddUint64(&stats.Found, 1)
					results <- sub
				}
				atomic.AddUint64(&stats.Tested, 1)
			}
		}()
	}

	// ------------------------------------------------
	// FEED JOBS
	// ------------------------------------------------
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text() + "." + domain
		}
		close(jobs)
	}()

	// ------------------------------------------------
	// SHUTDOWN
	// ------------------------------------------------
	go func() {
		wg.Wait()
		close(results)
		close(done)
	}()

	// ------------------------------------------------
	// COLLECT RESULTS (stdout handled in main.go)
	// ------------------------------------------------
	var found []string
	for r := range results {
		found = append(found, r)
	}

	return found, stats
}

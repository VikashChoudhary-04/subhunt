package bruteforce

import (
	"bufio"
	"os"
	"sync"
	"strconv"
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

	// 🔵 LIVE STATUS LINE
	if !quiet {
		ticker := time.NewTicker(500 * time.Millisecond)
		go func() {
			for {
				select {
				case <-ticker.C:
					elapsed := time.Since(start).Seconds()
					rate := uint64(0)
					if elapsed > 0 {
						rate = atomic.LoadUint64(&stats.Tested) / uint64(elapsed)
					}
					os.Stderr.WriteString(
						"\r[RUNNING] Tested: " +
							format(atomic.LoadUint64(&stats.Tested)) +
							" | Found: " +
							format(atomic.LoadUint64(&stats.Found)) +
							" | Rate: " +
							format(rate) + "/s",
					)
				case <-done:
					ticker.Stop()
					return
				}
			}
		}()
	}

	// 🔵 WORKERS
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

	// 🔵 FEED JOBS
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text() + "." + domain
		}
		close(jobs)
	}()

	// 🔵 CLOSE
	go func() {
		wg.Wait()
		close(results)
		close(done)
	}()

	// 🔵 COLLECT RESULTS
	var found []string
	for r := range results {
		found = append(found, r)
	}

	return found, stats
}

// helper (keeps output readable)
func format(n uint64) string {
	return strconv.FormatUint(n, 10)
}

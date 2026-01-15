package bruteforce

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/VikashChoudhary-04/subhunt/internal/dnsresolver"
)

var dnsCache = struct {
	sync.RWMutex
	data map[string]bool
}{
	data: make(map[string]bool),
}

type Stats struct {
	Tested uint64
	Found  uint64
}

func Brute(domain, wordlist string, workers int, quiet bool) ([]string, Stats) {
	wildcard := hasWildcard(domain)

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

	// ------------------------------------------------
	// STATUS LINE (stderr only)
	// ------------------------------------------------
	if !quiet {
		ticker := time.NewTicker(500 * time.Millisecond)

		var lastTested uint64
		lastTime := time.Now()

		go func() {
			for {
				select {
				case <-ticker.C:
					now := time.Now()
					elapsed := now.Sub(lastTime).Seconds()

					current := atomic.LoadUint64(&stats.Tested)
					var rate uint64

					// Prevent spike on first tick
					if elapsed > 0 && lastTested > 0 {
						rate = uint64(float64(current-lastTested) / elapsed)
					}

					lastTested = current
					lastTime = now

					os.Stderr.WriteString(
						"\r\033[K[RUNNING] Tested: " +
							strconv.FormatUint(current, 10) +
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
	// WORKERS
	// ------------------------------------------------
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for sub := range jobs {
				dnsCache.RLock()
				cached, ok := dnsCache.data[sub]
				dnsCache.RUnlock()

				if ok {
					if cached {
						results <- sub
						atomic.AddUint64(&stats.Found, 1)
					}
					atomic.AddUint64(&stats.Tested, 1)
					continue
				}

				resolved := dnsresolver.ResolveDoH(sub)

				dnsCache.Lock()
				dnsCache.data[sub] = resolved
				dnsCache.Unlock()

				if resolved {
					results <- sub
					atomic.AddUint64(&stats.Found, 1)
				}

				atomic.AddUint64(&stats.Tested, 1)
			}
		}()
	}

	// ------------------------------------------------
	// FEED JOBS (scanner fix applied)
	// ------------------------------------------------
	go func() {
		scanner := bufio.NewScanner(file)

		buf := make([]byte, 0, 1024*1024)
		scanner.Buffer(buf, 1024*1024)

		for scanner.Scan() {
			jobs <- scanner.Text() + "." + domain
		}
		close(jobs)

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "[!] Wordlist read error: %v\n", err)
		}
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
	// COLLECT RESULTS
	// ------------------------------------------------
	var found []string
	for r := range results {
		found = append(found, r)
	}

	if wildcard && !quiet {
		fmt.Fprintln(os.Stderr, "[!] Wildcard DNS detected — results may include false positives")
	}

	return found, stats
}

func hasWildcard(domain string) bool {
	for i := 0; i < 3; i++ {
		test := fmt.Sprintf("random-%d.%s", rand.Int(), domain)
		if dnsresolver.ResolveDoH(test) {
			return true
		}
	}
	return false
}

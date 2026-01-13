package bruteforce

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
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

	// Worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sub := range jobs {
				net.LookupHost(sub)
				atomic.AddUint64(&tested, 1)
			}
		}()
	}

	// Result collector (separate lookup to avoid duplicate DNS)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sub := range jobs {
				if _, err := net.LookupHost(sub); err == nil {
					results <- sub
				}
			}
		}()
	}

	// Feed jobs
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text() + "." + domain

			// Print progress every 1000 tests
			if atomic.LoadUint64(&tested)%1000 == 0 && atomic.LoadUint64(&tested) != 0 {
				fmt.Fprintf(os.Stderr, "[+] Tested: %d\n", atomic.LoadUint64(&tested))
			}
		}
		close(jobs)
	}()

	// Close results when done
	go func() {
		wg.Wait()
		close(results)
	}()

	var found []string
	for r := range results {
		found = append(found, r)
	}

	fmt.Fprintf(os.Stderr, "[+] Total tested: %d\n", tested)
	return found
}

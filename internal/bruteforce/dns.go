package bruteforce

import (
	"bufio"
	"fmt"
	"net"
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

	// 🔴 LIVE PROGRESS DISPLAY (IMMEDIATE)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		fmt.Fprintf(os.Stderr, "[+] Tested: 0")
		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(os.Stderr, "\r[+] Tested: %d", atomic.LoadUint64(&tested))
			case <-done:
				ticker.Stop()
				fmt.Fprintf(os.Stderr, "\r[✓] Finished. Total tested: %d\n", atomic.LoadUint64(&tested))
				return
			}
		}
	}()

	// Worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sub := range jobs {
				_, _ = net.LookupHost(sub)
				atomic.AddUint64(&tested, 1)
			}
		}()
	}

	// Feed jobs
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text() + "." + domain
		}
		close(jobs)
	}()

	// Close everything
	go func() {
		wg.Wait()
		close(results)
		close(done)
	}()

	var found []string
	for r := range results {
		found = append(found, r)
	}

	return found
}

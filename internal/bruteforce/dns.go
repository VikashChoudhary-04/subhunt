package bruteforce

import (
	"bufio"
	"net"
	"os"
	"sync"
)

func Brute(domain, wordlist string) []string {
	file, err := os.Open(wordlist)
	if err != nil {
		return nil
	}
	defer file.Close()

	const workers = 50 // safe default
	jobs := make(chan string)
	results := make(chan string)

	var wg sync.WaitGroup

	// Worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sub := range jobs {
				_, err := net.LookupHost(sub)
				if err == nil {
					results <- sub
				}
			}
		}()
	}

	// Reader
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			sub := scanner.Text() + "." + domain
			jobs <- sub
		}
		close(jobs)
	}()

	// Closer
	go func() {
		wg.Wait()
		close(results)
	}()

	var found []string
	for r := range results {
		found = append(found, r)
	}

	return found
}

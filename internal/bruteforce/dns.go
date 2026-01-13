package bruteforce

import (
	"bufio"
	"net"
	"os"
)

func Brute(domain, wordlist string) []string {
	file, err := os.Open(wordlist)
	if err != nil {
		return nil
	}
	defer file.Close()

	var found []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sub := scanner.Text() + "." + domain
		_, err := net.LookupHost(sub)
		if err == nil {
			found = append(found, sub)
		}
	}

	return found
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VikashChoudhary-04/subhunt/internal/bruteforce"
	"github.com/VikashChoudhary-04/subhunt/internal/resolver"
	"github.com/VikashChoudhary-04/subhunt/internal/utils"
)

func main() {
	domain := flag.String("d", "", "Target domain")
	passiveEnum := flag.Bool("passive", false, "Enable passive enumeration")
	threads := flag.Int("threads", 50, "Number of concurrent DNS threads")
	bruteforceList := flag.String("bruteforce", "", "Wordlist for DNS bruteforce")
	resolve := flag.Bool("resolve", false, "Resolve subdomains")
	flag.Parse()

	if *domain == "" {
		fmt.Println("Usage: subhunt -d example.com [--passive] [--bruteforce wordlist] [--resolve]")
		os.Exit(1)
	}

	var subs []string

	if *passiveEnum {
	fmt.Println("[*] Passive enumeration enabled (no crt.sh)")
	}



	if *bruteforceList != "" {
		results := bruteforce.Brute(*domain, *bruteforceList, *threads)
		subs = append(subs, results...)
	}

	subs = utils.Dedupe(subs)

	// Only resolve if bruteforce was NOT used
	if *resolve && *bruteforceList == "" {
	subs = resolver.Resolve(subs)
	}


	for _, s := range subs {
		fmt.Println(s)
	}
}

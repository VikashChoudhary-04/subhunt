package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VikashChoudhary-04/subhunt/internal/bruteforce"
	"github.com/VikashChoudhary-04/subhunt/internal/passive"
	"github.com/VikashChoudhary-04/subhunt/internal/resolver"
	"github.com/VikashChoudhary-04/subhunt/internal/utils"
)

func main() {
	domain := flag.String("d", "", "Target domain")
	passiveEnum := flag.Bool("passive", false, "Enable passive enumeration")
	bruteforceList := flag.String("bruteforce", "", "Wordlist for DNS bruteforce")
	resolve := flag.Bool("resolve", false, "Resolve subdomains")
	flag.Parse()

	if *domain == "" {
		fmt.Println("Usage: subhunt -d example.com [--passive] [--bruteforce wordlist] [--resolve]")
		os.Exit(1)
	}

	var subs []string

	if *passiveEnum {
		results, err := passive.CRTSH(*domain)
		if err != nil {
			fmt.Println("[!] crt.sh failed:", err)
		} else {
			subs = append(subs, results...)
		}
	}


	if *bruteforceList != "" {
		results := bruteforce.Brute(*domain, *bruteforceList)
		subs = append(subs, results...)
	}

	subs = utils.Dedupe(subs)

	if *resolve {
		subs = resolver.Resolve(subs)
	}

	for _, s := range subs {
		fmt.Println(s)
	}
}

package main

import (
	"flag"
	"fmt"

	"github.com/rebay1982/bhg/dns/dnsenum/pkg"
)

type result struct {
	IPAddress string
	Hostname  string
}

func main() {
	var (
		flDomain = flag.String("domain", "", "The domain on which to perform the guessing on")
		//flWordlist    = flag.String("wordlist", "", "The world list to use for enumeration")
		//flWorkerCount = flag.Int("c", 100, "The amount of workers to use")
		flServerAddr = flag.String("server", "8.8.8.8:53", "The DNS server to use")
	)

	flag.Parse()

	if *flDomain == "" { //|| *flWordlist == "" {
		fmt.Println("-domain and -worldlist are required")
	}

	fmt.Printf("Looking up [%s] against [%s]\n", *flDomain, *flServerAddr)
	results := dolookup(*flDomain, *flServerAddr)

	for _, result := range results {
		fmt.Printf("[%s]: [%s]\n", result.Hostname, result.IPAddress)
	}
}

func dolookup(fqdn, serverAddr string) []result {
	var results []result
	var cfqdn = fqdn

	// First check CNAMES
	for {
		cnames, err := lookup.LookupCNAME(cfqdn, serverAddr)
		if err == nil && len(cnames) > 0 {

			for _, cname := range cnames {
				fmt.Printf("Found cname [%s]\n", cname)
			}

			cfqdn = cnames[0] // Pick first CNAME
			continue
		}

		ips, err := lookup.LookupA(cfqdn, serverAddr)
		if err != nil {
			fmt.Println("DNS: Unable to find A Records.")
			break
		}

		for _, ip := range ips {
			results = append(results, result{IPAddress: ip, Hostname: fqdn})
		}

		break
	}
	return results
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rebay1982/bhg/dns/dnsenum/pkg/lookup"
	"github.com/rebay1982/bhg/dns/dnsenum/pkg/worker"
)

func main() {
	var (
		flDomain      = flag.String("domain", "", "The domain on which to perform the guessing on")
		flWordlist    = flag.String("wordlist", "", "The world list to use for enumeration")
		flWorkerCount = flag.Int("c", 100, "The amount of workers to use")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use")
	)

	flag.Parse()

	if *flDomain == "" { //|| *flWordlist == "" {
		fmt.Println("-domain and -worldlist are required")
	}

	var results []lookup.Result
	fqdns := make(chan string, *flWorkerCount)
	gather := make(chan []lookup.Result)
	tracker := make(chan worker.Empty)

	// Open the word list for scanning.
	fh, err := os.Open(*flWordlist)
	if err != nil {
		panic(err)
	}

	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	for i := 0; i < *flWorkerCount; i++ {
		go worker.Worker(tracker, fqdns, gather, *flServerAddr)
	}

	for scanner.Scan() {
		fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), *flDomain)
	}

	go func() {
		for r := range gather {
			results = append(results, r...)
		}
		var e worker.Empty
		tracker <- e
	}()

	close(fqdns)
	for i := 0; i < *flWorkerCount; i++ {
		<-tracker
	}
	close(gather)
	<-tracker

	// Print out results with a tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Hostname, r.IPAddress)
	}

	w.Flush()

}

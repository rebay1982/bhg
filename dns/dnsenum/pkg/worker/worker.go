package worker

import (
	"github.com/rebay1982/bhg/dns/dnsenum/pkg/lookup"
)

type Empty struct{} // Empty struct -- used to signify a worker is done.

func Worker(tracker chan Empty, fqdns chan string, gather chan []lookup.Result, serverAddr string) {
	for fqdn := range fqdns {
		results := lookup.DoLookup(fqdn, serverAddr)

		if len(results) > 0 {
			gather <- results
		}
	}

	var e Empty
	tracker <- e
}

package lookup

import (
	"errors"
	"fmt"
	"github.com/miekg/dns"
)

type Result struct {
	IPAddress string
	Hostname  string
}

// LookupA Looks up an A record for the fqdn against the serverAddr
func LookupA(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string

	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)

	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}

	if len(in.Answer) < 1 {
		return ips, errors.New("DNS: No answer.")
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips, nil
}

// LookupCNAME Looks up a CNAME
func LookupCNAME(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var fqdns []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)

	if err != nil {
		return fqdns, err
	}

	if len(in.Answer) < 1 {
		return fqdns, errors.New("DNS: No answer.")
	}
	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}

	return fqdns, nil

}

// DoLookup does a DNS lookup
func DoLookup(fqdn, serverAddr string) []Result {
	var results []Result
	var cfqdn = fqdn

	// First check CNAMES
	for {
		cnames, err := LookupCNAME(cfqdn, serverAddr)
		if err == nil && len(cnames) > 0 {

			for _, cname := range cnames {
				fmt.Printf("Found cname [%s]\n", cname)
			}

			cfqdn = cnames[0] // Pick first CNAME
			continue
		}

		ips, err := LookupA(cfqdn, serverAddr)
		if err != nil {
			fmt.Println("DNS: Unable to find A Records.")
			break
		}

		for _, ip := range ips {
			results = append(results, Result{IPAddress: ip, Hostname: fqdn})
		}

		break
	}
	return results
}

package lookup

import (
	"errors"
	"github.com/miekg/dns"
)

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

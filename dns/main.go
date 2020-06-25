package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg

	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)

	// use 8.8.8.8:53 for testing
	in, err := dns.Exchange(&msg, "127.0.0.1:53")
	if err != nil {
		panic(err)
	}

	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}

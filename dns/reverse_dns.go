package main

import (
	"flag"
	"fmt"

	"github.com/miekg/dns"
)

var (
	address string
)

func init() {
	flag.StringVar(&address, "addr", "", "Address to resolve")
	flag.Parse()
}

func main() {
	var msg dns.Msg

	fqdn := dns.Fqdn(address)
	msg.SetQuestion(fqdn, dns.TypeA)

	in, err := dns.Exchange(&msg, "8.8.8.8:53")
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

package main

import (
	"github.com/rebay1982/bhg/shodan/shodan"

	"fmt"
	"log"
	"os"
)

func main() {

	apikey := os.Getenv("SHODAN_API_KEY")
	c := shodan.New(apikey)

	info, err := c.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Query Credits [%d]\nScan Credits [%d]\nTelnet? [%t]\n",
		info.QueryCredits,
		info.ScanCredits,
		info.Telnet)
}

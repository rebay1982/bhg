package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Message string
	Status  string
}

func main() {

	// Reach out to google.com/robots.txt
	res, err := http.Post(
		"https://www.google.com/robots.txt",
		"application/json",
		nil,
	)

	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	log.Printf("%s -> %s\n", status.Status, status.Message)
}

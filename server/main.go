package main

import (
	"fmt"
	"log"
	"net/http"
)

// Our logging middleware
type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("Start [%s]\n", req.URL.Query().Get("name"))
	l.Inner.ServeHTTP(w, req)
	log.Println("Stop")
}

// Our basic HTTP serving function.
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}

	// Use the router here. --> &r
	http.ListenAndServe(":8000", &l)
}

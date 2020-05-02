package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// HANDLER FUNC  Our basic HTTP serving function.
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/users/{user:[a-z]+}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "User [%s]", user)
	}).Methods("GET")

	//f := http.HandlerFunc(hello)
	//l := logger{Inner: f}

	// Use the router here. --> &r

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)

}

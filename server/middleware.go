package main

import (
	"fmt"
	"log"
	"net/http"
)

func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving [%s]\n", r.URL.Path)

		h.ServeHTTP(w, r)

		log.Printf("Done [%s]\n", r.URL.Path)
	})
}

func RedirectMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			log.Printf("Redirecting [%s] to /", r.URL.Path)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)

			return
		}

		h.ServeHTTP(w, r)
	})
}

// Our basic HTTP serving function.
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello [%s]\n", r.URL.Query().Get("name"))
}

func main() {

	h := http.HandlerFunc(hello)
	r := RedirectMiddleware(h)
	l := LogMiddleware(r)

	http.ListenAndServe(":8080", l)

}

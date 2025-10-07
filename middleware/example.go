package middleware

import (
	"log"
	"net/http"
)

func Example(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("I am the example")
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // 8.55.20s

		log.Println("i was first")

		next.ServeHTTP(w, r) // 10s

		log.Println("I was second")

		log.Println(r.Method, r.URL.Path, time.Since(start))

	})

}

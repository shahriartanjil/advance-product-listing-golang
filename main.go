package main

import (
	"fmt"

	"ecommere.com/config"
)

func main() {
	cnf := config.GetConfig()

	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)

	// cmd.Serve()
}

// func corsMiddleware(next http.Handler) http.Handler {
// 	handleCors := func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Shahriar")
// 		w.Header().Set("Content-Type", "application/json")

// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)

// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	}

// 	handler := http.HandlerFunc(handleCors)

// 	return handler
// }

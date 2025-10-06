package cmd

import (
	"fmt"
	"net/http"

	"ecommere.com/global_router"
	"ecommere.com/handlers"
	"ecommere.com/middleware"
)

func Serve() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct)) //route
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))

	globRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :3001")
	err := http.ListenAndServe(":3001", globRouter) // "failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

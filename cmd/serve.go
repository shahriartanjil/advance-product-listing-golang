package cmd

import (
	"fmt"
	"net/http"

	"ecommere.com/global_router"
	"ecommere.com/handlers"
	"ecommere.com/middleware"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux() //router

	mux.Handle("GET /Shahriar", manager.With(middleware.Example, middleware.Logger)(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /route", middleware.Example(middleware.Logger(http.HandlerFunc(handlers.Test))))

	mux.Handle("GET /products", middleware.Example(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))
	mux.Handle("POST /products", middleware.Example(middleware.Logger(http.HandlerFunc(handlers.CreateProduct))))
	mux.Handle("GET /products/{productId}", middleware.Example(middleware.Logger(http.HandlerFunc(handlers.GetProductById))))

	globRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :3001")
	err := http.ListenAndServe(":3001", globRouter) // "failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

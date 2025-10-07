package cmd

import (
	"net/http"

	"ecommere.com/handlers"
	"ecommere.com/middleware"
)

func intiRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /Shahriar", manager.With(
		http.HandlerFunc(handlers.Test),
	))

	mux.Handle("GET /route", manager.With(
		http.HandlerFunc(handlers.Test),
	))

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	))
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	))
	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(handlers.GetProductById),
	))
}

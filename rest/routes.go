package rest

import (
	"net/http"

	"ecommere.com/rest/handlers"
	middleware "ecommere.com/rest/middlewares"
)

func intiRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		))

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		))

	mux.Handle("GET /products/{Id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		))

	mux.Handle("PUT /products/{Id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProducts),
		))

	mux.Handle("DELETE /products/{Id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
		))

}

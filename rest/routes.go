package rest

import (
	"net/http"

	"ecommere.com/rest/handlers"
	middleware "ecommere.com/rest/middlewares"
)

func intiRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// auth := middleware.AuthFromConfig()

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		))

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
			middleware.AuthenticateJwt,
		))

	mux.Handle("GET /products/{Id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		))

	mux.Handle("PUT /products/{Id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProducts),
			middleware.AuthenticateJwt,
		))

	mux.Handle("DELETE /products/{Id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.AuthenticateJwt,
		))

	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		))

	mux.Handle("POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		))

}

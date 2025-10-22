package product

import (
	"net/http"

	middleware "ecommere.com/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// auth := middleware.AuthFromConfig()

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		))

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthenticateJwt,
		))

	mux.Handle("GET /products/{Id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		))

	mux.Handle("PUT /products/{Id}",
		manager.With(
			http.HandlerFunc(h.UpdateProducts),
			h.middlewares.AuthenticateJwt,
		))

	mux.Handle("DELETE /products/{Id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJwt,
		))

}

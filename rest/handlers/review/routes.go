package review

import (
	"net/http"

	middleware "ecommere.com/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /reviews",
		manager.With(
			http.HandlerFunc(h.GetReviews),
		))
}

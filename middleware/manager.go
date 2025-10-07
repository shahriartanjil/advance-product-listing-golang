package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
func (mngr *Manager) With(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		n := next

		// middlewares = [example, logger] // idx = example - 0, logger -1
		//middleware.Logger(http.HandlerFunc(handlers.GetProducts))

		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			n = middleware(n)
		}

		return n
	}
}

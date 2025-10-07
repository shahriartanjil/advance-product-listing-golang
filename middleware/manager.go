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

func (mngr *Manager) Use(Middlewares ...Middleware) { /// builder pattern bola hoy
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, Middlewares...)

}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {

	n := next

	//middlewares = [ logger, example]
	// n = middleware.Logger(http.HandlerFunc(handlers.CreateProduct))
	//middlewares = [ logger, example]
	//n = middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	//globalMiddleware
	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}

	return n
}

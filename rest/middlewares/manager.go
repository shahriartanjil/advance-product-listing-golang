package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware //Logger, Example,CorsWithPreflight
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

	h := next

	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler

	//[CorsWithPreflight, Example,Logger]
	// h= logger(example(corsWithPreflight(mux)))
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}

	return h
}

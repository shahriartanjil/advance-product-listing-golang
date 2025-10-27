package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ecommere.com/config"
	"ecommere.com/rest/handlers/product"
	"ecommere.com/rest/handlers/user"
	middleware "ecommere.com/rest/middlewares"
)

type Server struct {
	cnf            config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// Register routes
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("✅ Server running on port", server.cnf.HttpPort)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		fmt.Println("❌ Error starting the server:", err)
		os.Exit(1)
	}
}

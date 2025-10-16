package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ecommere.com/config"
	middleware "ecommere.com/rest/middlewares"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()
	//router
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
		middleware.Auth(config.GetConfig().JwtSecretKey)
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	intiRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux) // "failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}

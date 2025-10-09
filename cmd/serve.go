package cmd

import (
	"fmt"
	"net/http"

	"ecommere.com/middleware"
)

func Serve() {

	manager := middleware.NewManager()
	//router
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	intiRoutes(mux, manager)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", wrappedMux) // "failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

package cmd

import (
	"fmt"
	"net/http"

	"ecommere.com/global_router"
	"ecommere.com/middleware"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Example)

	mux := http.NewServeMux() //router

	intiRoutes(mux, manager)

	globRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", globRouter) // "failed to start the server"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

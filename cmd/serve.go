package cmd

import (
	"ecommere.com/config"
	"ecommere.com/rest"
	"ecommere.com/rest/handlers/product"
	"ecommere.com/rest/handlers/review"
	"ecommere.com/rest/handlers/user"
	middleware "ecommere.com/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(&cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()

}

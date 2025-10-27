package cmd

import (
	"ecommere.com/config"
	"ecommere.com/repo"
	"ecommere.com/rest"
	"ecommere.com/rest/handlers/product"
	"ecommere.com/rest/handlers/user"
	middleware "ecommere.com/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		*cnf,
		productHandler,
		userHandler,
	)
	server.Start()

}

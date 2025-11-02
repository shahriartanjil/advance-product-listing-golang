package cmd

import (
	"fmt"
	"os"

	"ecommere.com/config"
	"ecommere.com/infra/db"
	"ecommere.com/repo"
	"ecommere.com/rest"
	"ecommere.com/rest/handlers/product"
	"ecommere.com/rest/handlers/user"
	middleware "ecommere.com/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	fmt.Println("%+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		*cnf,
		productHandler,
		userHandler,
	)
	server.Start()

}

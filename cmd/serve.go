package cmd

import (
	"fmt"
	"os"

	"ecommere.com/config"
	"ecommere.com/infra/db"
	"ecommere.com/product"
	"ecommere.com/repo"
	"ecommere.com/rest"
	prdHandler "ecommere.com/rest/handlers/product"
	usrHandler "ecommere.com/rest/handlers/user"
	middleware "ecommere.com/rest/middlewares"
	"ecommere.com/user"
)

func Serve() {
	cnf := config.GetConfig()

	// fmt.Printf("%+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Migration failed:", err)
		os.Exit(1)
	}

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	// handlers
	productHandler := prdHandler.NewHandler(middlewares, prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		*cnf,
		productHandler,
		userHandler,
	)
	server.Start()

}

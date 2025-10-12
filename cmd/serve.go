package cmd

import (
	"ecommere.com/config"
	"ecommere.com/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf)

}

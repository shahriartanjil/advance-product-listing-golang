package db

import (
	"fmt"

	"ecommere.com/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // This line registers the "postgres" driver
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"

	}
	return connString
}

// func GetConnectionString(cnf config.DBConfig) string {

// 	return "user=postgres password=123456789 host=localhost port=5432 dbname=ecommerce sslmode=disable"
// }

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}

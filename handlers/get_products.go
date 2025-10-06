package handlers

import (
	"net/http"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utility.SendData(w, database.ProductList, 200)
}

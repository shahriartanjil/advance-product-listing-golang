package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)

	utility.SendData(w, newProduct, 201)

}

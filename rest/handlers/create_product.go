package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) { 
	// parse jwt
	// parse header and payload or claims
	//hmac-sha-256 algo -> hash hmac(header, payload, secret key)
	// parse signature part from the jwt
	//if the signature and hash is same => forward to create products
	//otherwise 401 status code with Unauthorized

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	utility.SendData(w, createdProduct, 201)

}

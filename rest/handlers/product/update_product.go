package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func (h *Handler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("Id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid json", 400)
		return
	}

	newProduct.ID = pId

	database.Update(newProduct)

	utility.SendData(w, "Successfully updated product", 201)
}

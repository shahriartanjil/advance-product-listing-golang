package handlers

import (
	"net/http"
	"strconv"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}
	for _, product := range database.ProductList {
		if product.ID == pId {
			utility.SendData(w, product, 200)
			return
		}
	}

	utility.SendData(w, "no data availabe", 400)
}

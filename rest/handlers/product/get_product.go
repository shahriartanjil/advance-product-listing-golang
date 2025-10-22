package product

import (
	"net/http"
	"strconv"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("Id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}
	product := database.Get(pId)
	if product == nil {
		utility.SendError(w, 404, "Product not found")
		return
	}

	utility.SendData(w, product, 200)
}

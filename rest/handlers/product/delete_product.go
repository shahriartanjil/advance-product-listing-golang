package product

import (
	"net/http"
	"strconv"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("Id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}
	database.Delete(pId)

	utility.SendData(w, "Successfully deleted product", 201)
}

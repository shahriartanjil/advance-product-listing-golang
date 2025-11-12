package product

import (
	"net/http"
	"strconv"

	"ecommere.com/utility"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("Id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		utility.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	product, err := h.svc.Get(pId)
	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal server error ")
		return
	}
	if product == nil {
		utility.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	utility.SendData(w, http.StatusOK, product)
}

package product

import (
	"net/http"

	"ecommere.com/utility"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.svc.List()
	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utility.SendData(w, http.StatusOK, productList)
	// utility.SendData(w, map[string]string{"message": "Success"}, http.StatusOK)
}

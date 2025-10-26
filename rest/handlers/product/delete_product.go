package product

import (
	"fmt"
	"net/http"
	"strconv"

	"ecommere.com/utility"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("Id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusBadRequest, "Invalid product id")
		return
	}
	err = h.productRepo.Delete(pId)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusInternalServerError, "Internal sever error")
		return

	}

	utility.SendData(w, http.StatusOK, "Successfully deleted product")
}

package product

import (
	"net/http"
	"strconv"

	"ecommere.com/utility"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	// get page from query params
	// get limit from query params
	reqQuery := r.URL.Query()

	paseAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(paseAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal server error 1")
		return
	}

	utility.SendData(w, http.StatusOK, productList)
	// utility.SendData(w, map[string]string{"message": "Success"}, http.StatusOK)
}

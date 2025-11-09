package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ecommere.com/domain"
	"ecommere.com/utility"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) ReqUpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("Id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		utility.SendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err = h.svc.Update(domain.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		utility.SendError(w, http.StatusInternalServerError, "Internal server error")
	}

	utility.SendData(w, http.StatusOK, "Successfully updated product")
}

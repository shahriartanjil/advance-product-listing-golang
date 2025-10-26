package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommere.com/repo"
	"ecommere.com/utility"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		utility.SendData(w, http.StatusInternalServerError, "internal server error")
		return
	}

	utility.SendData(w, http.StatusCreated, createdProduct)

}

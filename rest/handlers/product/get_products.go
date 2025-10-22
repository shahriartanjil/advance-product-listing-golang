package product

import (
	"net/http"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hello")
	utility.SendData(w, database.List(), 200)
	// utility.SendData(w, map[string]string{"message": "Success"}, http.StatusOK)
}

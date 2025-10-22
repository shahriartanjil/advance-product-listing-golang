package review

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommere.com/database"
	"ecommere.com/utility"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request Data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	utility.SendData(w, createdUser, http.StatusCreated)

}

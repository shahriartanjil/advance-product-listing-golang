package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommere.com/database"
	"ecommere.com/utility"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid users", http.StatusBadRequest)
		return
	}

	usr := database.Find(userLogin.Email, userLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	utility.SendData(w, usr, http.StatusCreated)

}

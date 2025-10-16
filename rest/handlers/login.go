package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommere.com/config"
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

	cnf := config.GetConfig()

	accessToken, err := utility.CreateJwt(cnf.JwtSecretKey, utility.Payload{ //jwt secret key is called access token
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utility.SendData(w, accessToken, http.StatusCreated)

}

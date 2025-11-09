package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommere.com/utility"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utility.SendError(w, http.StatusBadRequest, "Invalid user")
		return
	}
	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		utility.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr == nil {
		utility.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := utility.CreateJwt(h.cnf.JwtSecretKey, utility.Payload{ //jwt secret key is called access token
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		// Password:  usr.Password,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utility.SendData(w, http.StatusCreated, accessToken)

}

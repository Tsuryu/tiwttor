package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Tsuryu/tiwttor/datamodel"
	"github.com/Tsuryu/tiwttor/db/userdao"
	"github.com/Tsuryu/tiwttor/jwt"
	"github.com/Tsuryu/tiwttor/model"
)

// Login : handle login from http request
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid user or password"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is mandatory"+err.Error(), 400)
		return
	}

	document, result := userdao.Login(user.Email, user.Password)
	if !result {
		http.Error(w, "Invalid user or password", 400)
		return
	}

	jwtKey, err := jwt.CreateJWT(document)
	if err != nil {
		http.Error(w, "Internal error on jwt creation"+err.Error(), 400)
		return
	}

	response := datamodel.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	// save on cookie
	expirationTime := time.Now().Add(24 * time.Hour) // adds 24 hours
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Tsuryu/tiwttor/db/userdao"
	"github.com/Tsuryu/tiwttor/model"
)

// Register : handle register user request
func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid user data "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is mandatory", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Invalid password, at least 6 characters are required", http.StatusBadRequest)
		return
	}

	_, result, _ := userdao.FindBy(user.Email)
	if result {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	_, status, err := userdao.Insert(user)
	if err != nil {
		http.Error(w, "Failed to create user "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

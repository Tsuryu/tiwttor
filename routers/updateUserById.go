package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Tsuryu/tiwttor/db/userdao"
	"github.com/Tsuryu/tiwttor/model"
)

// UpdateUserByID : updates an user by id
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid body"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = userdao.UpdateByID(user, UserID)
	if err != nil {
		http.Error(w, "Failed to update user profile."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

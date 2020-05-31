package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Tsuryu/tiwttor/db/userdao"
)

// FindUserByID : router to fetch user by id
func FindUserByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Parameter ID is mandatory", http.StatusBadRequest)
		return
	}

	profile, err := userdao.FindByID(ID)
	if err != nil {
		http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

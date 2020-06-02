package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tsuryu/tiwttor/db/userdao"
)

// FindManyUserBy : fetch many users
func FindManyUserBy(w http.ResponseWriter, r *http.Request) {
	userType := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}

	result, status := userdao.FindManyByID(UserID, pageNumber, search, userType)
	if !status {
		http.Error(w, "Failed to fetch users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

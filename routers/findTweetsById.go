package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tsuryu/tiwttor/db/tweetdao"
)

// FindTweetsByID : fetch tweets by id
func FindTweetsByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page is mandatory", http.StatusBadRequest)
		return
	}

	// alphabet to integer
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page is mandatory", http.StatusBadRequest)
		return
	}

	result, status := tweetdao.FindByID(ID, page)
	if !status {
		http.Error(w, "Failed to fetch tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

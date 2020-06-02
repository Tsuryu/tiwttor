package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tsuryu/tiwttor/db/tweetdao"
)

// FindManyTweetByID : fetch my followers tweet
func FindManyTweetByID(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page is mandatory", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}

	tweets, result := tweetdao.FindManyByID(UserID, page)
	if !result {
		http.Error(w, "Failed to fetch tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

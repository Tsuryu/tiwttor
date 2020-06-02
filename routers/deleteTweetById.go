package routers

import (
	"net/http"

	"github.com/Tsuryu/tiwttor/db/tweetdao"
)

// DeleteTweetByID : deletes a tweet
func DeleteTweetByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	err := tweetdao.DeleteByID(ID, UserID)
	if err != nil {
		http.Error(w, "Failed to delete a tweet."+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

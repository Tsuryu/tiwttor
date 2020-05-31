package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Tsuryu/tiwttor/db/tweetdao"
	"github.com/Tsuryu/tiwttor/model"
)

// InsertTweet : creates a tweet
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var tweet model.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, "Invalid tweet data "+err.Error(), http.StatusBadRequest)
		return
	}

	register := model.Tweet{
		UserID:  UserID,
		Message: tweet.Message,
		Date:    time.Now(),
	}

	_, status, err := tweetdao.Insert(register)
	if err != nil {
		http.Error(w, "Failed to create tweet."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Failed to create tweet."+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

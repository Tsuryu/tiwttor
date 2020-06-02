package routers

import (
	"net/http"

	"github.com/Tsuryu/tiwttor/db/followerdao"
	"github.com/Tsuryu/tiwttor/model"
)

// InsertFollower : creates a relation between users
func InsertFollower(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	var follower model.Follower
	follower.UserID = UserID
	follower.UserFollowedID = ID

	status, err := followerdao.Insert(follower)
	if err != nil || !status {
		http.Error(w, "Failed to follow", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

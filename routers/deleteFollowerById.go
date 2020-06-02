package routers

import (
	"net/http"

	"github.com/Tsuryu/tiwttor/db/followerdao"
	"github.com/Tsuryu/tiwttor/model"
)

// DeleteFollowerByID : deletes a relation between users
func DeleteFollowerByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var follower model.Follower
	follower.UserID = UserID
	follower.UserFollowedID = ID

	status, err := followerdao.DeleteByID(follower)
	if err != nil || !status {
		http.Error(w, "Failed to unfollow", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

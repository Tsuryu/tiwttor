package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Tsuryu/tiwttor/db/followerdao"
	"github.com/Tsuryu/tiwttor/model"
)

// FindFollowerByID : check if the relation between users exists
func FindFollowerByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	var follower model.Follower
	follower.UserID = UserID
	follower.UserFollowedID = ID

	var response model.CommonResponse

	status, err := followerdao.FindByID(follower)
	if err != nil || !status {
		response.Result = "Error"
		w.WriteHeader(http.StatusNotFound)
	} else {
		response.Result = "OK"
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

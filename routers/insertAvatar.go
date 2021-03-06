package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Tsuryu/tiwttor/db/userdao"
	"github.com/Tsuryu/tiwttor/model"
)

// InsertAvatar : use to customize user profile avatar
func InsertAvatar(w http.ResponseWriter, r *http.Request) {
	requestFile, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileName string = UserID + "." + extension
	var filePath string = "upload/avatar/" + fileName

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Failed to upload the avatar."+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(file, requestFile)
	if err != nil {
		http.Error(w, "Failed to upload the avatar."+err.Error(), http.StatusBadRequest)
		return
	}

	var user model.User
	var status bool

	user.Avatar = fileName
	status, err = userdao.UpdateByID(user, UserID)
	if err != nil || !status {
		http.Error(w, "Failed to save the avatar."+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Tsuryu/tiwttor/db/userdao"
)

// FindBannerByID : get banner from folder
func FindBannerByID(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	user, err := userdao.FindByID(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	file, err := os.Open("upload/banner/" + user.Banner)
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to copy the banner", http.StatusNotFound)
	}
}

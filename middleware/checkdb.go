package middleware

import (
	"net/http"

	"github.com/Tsuryu/tiwttor/db"
)

// CheckDB : validates database connection status
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Database connection lost", 500)
		}

		next.ServeHTTP(w, r)
	}
}

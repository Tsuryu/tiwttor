package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/Tsuryu/tiwttor/middleware"
	"github.com/Tsuryu/tiwttor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers : set handler, port and runs server
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

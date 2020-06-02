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
	router.HandleFunc("/profile", middleware.CheckDB(middleware.CheckJWT(routers.FindUserByID))).Methods("GET")
	router.HandleFunc("/profile", middleware.CheckDB(middleware.CheckJWT(routers.UpdateUserByID))).Methods("PUT")
	router.HandleFunc("/profiles", middleware.CheckDB(middleware.CheckJWT(routers.FindManyUserBy))).Methods("GET")

	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routers.InsertTweet))).Methods("POST")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routers.FindTweetsByID))).Methods("GET")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routers.DeleteTweetByID))).Methods("DELETE")
	router.HandleFunc("/tweets", middleware.CheckDB(middleware.CheckJWT(routers.FindManyTweetByID))).Methods("GET")

	router.HandleFunc("/avatar", middleware.CheckDB(middleware.CheckJWT(routers.InsertAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middleware.CheckDB(routers.FindAvatarByID)).Methods("GET")
	router.HandleFunc("/banner", middleware.CheckDB(middleware.CheckJWT(routers.InsertBanner))).Methods("POST")
	router.HandleFunc("/banner", middleware.CheckDB(routers.FindBannerByID)).Methods("GET")

	router.HandleFunc("/follow", middleware.CheckDB(middleware.CheckJWT(routers.InsertFollower))).Methods("POST")
	router.HandleFunc("/follow", middleware.CheckDB(middleware.CheckJWT(routers.DeleteFollowerByID))).Methods("DELETE")
	router.HandleFunc("/follow", middleware.CheckDB(middleware.CheckJWT(routers.FindFollowerByID))).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

package main

import (
	"fmt"
	"log"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/handler"
)

func main() {
	fmt.Println("Starting server")
	if !db.CheckConnection() {
		log.Fatal("No database connection")
		return
	}

	fmt.Println("Server up and running")
	handler.Handlers()
}

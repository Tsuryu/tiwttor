package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection : client to work with mongodb
var Connection = ConnectDB()

var clientOptions = options.Client().ApplyURI(os.Getenv("MONGO_URL"))

// ConnectDB : to connect to mongo
func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Mongo connection up")
	return client
}

// CheckConnection : ping mongo server to check if it is up
func CheckConnection() bool {
	err := Connection.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}

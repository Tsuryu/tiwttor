package followerdao

import (
	"context"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
)

// Insert : creates a relation between users
func Insert(follower model.Follower) (bool, error) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := db.Connection.Database("twittor")
	collection := db.Collection("follower")

	_, err := collection.InsertOne(context, follower)
	if err != nil {
		return false, err
	}

	return true, nil
}

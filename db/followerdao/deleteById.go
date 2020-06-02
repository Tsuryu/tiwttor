package followerdao

import (
	"context"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
)

// DeleteByID : delete a relation between users
func DeleteByID(follower model.Follower) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := db.Connection.Database("twittor")
	collection := db.Collection("follower")

	_, err := collection.DeleteOne(ctx, follower)
	if err != nil {
		return false, err
	}

	return true, nil
}

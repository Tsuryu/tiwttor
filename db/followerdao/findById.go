package followerdao

import (
	"context"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
	"go.mongodb.org/mongo-driver/bson"
)

// FindByID : fetch a relation between users
func FindByID(follower model.Follower) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := db.Connection.Database("twittor")
	collection := db.Collection("follower")

	filter := bson.M{
		"userId":         follower.UserID,
		"userFollowedId": follower.UserFollowedID,
	}

	var result model.Follower
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}

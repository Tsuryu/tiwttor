package tweetdao

import (
	"context"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteByID : deletes a tweet
func DeleteByID(ID string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := db.Connection.Database("twittor")
	collection := db.Collection("tweet")

	objectID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id":    objectID,
		"userId": userID,
	}

	_, err := collection.DeleteOne(ctx, filter)
	return err
}

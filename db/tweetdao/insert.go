package tweetdao

import (
	"context"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert : creates a tweet
func Insert(tweet model.Tweet) (string, bool, error) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := db.Connection.Database("twittor")
	collection := db.Collection("tweet")

	register := bson.M{
		"userId":  tweet.UserID,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := collection.InsertOne(context, register)
	if err != nil {
		return "", false, err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)
	return objectID.String(), true, nil
}

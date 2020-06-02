package tweetdao

import (
	"context"
	"log"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindByID : get tweets by user id
func FindByID(ID string, page int) ([]*model.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := db.Connection.Database("twittor")
	collection := db.Collection("tweet")

	var result []*model.Tweet

	filter := bson.M{
		"userId": ID,
	}

	// options := options.Find()
	var options options.FindOptions
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}}) // descendente
	options.SetSkip((int64(page) - 1) * 20)

	cursor, err := collection.Find(ctx, filter, &options)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var register model.Tweet
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}

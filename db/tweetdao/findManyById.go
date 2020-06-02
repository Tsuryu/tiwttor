package tweetdao

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/dto"
	"go.mongodb.org/mongo-driver/bson"
)

// FindManyByID : fetch my follower's tweets
func FindManyByID(ID string, page int) ([]dto.TweetDto, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := db.Connection.Database("twittor")
	collection := db.Collection("follower")

	// filter := make([]bson.M, 8)
	filter := []bson.M{}
	filter = append(filter, bson.M{"$match": bson.M{"userId": ID}})
	filter = append(filter, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userFollowedId",
			"foreignField": "userId",
			"as":           "tweet",
		},
	})
	// undo master-slave result
	filter = append(filter, bson.M{"$unwind": "$tweet"})
	filter = append(filter, bson.M{"$sort": bson.M{"date": -1}})     // -1 = desc, 1 = asc
	filter = append(filter, bson.M{"$skip": (int64(page) - 1) * 20}) // skip before limit, otherwise = error
	filter = append(filter, bson.M{"$limit": 20})

	cursor, err := collection.Aggregate(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	var result []dto.TweetDto

	err = cursor.All(ctx, &result)
	if err != nil {
		fmt.Println("Cursor failed." + err.Error())
		return result, false
	}

	return result, true
}

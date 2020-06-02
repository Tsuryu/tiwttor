package userdao

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/db/followerdao"
	"github.com/Tsuryu/tiwttor/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindManyByID : fetch many users
func FindManyByID(ID string, page int, search string, searchType string) ([]*model.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	database := db.Connection.Database("twittor")
	collection := database.Collection("user")

	var results []*model.User

	// order affects result
	var findOptions options.FindOptions
	findOptions.SetSkip((int64(page) - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"firstName": bson.M{
			"$regex": `(?i)` + search,
		},
	}

	cursor, err := collection.Find(ctx, query, &findOptions)
	if err != nil {
		fmt.Println("Failed to fetch users." + err.Error())
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Failed to set cursor." + err.Error())
			return results, false
		}

		var follower model.Follower
		follower.UserID = ID
		follower.UserFollowedID = user.ID.Hex()

		include = false

		found, err = followerdao.FindByID(follower)
		if searchType == "new" && !found {
			include = true
		}

		if searchType == "follow" && found {
			include = true
		}

		if follower.UserFollowedID == ID {
			include = false
		}

		if include {
			user.Password = ""
			user.Biography = ""
			user.Banner = ""
			user.WebSite = ""
			user.Location = ""
			user.Email = ""

			results = append(results, &user)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)

	return results, true
}

package userdao

import (
	"context"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateByID : updates user by id
func UpdateByID(user model.User, ID string) (bool, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	database := db.Connection.Database("twittor")
	collection := database.Collection("user")

	// map string, interface
	register := make(map[string]interface{})
	if len(user.FirstName) > 0 {
		register["firstName"] = user.FirstName
	}

	if len(user.LastName) > 0 {
		register["lastName"] = user.LastName
	}

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}

	if len(user.Location) > 0 {
		register["location"] = user.Location
	}

	if len(user.WebSite) > 0 {
		register["webSite"] = user.WebSite
	}

	register["birthDate"] = user.BirthDate

	updateString := bson.M{
		"$set": register,
	}

	objectID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": bson.M{
			"$eq": objectID, // gt
		},
	}

	_, err := collection.UpdateOne(context, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}

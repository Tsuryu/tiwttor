package userdao

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByID : fetch an user by id
func FindByID(ID string) (model.User, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	database := db.Connection.Database("twittor")
	collection := database.Collection("user")

	var profile model.User
	objectID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objectID,
	}

	err := collection.FindOne(context, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Profile not found")
		return profile, err
	}

	return profile, nil
}

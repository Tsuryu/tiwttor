package userdao

import (
	"time"

	"github.com/Tsuryu/tiwttor/db"
	"github.com/Tsuryu/tiwttor/model"
	"github.com/Tsuryu/tiwttor/util"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
)

// Insert : creates an user
func Insert(user model.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := db.Connection.Database("twittor")
	collection := database.Collection("user")

	user.Password, _ = util.Encrypt(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil
}

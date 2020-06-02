package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TweetMongoDto : struct required by cursor all
type TweetMongoDto struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"userId" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"userFollowedId" json:"userFollowedId,omitempty"`
	Tweet             struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}

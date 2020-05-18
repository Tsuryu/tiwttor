package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User : mongo db model, bson = dato de entrada, json = navegador
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDate time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	WebPage   string             `bson:"webPage" json:"webPage,omitempty"`
}

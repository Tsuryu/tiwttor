package model

import "time"

// Tweet : user tweeet struct
type Tweet struct {
	UserID  string    `bson:"userId" json:"userId,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}

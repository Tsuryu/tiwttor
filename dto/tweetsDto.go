package dto

import "github.com/Tsuryu/tiwttor/model"

// TweetDto : struct to fetch tweets
type TweetDto struct {
	UserID string      `bson:"userId" json:"userId,omitempty"`
	Tweet  model.Tweet `bson:"tweet" json:"tweet,omitempty"`
}

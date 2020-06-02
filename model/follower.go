package model

// Follower : relation between users
type Follower struct {
	UserID string `bson:"userId" json:"userId"`
	UserFollowedID string `bson:"userFollowedId" json:"userFollowedId"`
}
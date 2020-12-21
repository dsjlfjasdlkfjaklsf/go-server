package model

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment 评论类型
type Comment struct {

	// The ID of the Blog
	BlogID primitive.ObjectID `json:"BlogID" bson:"BlogID"`
	
	// The ID of the owner
	OwnID string `json:"OwnID" bson:"OwnID"`

	// The name of the one who comments
	OwnName string `json:"ownName" bson:"ownName"`

	// The time of the comment
	CommentTime int64 `json:"commentTime" bson:"commentTime"`

	// The content of the comment
	Content string `json:"content" bson:"content"`
}

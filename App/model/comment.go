package model

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {

	// The ID of the Blog
	BlogId primitive.ObjectID `json:"BlogId,omitempty"`

	// The name of the one who comments
	OwnName string `json:"ownName,omitempty"`

	// The time of the comment
	CommentTime float32 `json:"commentTime,omitempty"`

	// The content of the comment
	Content string `json:"content,omitempty"`
}

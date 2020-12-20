package model

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment 评论类型
type Comment struct {

	// The ID of the Blog
	BlogID primitive.ObjectID `json:"BlogId"`

	// The name of the one who comments
	OwnName string `json:"ownName"`

	// The time of the comment
	CommentTime int64 `json:"commentTime"`

	// The content of the comment
	Content string `json:"content"`
}

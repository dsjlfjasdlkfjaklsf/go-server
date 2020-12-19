package model

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog 博客
type Blog struct {
	BlogID primitive.ObjectID `json:"BlogID,omitempty" bson:"BlogID"`

	AuthorID primitive.ObjectID `json:"AuthorID,omitempty" bson:"AuthorID"`

	AuthorName string `json:"AuthorName,omitempty" bson:"AuthorName"`

	CreateTime int64 `json:"CreateTime,omitempty" bson:"CreateTime"`

	Title string `json:"Title,omitempty" bson:"Title"`

	Abstract string `json:"Abstract,omitempty" bson:"Abstract"`

	Content string `json:"Content,omitempty" bson:"Content"`
}

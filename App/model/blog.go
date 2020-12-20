package model

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog 博客
type Blog struct {
	BlogID primitive.ObjectID `json:"BlogID" bson:"BlogID"`

	AuthorID primitive.ObjectID `json:"AuthorID" bson:"AuthorID"`

	AuthorName string `json:"AuthorName" bson:"AuthorName"`

	CreateTime int64 `json:"CreateTime" bson:"CreateTime"`

	Title string `json:"Title" bson:"Title"`

	Abstract string `json:"Abstract" bson:"Abstract"`

	Content string `json:"Content" bson:"Content"`
}

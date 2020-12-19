package model

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {

	// The ID of the Blog
	BlogID primitive.ObjectID `json:"BlogID,omitempty"`

	// all of the tags
	Content string `json:"content,omitempty"`
}

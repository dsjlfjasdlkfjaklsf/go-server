package model

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {

	ID primitive.ObjectID `json:"ID,omitempty" bson:"ID"`

	Name string `json:"Name,omitempty" bson:"Name"`

	Password string `json:"Password,omitempty" bson:"Password"`

	Bio string `json:"Bio,omitempty" bson:"Bio"`

	Level float32 `json:"Level,omitempty" bson:"Level"`
}

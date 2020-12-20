package model

type UserInfo struct {
	ID string `json:"ID" bson:"ID"`

	Name string `json:"Name" bson:"Name"`

	Password string `json:"Password" bson:"Password"`

	Bio string `json:"Bio" bson:"Bio"`

	Level float32 `json:"Level" bson:"Level"`
}

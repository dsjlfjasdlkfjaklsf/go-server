package model

type RegistBody struct {

	ID string `json:"ID,omitempty"`

	Name string `json:"Name,omitempty"`

	Password string `json:"Password,omitempty"`
}
type LoginBody struct {

	ID string `json:"ID,omitempty"`

	Password string `json:"Password,omitempty"`
}
type BlogBody struct {

	Title string `json:"Title,omitempty"`

	Abstract string `json:"Abstract,omitempty"`

	Content string `json:"Content,omitempty"`
}
type Body3 struct {

	OwnName string `json:"ownName,omitempty"`

	Content string `json:"content,omitempty"`
}
type Body4 struct {

	Content string `json:"content,omitempty"`
}
package model

type RegistBody struct {

	ID string `json:"ID"`

	Name string `json:"Name"`

	Password string `json:"Password"`
}
type LoginBody struct {

	ID string `json:"ID"`

	Password string `json:"Password"`
}
type BlogBody struct {

	Title string `json:"Title"`

	Abstract string `json:"Abstract"`

	Content string `json:"Content"`
}
type TagBody struct {

	ID string `json:"ID"`

	Content string `json:"Content"`
}
type Body3 struct {

	OwnName string `json:"ownName"`

	Content string `json:"content"`
}
type Body4 struct {

	Content string `json:"content"`
}
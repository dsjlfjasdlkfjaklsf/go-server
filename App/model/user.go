package model

type UserInfo struct {

	ID string `json:"ID,omitempty"`

	Name string `json:"Name,omitempty"`

	Bio string `json:"Bio,omitempty"`

	Level float32 `json:"Level,omitempty"`
}

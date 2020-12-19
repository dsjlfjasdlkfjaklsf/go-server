package model

type Tag struct {

	// The ID of the Blog
	BlogID string `json:"BlogID,omitempty"`

	// all of the tags
	Content string `json:"content,omitempty"`
}

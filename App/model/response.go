package model

type Response struct {

	// true for success and false for failure
	State bool `json:"state"`

	// the return message of the interfaces
	Response interface{} `json:"response"`
}

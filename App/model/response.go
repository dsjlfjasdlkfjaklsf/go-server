package model

type Response struct {

	// true for success and false for failure
	State bool `json:"state"`

	// the return message of the interfaces
	Response interface{} `json:"response"`
}

type InlineResponse200 struct {

	// true for success and false for failure
	State bool `json:"state"`

	Response *UserInfo `json:"response"`
}

type InlineResponse2001 struct {
	State bool `json:"state"`

	Response []Blog `json:"response"`
}

type InlineResponse2002 struct {
	State bool `json:"state"`

	Response *Blog `json:"response"`
}

type InlineResponse2003 struct {
	State bool `json:"state"`

	Response *Comment `json:"response"`
}

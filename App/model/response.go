package model

type Response struct {

	// true for success and false for failure
	State bool `json:"state,omitempty"`

	// the return message of the interfaces
	Response string `json:"response,omitempty"`
}

type InlineResponse200 struct {

	// true for success and false for failure
	State bool `json:"state,omitempty"`

	Response *UserInfo `json:"response,omitempty"`
}

type InlineResponse2001 struct {

	State bool `json:"state,omitempty"`

	Response []Blog `json:"response,omitempty"`
}

type InlineResponse2002 struct {

	State bool `json:"state,omitempty"`

	Response *Blog `json:"response,omitempty"`
}

type InlineResponse2003 struct {

	State bool `json:"state,omitempty"`

	Response *Comment `json:"response,omitempty"`
}
package util_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"github.com/stretchr/testify/assert"
)

func TestSpiltPath(t *testing.T) {
	url := "/user/dakfjalfsdf/blogs"
	urls := strings.Split(url, "/")
	assert.Equal(t, "user", urls[1], "spiltPath error!")
	assert.Equal(t, "dakfjalfsdf", urls[2], "spiltPath error!")
	assert.Equal(t, "blogs", urls[3], "spiltPath error!")
}

func TestJsonMarshal(t *testing.T) {
	res := &model.Response{
		State:    false,
		Response: "asdasdf"}

	msg, _ := json.Marshal(res)
	assert.Equal(t, "{\"state\":false,\"response\":\"asdasdf\"}", string(msg), "")
}

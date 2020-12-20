package util_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiltPath(t *testing.T) {
	url := "/user/dakfjalfsdf/blogs"
	urls := strings.Split(url, "/")
	assert.Equal(t, "user", urls[1], "spiltPath error!")
	assert.Equal(t, "dakfjalfsdf", urls[2], "spiltPath error!")
	assert.Equal(t, "blogs", urls[3], "spiltPath error!")
}

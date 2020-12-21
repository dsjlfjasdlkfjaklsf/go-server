package util_test

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	. "github.com/dsjlfjasdlkfjaklsf/go-server/App/util"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	// 检查生成是否会失败
	token, err := CreateToken("123456","name")
	assert.Equal(t, nil, err, "CreateToken Test Error!")
	// 检查解析是否会成功
	uid, uname, err := ParseToken(token)
	assert.Equal(t, nil, err, "ParseToken Test Error1!")
	assert.Equal(t, "123456", uid, "ParseToken Test Fail!")
	assert.Equal(t, "name", uname, "ParseToken Test Fail!")
	// 检查历史token是否会被拒绝
	uid, uname, err = ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDg1MzQ3MDEsInVpZCI6IjEyMzQ1NiIsInVuYW1lIjoibmFtZSJ9.EujdglWwegCJfFaD8nd_YV3wY0rAzw7HHby4Yktjqrs")
	validationError, ok := err.(*jwt.ValidationError)
	assert.Equal(t, true, ok, "ParseToken Test Error2!")
	// 检查错误类型是否有值
	assert.NotEqual(t,
		uint32(0),
		validationError.Errors,
		"ParseToken Test Error3!")
}

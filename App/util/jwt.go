package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = "c1a877e1f469b"

// CreateToken 根据用户id生成token
func CreateToken(uid,uname string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"uname":uname,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析token
func ParseToken(token string) (string,string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string),claim.Claims.(jwt.MapClaims)["uname"].(string), nil
}

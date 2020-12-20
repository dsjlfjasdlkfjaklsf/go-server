package controller

import (
	"net/http"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/util"
)

// CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetUserByID 查看指定用户的信息
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// handler := util.CreateHandler(w, r)
	// id := handler.DecodePath(2)
	// user,err :=	Service.User.GetUserByID(id)
	// handler.Send(user)
}

// LoginUser 用户登录
func LoginUser(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	body := &model.LoginBody{}
	err := handler.DecodePost(&body)
	if err != nil {
		handler.Send("invalid request")
	}
	token, _ := util.CreateToken(body.ID)
	handler.Send(token)
}

// LogoutUser 用户登出
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	handler.Send("success")
}

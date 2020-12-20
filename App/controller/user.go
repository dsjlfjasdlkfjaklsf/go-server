package controller

import (
	"net/http"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/util"
)

// CreateUser 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	body := model.RegistBody{}
	err := handler.DecodePost(&body)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	_, err = Service.User.CreateUser(body.ID, body.Name, body.Password)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}

// GetUserByID 查看指定用户的信息
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	id := handler.DecodePath(2)
	user, err := Service.User.GetUserByID(id)
	user.Password = "********"
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send(user, true)
}

// LoginUser 用户登录
func LoginUser(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	body := model.LoginBody{}
	err := handler.DecodePost(&body)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	_, err = Service.User.LoginUser(body.ID, body.Password)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	token, _ := util.CreateToken(body.ID,body.Name)
	handler.Send(token, false)
}

// LogoutUser 用户登出
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	handler.Send("success", true)
}

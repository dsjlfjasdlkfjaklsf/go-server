package controller

import (
	"net/http"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/util"
)

func GetCommentByBlogId(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	id := handler.DecodePath(2)
	comments, err := Service.Comment.GetCommentsByBlogID(id)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send(comments, true)
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	body := model.CommentBody{}
	err := handler.DecodePost(&body)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	token := handler.DecodeToken()
	uid, uname, err := util.ParseToken(token)
	_, err = Service.Comment.PostComment(body.BlogID, &model.Comment{
		BlogID:  body.BlogID,
		OwnID:   uid,
		OwnName: uname,
		Content: body.Content,
	})
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}

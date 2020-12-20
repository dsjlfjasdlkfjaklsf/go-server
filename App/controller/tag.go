package controller

import (
	"net/http"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/util"
)

func GetTagByBlogId(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	id := handler.DecodePath(2)
	tags, err := Service.Tag.GetTagByBlogID(id)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send(tags, true)
}

func PostTag(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	body := model.TagBody{}
	err := handler.DecodePost(&body)
	tag := model.Tag{}
	tag.Content = body.Content
	blogID := handler.DecodePath(2)
	_, err = Service.Tag.PostTag(blogID, tag)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}

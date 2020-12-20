package controller

import (
	// "github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"net/http"
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
	ID, err := primitive.ObjectIDFromHex(body.ID)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	tag.BlogID = ID
	tag.Content = body.Content

	_, err = Service.Tag.PostTag(tag)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}

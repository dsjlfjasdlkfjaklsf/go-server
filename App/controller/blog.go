package controller

import (
	// "github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"net/http"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/util"
)

func AddBlog(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	body := model.RegistBody{}
	err := handler.DecodePost(&body)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	_, err = Service.Blog.AddBlog(body)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	id := handler.DecodePath(2)
	err = Service.Blog.DeleteBlog(id)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send("success", true)
}

func GetBlogByBlogID(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	id := handler.DecodePath(2)
	blog, err := Service.Blog.GetBlogByBlogID(id)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send(blog, true)
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	blogs, err := Service.Blog.GetBlogs()
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send(blogs, true)
}

func GetBlogsByID(w http.ResponseWriter, r *http.Request) {
	handler := util.CreateHandler(w, r)
	id := handler.DecodePath(2)
	blogs, err := Service.Blog.GetBlogsByID(id)
	if err != nil {
		handler.Send(err.Error(), false)
		return
	}
	handler.Send(blogs, true)
}

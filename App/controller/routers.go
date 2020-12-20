package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/service"
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/util"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var Service *service.Service

func NewRouter(s *service.Service) *mux.Router {
	Service = s
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = util.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index 首页请求，用于测试
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddBlog",
		strings.ToUpper("Post"),
		"/blog",
		AddBlog,
	},

	Route{
		"DeleteBlog",
		strings.ToUpper("Delete"),
		"/blog/{BlogID}",
		DeleteBlog,
	},

	Route{
		"GetBlogByBlogID",
		strings.ToUpper("Get"),
		"/blog/{BlogID}",
		GetBlogByBlogID,
	},

	Route{
		"GetBlogs",
		strings.ToUpper("Get"),
		"/blogs",
		GetBlogs,
	},

	Route{
		"GetBlogsByID",
		strings.ToUpper("Get"),
		"/blogs/{ID}",
		GetBlogsByID,
	},

	Route{
		"GetCommentByBlogId",
		strings.ToUpper("Get"),
		"/comment/{BlogID}",
		GetCommentByBlogId,
	},

	Route{
		"PostComment",
		strings.ToUpper("Post"),
		"/comment/{BlogID}",
		PostComment,
	},

	Route{
		"GetTagByBlogId",
		strings.ToUpper("Get"),
		"/tag/{BlogID}",
		GetTagByBlogId,
	},

	Route{
		"PostTag",
		strings.ToUpper("Post"),
		"/tag/{BlogID}",
		PostTag,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/user/signup",
		CreateUser,
	},

	Route{
		"GetUserByID",
		strings.ToUpper("Get"),
		"/user/{ID}/info",
		GetUserByID,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Post"),
		"/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		strings.ToUpper("Get"),
		"/user/logout",
		LogoutUser,
	},
}

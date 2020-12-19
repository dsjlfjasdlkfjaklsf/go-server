package service

import (
	"context"
	"time"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

// BlogService 博客的服务
type BlogService struct {
	DB *mgo.Collection
}

// AddBlog 添加博客
func (service *BlogService) AddBlog(blog model.Blog) (primitive.ObjectID, error) {
	blog.BlogID = primitive.NewObjectID()
	blog.CreateTime = time.Now().Unix()
	_, err := service.DB.InsertOne(context.Background(), &blog)
	return blog.BlogID, err
}

// DeleteBlog 删除博客
func (service *BlogService) DeleteBlog(id string) (err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = service.DB.DeleteOne(context.Background(), bson.D{{"_id", ID}})
	return
}

// GetBlogByBlogID 获取到某个博客
func (service *BlogService) GetBlogByBlogID(blogID string) (blog model.Blog, err error) {
	return
}

// GetBlogs 获取所有人的所有博客
func (service *BlogService) GetBlogs() (blogs []model.Blog, err error) {

	return
}

// GetBlogsByID 获取到某个人的所有
func (service *BlogService) GetBlogsByID(userID string) (blogs []model.Blog, err error) {
	return
}

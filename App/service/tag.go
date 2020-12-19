package service

import (
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TagService 标签服务
type TagService struct {
	DB *mongo.Collection
}

// GetTagByBlogID 获取某个博客的所有tag
func (service *TagService) GetTagByBlogID(id string) (tag model.Tag, err error) {
	return
}

// PostTag 修改每个博客的tag
func (service *TagService) PostTag(tag model.Tag) (ID primitive.ObjectID, err error) {
	return
}

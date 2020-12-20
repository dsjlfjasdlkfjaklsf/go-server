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
func (service *TagService) GetTagByBlogID(id string) (tags []model.Tag, err error) {
	var cursor *Cursor
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	cursor, err = service.DB.Find(context.Background(), bson.D{{"BlogID", ID}})
	if err != nil {
		return
	}
	var tags []model.Tag
	err = cursor.All(context.Background(), &tags)
	return
}

// PostTag 修改每个博客的tag
func (service *TagService) PostTag(id string, tag model.Tag) (ID primitive.ObjectID, err error) {
	tag.BlogID = primitive.ObjectIDFromHex(id)
	_, err := service.DB.InsertOne(context.Background(), &tag)
	return tag.BlogID, err;
}

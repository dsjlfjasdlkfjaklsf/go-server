package service

import (
	"context"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CommentService 评论服务
type CommentService struct {
	DB *mongo.Collection
}

// GetCommentByBlogID 根据BlogID获得评论
func (service *CommentService) GetCommentByBlogID(id string) (comment model.Comment, err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = service.DB.DeleteOne(context.Background(), bson.D{{"_id", ID}})
	return
}

// PostComment 发评论
func (service *CommentService) PostComment(commentID string, comment model.Comment) (id primitive.ObjectID, err error) {
	return
}

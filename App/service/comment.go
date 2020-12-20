package service

import (
	"context"
	"time"

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
func (service *CommentService) GetCommentByBlogID(id string) (comments []model.Comment, err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	cursor, err := service.DB.Find(context.Background(), bson.D{{"BlogId", ID}})
	if err != nil {
		return
	}
	err = cursor.All(context.Background(), &comments)
	return
}

// PostComment 发评论
func (service *CommentService) PostComment(commentID string, comment model.Comment) (id primitive.ObjectID, err error) {
	comment.BlogID, err = primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return
	}
	comment.CommentTime = time.Now().Unix()
	_, err = service.DB.InsertOne(context.Background(), &comment)
	return comment.BlogID, err
}

package service

import (
	"context"
	"errors"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DB *mongo.Collection
}

func (service *UserService) CreateUser(id string, name string, password string) (ID primitive.ObjectID, err error) {
	user := model.UserInfo{}
	user.ID, err = primitive.ObjectIDFromHex(id)
	user.Name = name
	user.Password = password
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", ID}}).Err()
	if err != mongo.ErrNoDocuments {
		return user.ID, errors.New("The userID already exists.")
	}
	_, err = service.DB.InsertOne(context.Background(), &user)
	return user.ID, err
}

func (service *UserService) GetUserByID(id string) (user model.UserInfo, err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", ID}}).Decode(&user)
	return
}

func (service *UserService) LoginUser(id string, password string) (ID primitive.ObjectID, err error) {
	ID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	var user model.UserInfo
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", ID}}).Decode(&user)
	if err == nil && user.Password != password {
		err = errors.New("The password is incorrect.")
	}
	return
}

package service

import (
	"context"
	"errors"

	"github.com/dsjlfjasdlkfjaklsf/go-server/App/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DB *mongo.Collection
}

func (service *UserService) CreateUser(id string, name string, password string) (ID string, err error) {
	user := model.UserInfo{}
	user.ID = id
	user.Name = name
	user.Password = password
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", id}}).Err()
	if err != mongo.ErrNoDocuments {
		return user.ID, errors.New("The userID already exists.")
	}
	_, err = service.DB.InsertOne(context.Background(), &user)
	return user.ID, err
}

func (service *UserService) GetUserByID(id string) (user model.UserInfo, err error) {
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", id}}).Decode(&user)
	return
}

func (service *UserService) LoginUser(id string, password string) (name string, err error) {
	var user model.UserInfo
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", id}}).Decode(&user)
	if err != nil {
		err = errors.New("No such userID.")
	} else if user.Password != password {
		err = errors.New("The password is incorrect.")
	}
	name = user.Name
	return
}

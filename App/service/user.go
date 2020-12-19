package service

import (
	mgo "go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DB *mgo.Collection
}

func (service *UserService) CreateUser() {

}

func (service *UserService) GetUserByID() {

}

func (service *UserService) LoginUser() {

}

func (service *UserService) LogoutUser() {

}

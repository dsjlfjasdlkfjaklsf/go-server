package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo 数据库配置
type Mongo struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
}

type Service struct {
	Config  Mongo
	DB      *mongo.Database
	User    UserService
	Tag     TagService
	Blog    BlogService
	Comment CommentService
}

// InitMongo 初始化数据库
func (service *Service) InitMongo(conf Mongo) error {
	service.Config = conf
	if service.DB != nil {
		service.DB.Client().Disconnect(context.TODO())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(
			"mongodb+srv://"+
				conf.User+
				":"+conf.Password+
				"@"+conf.Host+
				"/"+conf.Name+"?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	service.DB = client.Database(conf.Name)
	service.User.DB = service.DB.Collection("users")
	service.Tag.DB = service.DB.Collection("tags")
	service.Blog.DB = service.DB.Collection("blogs")
	service.Comment.DB = service.DB.Collection("comments")
	log.Printf("MongoDB Connect Success!")
	return nil
}

// NewModel 初始化Service
func NewService(c Mongo) (*Service, error) {
	service := new(Service)
	err := service.InitMongo(c)
	return service, err
}

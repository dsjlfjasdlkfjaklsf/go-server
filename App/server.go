package app

import (
	"log"
	"net/http"

	sw "github.com/dsjlfjasdlkfjaklsf/go-server/App/controller"
	"github.com/dsjlfjasdlkfjaklsf/go-server/App/service"
)

// Config 配置文件
type Config struct {
	Mongo  service.Mongo `yaml:"Mongo"`  // mongoDB配置
	Server ServerConfig  `yaml:"Server"` // iris配置
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host string `yaml:"Host"` // 服务器监听地址
	Port string `yaml:"Port"` // 服务器监听端口
}

// RunServer 开始运行服务
func RunServer(c Config) {
	// 初始化数据库
	// 初始化服务
	Service, err := service.NewService(c.Mongo)
	if err != nil {
		log.Fatal(err)
	}
	// 启动服务器
	log.Printf("Server started at " + c.Server.Host + ":" + c.Server.Port)
	router := sw.NewRouter(Service)
	log.Fatal(http.ListenAndServe(c.Server.Host+":"+c.Server.Port, router))
	// ":8080"
}

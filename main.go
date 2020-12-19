package main

import (
	"io/ioutil"
	"log"

	app "github.com/dsjlfjasdlkfjaklsf/go-server/App"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	// 加载配置文件
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("Can't find the config file in config/config.yaml")
		return
	}
	log.Printf("Load the config file in config/config.yaml")
	conf := app.Config{}
	yaml.Unmarshal(data, &conf)
	app.RunServer(conf)
}

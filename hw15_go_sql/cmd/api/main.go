package main

import (
	"flag"
	"log"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/api"
	"github.com/BurntSushi/toml"
)

var (
	configPath string = "config/api.toml"
)

func init() {
	// get enviroment
	flag.StringVar(&configPath, "path", "config/api.toml", "path config fil in .toml format")
}
func main() {
	log.Println("first step")
	//server instance initialization
	// log, err := logger.Newlogger()
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find config file", err)
	}
	server, err := api.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}

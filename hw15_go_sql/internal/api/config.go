package api

import "github.com/AlexSH61/homework_basic/hw15_go_sql/internal/store"

// general instance for api server of rest
type Config struct {
	BindAddr    string `toml:bind_addr`
	LoggerLevel string `toml:"logger_level"`
	Store       *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Store:       store.NewConfig(),
	}
}

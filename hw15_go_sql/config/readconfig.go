package config

import (
	"encoding/json"
	"os"
)

type Setting struct {
	ServerHost string `json:"ServerHost"`
	ServerPort string `json:"ServerPort"`
	PgHost     string `json:"PgHost"`
	PgPort     string `json:"PgPort"`
	PgUser     string `json:"PgUser"`
	PgPass     string `json:"PgPass"`
	PgNameBD   string `json:"PgNameBD"`
}

func ReadConfig(file string) (*Setting, error) {
	set, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var setting Setting
	err = json.Unmarshal(set, &setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

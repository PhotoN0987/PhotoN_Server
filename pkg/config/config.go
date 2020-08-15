package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config config
type Config struct {
	Server struct {
		Port    string
		Timeout int
	}
	DataBase struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
	AwsCredentials struct {
		AwsSecretAccessKey string
		AwsAccessKey       string
	}
}

// NewConfig create config
func NewConfig() *Config {

	jsonString, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	c := new(Config)

	// UnmarshalしてConfigにマッピング
	err = json.Unmarshal(jsonString, c)
	if err != nil {
		panic(err)
	}

	return c
}

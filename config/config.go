package config

import (
	"encoding/json"
	"io/ioutil"
)

type configClass struct {
	Episode     string
	Environment struct {
		Type string
		Port string
	}
}

type Config struct {
	config *configClass
}

func (cfg *Config) SetConfig() {
	fileContent, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(fileContent, &cfg.config); err != nil {
		panic(err)
	}
}

func (cfg *Config) GetConfig() *configClass {
	if cfg.config == nil {
		cfg.SetConfig()
	}

	return cfg.config
}

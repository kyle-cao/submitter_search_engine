package config

import (
	"encoding/json"
	"os"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) get() (jsonMap map[string]interface{}, err error) {

	pwd, _ := os.Getwd()
	content, err := os.ReadFile(pwd + "/config/config.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &jsonMap)
	if err != nil {
		return nil, err
	}

	return
}

func (c *Config) GetSubmitConfig(engine []string) (config map[string]interface{}, err error) {

	config, err = c.get()
	if err != nil {
		return
	}
	if len(engine) == 0 || len(engine) == 1 && engine[0] == "" {
		return
	}

	engineMap := make(map[string]struct{}, len(engine))
	for _, k := range engine {
		engineMap[k] = struct{}{}
	}
	for k := range config {
		if _, exists := engineMap[k]; !exists {
			delete(config, k)
		}
	}

	return
}

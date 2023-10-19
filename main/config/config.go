package config

import (
	"os"
	"sync/atomic"

	"gopkg.in/yaml.v3"
)

type Path struct {
	Name   string `yaml:"name"`
	Prefix string `yaml:"prefix"`
	Url    string `yaml:"url"`
}

type Config struct {
	ProxyPaths []*Path `yaml:"proxy_paths"`
}

const configFileName = "config.yaml"

var configStore atomic.Value

func InitConfig() error {
	data, err := os.ReadFile(configFileName)
	if err != nil {
		return err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}

	configStore.Store(config)

	return nil
}

func GetConfig() *Config {
	return configStore.Load().(*Config)
}

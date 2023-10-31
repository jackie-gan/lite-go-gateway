package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Path struct {
	Name   string `yaml:"name"`
	Prefix string `yaml:"prefix"`
	Url    string `yaml:"url"`
}

type Config struct {
	ProxyPaths []*Path `yaml:"proxy_paths"`
	Log        struct {
		Path string `yaml:"path"`
		Age  int32  `yaml:"age"`
	} `yaml:"log"`
}

const configFileName = "config.yaml"

var C *Config = &Config{}

func InitConfig() error {
	data, err := os.ReadFile(configFileName)
	if err != nil {
		fmt.Println("Read Config File Fail", err)
		return err
	}

	err = yaml.Unmarshal(data, C)
	if err != nil {
		return err
	}

	fmt.Print(C)

	return nil
}

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Paths struct {
		Models       string `yaml:"models"`
		Repositories string `yaml:"repositories"`
		Services     string `yaml:"services"`
		Handlers     string `yaml:"handlers"`
	} `yaml:"paths"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Address string        `yaml:"address"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"server"`
	ExternalAPI struct {
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"external_api"`
}

func LoadConfig(filename string) (*Config, error) {
	var cfg Config
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

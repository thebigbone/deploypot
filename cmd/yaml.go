package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Name      string   `yaml:"name"`
		Repo_URL  string   `yaml:"repo_url"`
		Directory string   `yaml:"directory"`
		Language  string   `yaml:"language"`
		Domain    string   `yaml:"domain"`
		Proxy     string   `yaml:"proxy"`
		Arguments []string `yaml:"arguments"`
	} `yaml:"app"`
}

func parseConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
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

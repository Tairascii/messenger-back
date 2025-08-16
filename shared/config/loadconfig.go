package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	configEnvKey = "CONFIG_FILE"
)

func LoadConfig[Config any]() (*Config, error) {
	f, err := os.Open(os.Getenv(configEnvKey))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

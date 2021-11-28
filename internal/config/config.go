package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	// DefaultPath - default path for config.
	DefaultPath = "./cmd/config.yaml"
)

// Config defines the properties of the application configuration.
type Config struct {
	PINGenerationKey        string            `yaml:"pin-generation-key"`
	Decimalization          map[string]string `yaml:"decimalization"`
	PinVerificationKeyIndex string            `yaml:"pin-verification-key-index"`
}

func New(cfgFilePath string) (*Config, error) {
	data, err := os.ReadFile(cfgFilePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

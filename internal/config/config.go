package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Lowercase bool `yaml:"lowercase"`
	English   bool `yaml:"english"`
	Symbols   bool `yaml:"symbols"`
	Sensitive bool `yaml:"sensitive"`
}

func Default() Config {
	return Config{
		Lowercase: true,
		English:   true,
		Symbols:   true,
		Sensitive: true,
	}
}

func Load(path string) (Config, error) {
	cfg := Default()
	if path == "" {
		return cfg, nil
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

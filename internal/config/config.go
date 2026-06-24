package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Target struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Config struct {
	Targets        []Target `yaml:"targets"`
	TimeoutSeconds int      `yaml:"timeout_seconds"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	if len(cfg.Targets) == 0 {
		return nil, fmt.Errorf("config has no targets")
	}
	if cfg.TimeoutSeconds <= 0 {
		cfg.TimeoutSeconds = 10
	}

	return &cfg, nil
}

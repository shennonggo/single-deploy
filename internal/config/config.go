package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Validate() error {
	if len(c.Projects) == 0 {
		return fmt.Errorf("no projects defined in config")
	}

	for _, p := range c.Projects {
		if p.Name == "" {
			return fmt.Errorf("project name cannot be empty")
		}
		if p.Path == "" {
			return fmt.Errorf("project path cannot be empty for %s", p.Name)
		}
	}

	return nil
}

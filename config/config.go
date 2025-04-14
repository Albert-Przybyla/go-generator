package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML config: %w", err)
	}

	if len(os.Args) < 2 {
		return nil, fmt.Errorf("missing entity name (expected as first argument)")
	}

	rawName := os.Args[1]
	cfg.StructName = toPascalCase(rawName)
	cfg.VarName = toCamelCase(rawName)
	cfg.FileName = toSnakeCase(rawName)

	repoName, err := getModuleName()
	if err != nil {
		return nil, err
	}
	cfg.RepoName = repoName

	return &cfg, nil
}

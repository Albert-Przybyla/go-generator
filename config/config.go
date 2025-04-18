package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(path string) (*Config, error) {
	simpleFlag := flag.Bool("simple", false, "use simple mode")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return nil, fmt.Errorf("missing entity name (expected as first argument)")
	}
	rawName := args[0]

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML config: %w", err)
	}

	cfg.PascalCase = toPascalCase(rawName)
	cfg.CamelCase = toCamelCase(rawName)
	cfg.SnakeCase = toSnakeCase(rawName)
	cfg.KebabCase = toKebabCase(rawName)

	repoName, err := getModuleName()
	if err != nil {
		return nil, err
	}
	cfg.RepoName = repoName

	cfg.Simple = *simpleFlag

	return &cfg, nil
}

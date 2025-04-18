package config

type Config struct {
	SnakeCase  string
	KebabCase  string
	CamelCase  string
	PascalCase string

	Simple   bool
	RepoName string
	Paths    struct {
		Models       string `yaml:"models"`
		Repositories string `yaml:"repositories"`
		Services     string `yaml:"services"`
		Handlers     string `yaml:"handlers"`
		Mappers      string `yaml:"mappers"`
		Dtos         string `yaml:"dtos"`
	} `yaml:"paths"`
}

package config

type Config struct {
	StructName string
	VarName    string
	RepoName   string
	FileName   string
	Paths      struct {
		Models       string `yaml:"models"`
		Repositories string `yaml:"repositories"`
		Services     string `yaml:"services"`
		Handlers     string `yaml:"handlers"`
	} `yaml:"paths"`
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Albert-Przybyla/go-generator/config"
)

func main() {
	cfg, err := config.LoadConfig(".conf.yaml")
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	files := map[string]struct {
		Template string
		Dir      string
		Ext      string
	}{
		// "model":      {"templates/model.tmpl", cfg.Paths.Models, ".go"},
		"repository": {"templates/repository.tmpl", cfg.Paths.Repositories, ".go"},
		"service":    {"templates/service.tmpl", cfg.Paths.Services, ".go"},
		// "handler":    {"templates/handler.tmpl", cfg.Paths.Handlers, ".go"},
		// "model_test": {"templates/model_test.tmpl", cfg.Paths.Models, "_test.go"},
	}

	if err := generateFiles(files, cfg); err != nil {
		log.Fatalf("Template generation error: %v", err)
	}

	fmt.Println("Files generated successfully.")
}

func generateFiles(files map[string]struct {
	Template string
	Dir      string
	Ext      string
}, data *config.Config) error {
	for suffix, file := range files {
		if err := os.MkdirAll(file.Dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create dir %s: %w", file.Dir, err)
		}

		filename := fmt.Sprintf("%s_%s%s", data.FileName, suffix, file.Ext)
		outPath := filepath.Join(file.Dir, filename)

		if err := renderTemplateToFile(file.Template, outPath, data); err != nil {
			return err
		}
	}
	return nil
}
func renderTemplateToFile(tmplPath, outPath string, data *config.Config) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", tmplPath, err)
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outPath, err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

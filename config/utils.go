package config

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getModuleName() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

func toPascalCase(s string) string {
	parts := splitWords(s)
	for i, word := range parts {
		parts[i] = capitalize(word)
	}
	return strings.Join(parts, "")
}

func toCamelCase(s string) string {
	parts := splitWords(s)
	for i, word := range parts {
		if i == 0 {
			parts[i] = strings.ToLower(word)
		} else {
			parts[i] = capitalize(word)
		}
	}
	return strings.Join(parts, "")
}

func splitWords(s string) []string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	return strings.Fields(s)
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func toSnakeCase(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.ToLower(str)
	words := strings.Fields(str)
	return strings.Join(words, "_")
}

func toKebabCase(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.ToLower(str)
	words := strings.Fields(str)
	return strings.Join(words, "-")
}

package config

import (
	"encoding/json"
	"os"
)

// Global map to hold the redirects
var TestNumberRedirects map[string][]string

func LoadConfig(filePath string) error {
	// 1. Read the file
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 2. Unmarshal the JSON into our map
	err = json.Unmarshal(fileBytes, &TestNumberRedirects)
	if err != nil {
		return err
	}

	return nil
}

func GetTestNumberRedirects(number string) []string {
	return TestNumberRedirects[number]
}

package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func LoadConfig(credentialsPath string) (*oauth2.Config, error) {
	b, err := os.ReadFile(credentialsPath)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		return nil, err
	}

	return config, nil
}

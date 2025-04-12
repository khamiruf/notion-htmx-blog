package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	NotionAPIKey   string
	NotionDBID     string
	TemplatesPath  string
	StaticFilePath string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	notionAPIKey := os.Getenv("NOTION_API_KEY")
	if notionAPIKey == "" {
		return nil, fmt.Errorf("NOTION_API_KEY is required")
	}

	notionDBID := os.Getenv("NOTION_DATABASE_ID")
	if notionDBID == "" {
		return nil, fmt.Errorf("NOTION_DATABASE_ID is required")
	}

	return &Config{
		Port:           port,
		NotionAPIKey:   notionAPIKey,
		NotionDBID:     notionDBID,
		TemplatesPath:  "web/templates",
		StaticFilePath: "web/static",
	}, nil
}

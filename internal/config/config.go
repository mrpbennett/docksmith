package config

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
})

type Config struct {
	GitHubWebhookSecret string
	RegistryURL         string
	RegistryUser        string
	RegistryPassword    string
	RepoDir             string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Failed to load environment variables", "error", err)
	}

	cfg := &Config{
		GitHubWebhookSecret: os.Getenv("GITHUB_WEBHOOK_SECRET"),
		RegistryURL:         os.Getenv("REGISTRY_URL"),
		RegistryUser:        os.Getenv("REGISTRY_USER"),
		RegistryPassword:    os.Getenv("REGISTRY_PASSWORD"),
		RepoDir:             os.Getenv("REPO_DIR"),
	}

	if cfg.GitHubWebhookSecret == "" {
		logger.Error("GitHub webhook secret is not set")
	}

	return cfg
}

package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	APIKey      string `env:"API_KEY"`
	URL         string `env:"URL"`
	GithubToken string `env:"GITHUB_TOKEN"`
}

func Load() *Config {
	var (
		cfg  Config
		path = ".env"
	)

	if err := godotenv.Load(path); err != nil && !os.IsNotExist(err) {
		log.Fatalf("failed loading .env file: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed parsing .env file: %v", err)
	}

	return &cfg
}

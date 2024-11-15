package config

import "os"

type Mode string

const (
	DEVELOPMENT Mode = "DEVELOPMENT"
	PRODUCTION  Mode = "PRODUCTION"
)

type Config struct {
	APIKey     string
	WeatherURL string
}

func (c *Config) Load() error {

	c.APIKey = os.Getenv("WEATHER_API_KEY")
	c.WeatherURL = os.Getenv("WEATHER_URL")

	return nil
}

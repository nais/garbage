package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Project             string `envconfig:"GOOGLE_PROJECT_ID"`
	Location            string `envconfig:"GOOGLE_LOCATION"`
	ServiceAccountEmail string `envconfig:"SERVICE_ACCOUNT_EMAIL"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	return cfg, err
}

package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppID          int    `env:"APP_ID"`
	InstallationID int    `env:"INSTALLATION_ID"`
	PrivateKey     string `env:"PRIVATE_KEY"`
	GitHubToken    string `env:"GITHUB_TOKEN"`
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	return &Config{
		AppID:          viper.GetInt("APP_ID"),
		InstallationID: viper.GetInt("INSTALLATION_ID"),
		PrivateKey:     viper.GetString("PRIVATE_KEY"),
		GitHubToken:    viper.GetString("GITHUB_TOKEN"),
	}
}

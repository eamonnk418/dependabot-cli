package app

import (
	"github.com/eamonnk418/dependabot-cli/internal/config"
	"github.com/eamonnk418/dependabot-cli/internal/log"
)

type App struct {
	Config *config.Config
	Logger *log.Logger
}

func NewApplication(cfg *config.Config, log *log.Logger) *App {
	return &App{Config: cfg, Logger: log}
}

package main

import (
	"context"

	"github.com/eamonnk418/dependabot-cli/internal/app"
	"github.com/eamonnk418/dependabot-cli/internal/config"
	"github.com/eamonnk418/dependabot-cli/internal/log"
	"github.com/eamonnk418/dependabot-cli/pkg/cmd"
)

func main() {
	cfg := config.LoadConfig()
	logger := log.NewLogger()
	app := app.NewApplication(cfg, logger)
	ctx := context.WithValue(context.Background(), "app", app)
	cmd.Execute(ctx)
}

package main

import (
	"context"
	"time"

	"github.com/eamonnk418/dependabot-cli/pkg/cmd"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	cmd.Execute(ctx)
}
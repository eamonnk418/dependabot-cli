package config

import (
	"fmt"

	"github.com/eamonnk418/dependabot-cli/internal/app"
	"github.com/spf13/cobra"
)

type createOptions struct {
	dryRun bool
}

func NewCmdGenerate() *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Dependabot configuration file for a repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, ok := cmd.Context().Value("app").(*app.App)
			if !ok {
				return fmt.Errorf("failed to get app from context")
			}

			repo, err := cmd.Flags().GetString("repo")
			if err != nil {
				return fmt.Errorf("failed to get repository flag: %w", err)
			}

			if opts.dryRun {
				app.Logger.InfoContext(cmd.Context(), "Creating dependabot.yml file", "repo", repo)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.dryRun, "dry-run", "d", false, "Print the configuration file to the console instead of writing it to a file")

	return cmd
}

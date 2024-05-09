package repo

import (
	"fmt"

	"github.com/eamonnk418/dependabot-cli/internal/app"
	"github.com/spf13/cobra"
)

func NewCmdRepo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Manage repositories",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, ok := cmd.Context().Value("app").(*app.App)
			if !ok {
				return fmt.Errorf("could not get app from context")
			}

			app.Logger.InfoContext(cmd.Context(), "Listing repositories", "token", app.Config.GitHubToken)

			return nil
		},
	}

	return cmd
}

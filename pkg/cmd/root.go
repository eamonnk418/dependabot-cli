package cmd

import (
	"context"

	"github.com/MakeNowJust/heredoc"
	"github.com/eamonnk418/dependabot-cli/pkg/cmd/repo"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dependabot-cli",
		Short: "A CLI for generating dependabot configuration files",
		Long: heredoc.Doc(`
			This CLI tool helps you generate Dependabot configuration files easily.
			Dependabot is a tool that automates security updates and other dependency updates for your projects.
			With this tool, you can quickly create configuration files for Dependabot to use.
		`),
	}

	cmd.AddCommand(repo.NewCmdRepo())

	return cmd
}

func Execute(ctx context.Context) {
	cobra.CheckErr(NewCmdRoot().ExecuteContext(ctx))
}

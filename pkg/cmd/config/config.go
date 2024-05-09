package config

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

type configOptions struct {
	repo string
}

func NewCmdConfig() *cobra.Command {
	var opts configOptions

	cmd := &cobra.Command{
		Use:   "config",
		Short: "A CLI for generating dependabot configuration files",
		Long: heredoc.Doc(`
			This CLI tool helps you generate Dependabot configuration files easily. 
			Dependabot is a tool that automates security updates and other dependency updates for your projects. 
			With this tool, you can quickly create configuration files for Dependabot to use.
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewCmdGenerate())

	cmd.PersistentFlags().StringVarP(&opts.repo, "repo", "r", "", "The repository to generate a configuration file for")

	return cmd
}

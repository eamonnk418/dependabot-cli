package cmd

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/eamonnk418/dependabot-cli/internal/dependabot"
	"github.com/eamonnk418/dependabot-cli/internal/dependabot/factory"
	"github.com/eamonnk418/dependabot-cli/internal/dependabot/model"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create Dependabot template factory
			tmplFactory, err := factory.NewTemplateFactory("npm", &model.DependabotConfig{
				Version: 2,
				Updates: []model.DependencyUpdate{
					{
						PackageEcosystem: "npm",
						Directory:        "/",
						Schedule: model.UpdateSchedule{
							Interval: "daily",
						},
					},
				},
			})
			if err != nil {
				return err
			}

			// Create Dependabot service
			svc := dependabot.NewDependabotService(tmplFactory)

			// Create Dependabot template
			tmpl, err := svc.RenderTemplate()
			if err != nil {
				return err
			}

			// Print the template
			fmt.Println(tmpl)

			return nil
		},
	}

	return cmd
}

func Execute(ctx context.Context) {
	cobra.CheckErr(NewCmdRoot().ExecuteContext(ctx))
}

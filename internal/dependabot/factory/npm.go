package factory

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/eamonnk418/dependabot-cli/internal/dependabot/model"
)

//go:embed template/npm.tmpl
var npmTemplate string

// NpmTemplateFactory is a factory for creating npm Dependabot templates.
type NpmTemplateFactory struct {
	Dependabot *model.DependabotConfig
}

// NewNpmTemplateFactory creates a new instance of NpmTemplateFactory.
func NewNpmTemplateFactory(dependabot *model.DependabotConfig) *NpmTemplateFactory {
	return &NpmTemplateFactory{
		Dependabot: dependabot,
	}
}

// CreateTemplate creates an npm Dependabot template.
func (f *NpmTemplateFactory) CreateTemplate() (string, error) {
	tmpl := template.Must(template.New("npm").Parse(npmTemplate))

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "npm", f.Dependabot); err != nil {
		return "", fmt.Errorf("failed to execute npm template: %w", err)
	}

	return buf.String(), nil
}

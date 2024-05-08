package factory

import (
	"errors"

	"github.com/eamonnk418/dependabot-cli/internal/dependabot/model"
)

var (
	ErrUnsupportedPackageEcosystem = errors.New("unsupported package ecosystem")
)

type TemplateFactory interface {
	CreateTemplate() (string, error)
}

func NewTemplateFactory(packageEcosystem string, dependabot *model.DependabotConfig) (TemplateFactory, error) {
	switch packageEcosystem {
	case "npm":
		return NewNpmTemplateFactory(dependabot), nil
	default:
		return nil, ErrUnsupportedPackageEcosystem
	}
}

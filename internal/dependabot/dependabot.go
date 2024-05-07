package dependabot

import "github.com/eamonnk418/dependabot-cli/internal/dependabot/factory"

type Service struct {
	TemplateFactory factory.TemplateFactory
}

func NewDependabotService(f factory.TemplateFactory) *Service {
	return &Service{
		TemplateFactory: f,
	}
}

func (s *Service) RenderTemplate() (string, error) {
	return s.TemplateFactory.CreateTemplate()	
}

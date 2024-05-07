package factory

type PipTemplateFactory struct{}

func NewPipTemplateFactory() *PipTemplateFactory {
	return &PipTemplateFactory{}
}

func (f *PipTemplateFactory) CreateTemplate() (string, error) {
	return "pip template", nil
}

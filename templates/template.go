package templates

type TemplateType string

type Template interface {
	Type() TemplateType
}

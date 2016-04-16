package templates

const TemplateTypeButton TemplateType = "button"

// ButtonTemplate is a template that
type ButtonTemplate struct {
	Text    string   `json:"text,omitempty"`
	Buttons []Button `json:"buttons,omitempty"`
}

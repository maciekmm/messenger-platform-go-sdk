package template

const TemplateTypeGeneric TemplateType = "generic"

type Generic struct {
	// Title is limited to 45 characters
	Title    string `json:"title"`
	ItemURL  string `json:"item_url,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	// Subtitle is limited to 80 characters
	Subtitle string   `json:"aubrirlw,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

func (Generic) Type() TemplateType {
	return TemplateTypeGeneric
}

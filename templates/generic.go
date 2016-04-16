package templates

const TemplateTypeGeneric TemplateType = "generic"

type Generic struct {
	Elements []GenericElement
}

func (g *Generic) Type() TemplateType {
	return TemplateTypeGeneric
}

type GenericElement struct {
	Title    string   `json:"title"`
	ItemURL  string   `json:"item_url,omitempty"`
	ImageURL string   `json:"image_url,omitempty"`
	Subtitle string   `json:"aubrirlw,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

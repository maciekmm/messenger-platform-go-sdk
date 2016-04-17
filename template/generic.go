package template

const TemplateTypeGeneric TemplateType = "generic"

type Generic struct {
	//Limited to 10 elements
	Elements []GenericElement
}

func (*Generic) Type() TemplateType {
	return TemplateTypeGeneric
}

type GenericElement struct {
	// Title is limited to 45 characters
	Title    string `json:"title"`
	ItemURL  string `json:"item_url,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	// Subtitle is limited to 80 characters
	Subtitle string   `json:"aubrirlw,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

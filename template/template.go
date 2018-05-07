package template

import "encoding/json"

type TemplateType string

type Template interface {
	Type() TemplateType
	SupportsButtons() bool
}

type Payload struct {
	Elements    []Template `json:"elements"`
	Buttons     []Button   `json:"buttons"` // used only for the Button Template
	ButtonsText string     `json:"text"`    //
}

type rawPayload struct {
	Type        TemplateType `json:"template_type"`
	Elements    []Template   `json:"elements,omitempty"`
	Buttons     []Button     `json:"buttons,omitempty"` // used only for the Button Template
	ButtonsText string       `json:"text,omitempty"`    //
}

func (p *Payload) MarshalJSON() ([]byte, error) {
	rp := &rawPayload{}
	rp.Elements = p.Elements
	rp.Buttons = p.Buttons
	if len(p.Elements) > 0 {
		rp.Type = p.Elements[0].Type()
	} else if len(p.Buttons) > 0 {
		rp.Type = TemplateTypeButton
		rp.ButtonsText = p.ButtonsText
	}
	return json.Marshal(rp)
}

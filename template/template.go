package template

import (
	"encoding/json"
	"errors"
)

type TemplateType string

type Template interface {
	Type() TemplateType
	SupportsButtons() bool
}

type Payload struct {
	Elements []Template `json:"elements"`
}

type rawPayload struct {
	Type     TemplateType `json:"template_type"`
	Elements []Template   `json:"elements"`
}

func (p *Payload) MarshalJSON() ([]byte, error) {
	rp := &rawPayload{}
	if len(p.Elements) < 1 {
		return []byte{}, errors.New("Elements slice cannot be empty")
	}
	rp.Elements = p.Elements
	rp.Type = p.Elements[0].Type()
	return json.Marshal(rp)
}

package messenger

import (
	"encoding/json"
)

type AttachmentType string

const (
	AttachmentTypeTemplate AttachmentType = "template"
	AttachmentTypeImage    AttachmentType = "image"
	AttachmentTypeVideo    AttachmentType = "video"
	AttachmentTypeAudio    AttachmentType = "audio"
	AttachmentTypeLocation AttachmentType = "location"
)

type Attachment struct {
	Type    AttachmentType `json:"type"`
	Payload interface{}    `json:"payload,omitempty"`
}

type rawAttachment struct {
	Type    AttachmentType  `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func (a *Attachment) UnmarshalJSON(b []byte) error {
	raw := &rawAttachment{}
	err := json.Unmarshal(b, raw)
	if err != nil {
		return err
	}
	a.Type = raw.Type
	var payload interface{}

	switch a.Type {
	case AttachmentTypeLocation:
		payload = &Location{}
	case AttachmentTypeTemplate:
		//TODO: implement template unmarshalling
		a.Payload = raw.Payload
		return nil
	default:
		payload = &Resource{}
	}

	err = json.Unmarshal(raw.Payload, payload)
	if err != nil {
		return err
	}
	a.Payload = payload
	return nil
}

type Resource struct {
	URL      string `json:"url"`
	Reusable bool   `json:"is_reusable,omitempty"`
}

type ReusableAttachment struct {
	AttachmentID string `json:"attachment_id"`
}

type Location struct {
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

package messenger

type AttachmentType string

const (
	AttachmentTypeTemplate AttachmentType = "template"
	AttachmentTypeImage    AttachmentType = "image"
	AttachmentTypeVideo    AttachmentType = "video"
)

type Attachment struct {
	Type    AttachmentType
	Payload interface{}
}

// func (a *Attachment) MarshalJSON() ([]byte, error) {
// 	att := &rawAttachment{
// 		Payload: a.Payload,
// 	}
// 	switch a.Payload.(type) {
// 	case template.Payload:
// 		att.Type = AttachmentTypeTemplate
// 	case Image:
// 		att.Type = AttachmentTypeImage
// 	default:
// 		return []byte{}, errors.New("Invalid payload")
// 	}
// 	return json.Marshal(att)
// }

type Resource struct {
	URL string `json:"url"`
}

package messenger

import (
	"errors"

	"github.com/maciekmm/messenger-platform-go-sdk/template"
)

type ContentType string

const (
	ContentTypeText     ContentType = "text"
	ContentTypeLocation ContentType = "location"
)

type SendMessage struct {
	Text         string       `json:"text,omitempty"`
	Attachment   *Attachment  `json:"attachment,omitempty"`
	QuickReplies []QuickReply `json:"quick_replies,omitempty"`
	Metadata     string       `json:"metadata,omitempty"`
}

type QuickReply struct {
	ContentType ContentType `json:"content_type"`
	Title       string      `json:"title,omitempty"`
	Payload     string      `json:"payload,omitempty"`
	ImageURL    string      `json:"image_url,omitempty"`
}

// Recipient describes the person who will receive the message
// Either ID or PhoneNumber has to be set
type Recipient struct {
	ID          string `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

// NotificationType describes the behavior phone will execute after receiving the message
type NotificationType string

const (
	// NotificationTypeRegular will emit a sound/vibration and a phone notification
	NotificationTypeRegular NotificationType = "REGULAR"
	// NotificationTypeSilentPush will just emit a phone notification
	NotificationTypeSilentPush NotificationType = "SILENT_PUSH"
	// NotificationTypeNoPush will not emit sound/vibration nor a phone notification
	NotificationTypeNoPush NotificationType = "NO_PUSH"
)

type MessageQuery struct {
	Recipient        Recipient        `json:"recipient"`
	Message          SendMessage      `json:"message"`
	NotificationType NotificationType `json:"notification_type,omitempty"`
}

func (mq *MessageQuery) RecipientID(recipientID string) error {
	if mq.Recipient.PhoneNumber != "" {
		return errors.New("Only one user identification (phone or id) can be specified.")
	}
	mq.Recipient.ID = recipientID
	return nil
}

func (mq *MessageQuery) RecipientPhoneNumber(phoneNumber string) error {
	if mq.Recipient.ID != "" {
		return errors.New("Only one user identification (phone or id) can be specified.")
	}
	mq.Recipient.PhoneNumber = phoneNumber
	return nil
}

func (mq *MessageQuery) Notification(notification NotificationType) *MessageQuery {
	mq.NotificationType = notification
	return mq
}

func (mq *MessageQuery) Text(text string) error {
	if mq.Message.Attachment != nil && mq.Message.Attachment.Type == AttachmentTypeTemplate {
		return errors.New("Can't set both text and template.")
	}
	mq.Message.Text = text
	return nil
}

func (mq *MessageQuery) resource(typ AttachmentType, url string) error {
	if mq.Message.Attachment == nil {
		mq.Message.Attachment = &Attachment{}
	}
	if mq.Message.Attachment.Payload != nil {
		return errors.New("Attachment already specified.")
	}
	mq.Message.Attachment.Type = typ
	mq.Message.Attachment.Payload = &Resource{URL: url}
	return nil
}

func (mq *MessageQuery) Audio(url string) error {
	return mq.resource(AttachmentTypeAudio, url)
}

func (mq *MessageQuery) Video(url string) error {
	return mq.resource(AttachmentTypeVideo, url)
}

func (mq *MessageQuery) Image(url string) error {
	return mq.resource(AttachmentTypeImage, url)
}

func (mq *MessageQuery) Template(tpl template.Template) error {
	if mq.Message.Attachment == nil {
		mq.Message.Attachment = &Attachment{}
	}
	if mq.Message.Attachment.Type != AttachmentTypeTemplate && mq.Message.Attachment.Payload != nil {
		return errors.New("Non-template attachment already specified.")
	}

	if mq.Message.Attachment.Payload == nil {
		mq.Message.Attachment.Type = AttachmentTypeTemplate
		mq.Message.Attachment.Payload = &template.Payload{}
	}

	payload := mq.Message.Attachment.Payload.(*template.Payload)

	for _, v := range payload.Elements {
		if v.Type() != tpl.Type() {
			return errors.New("All templates have to have thesame type.")
		}
	}

	payload.Elements = append(payload.Elements, tpl)
	return nil
}

// Documentation: https://developers.facebook.com/docs/messenger-platform/send-api-reference/quick-replies#quick_reply
func (mq *MessageQuery) QuickReply(qr QuickReply) error {
	//max 10 quick replies
	if len(mq.Message.QuickReplies) >= 10 {
		return errors.New("Number of quick replies is limited to 10")
	}

	//if location then no title or payload
	if qr.ContentType == ContentTypeLocation && (len(qr.Title) != 0 || len(qr.Payload) != 0) {
		return errors.New("Location ContentType supports neither title nor payload")
	}

	//title has a 20 character limit
	if len(qr.Title) > 20 {
		return errors.New("Title is too long, it has a 20 character limit.")
	}
	//payload has a 1000 character limit
	if len(qr.Payload) > 1000 {
		return errors.New("Payload is too long, it has a 1000 character limit.")
	}

	if len(qr.ContentType) == 0 {
		qr.ContentType = ContentTypeText
	}

	mq.Message.QuickReplies = append(mq.Message.QuickReplies, qr)
	return nil
}

func (mq *MessageQuery) Metadata(metadata string) error {
	if len(metadata) > 1000 {
		return errors.New("Metadata is too long, it has a 1000 character limit.")
	}
	mq.Message.Metadata = metadata
	return nil
}

package messenger

import "encoding/json"

type upstreamEvent struct {
	Object  string          `json:"object"`
	Entries []*MessageEvent `json:"entry"`
}

type Event struct {
	ID   json.Number `json:"id"`
	Time int64       `json:"time"`
}

type MessageOpts struct {
	Sender struct {
		ID string `json:"id"`
	} `json:"sender"`
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Timestamp int64 `json:"timestamp"`
}

type MessageEvent struct {
	Event
	Messaging []struct {
		MessageOpts
		Message  *MessageEcho `json:"message,omitempty"`
		Delivery *Delivery    `json:"delivery,omitempty"`
		Postback *Postback    `json:"postback,omitempty"`
		Optin    *Optin       `json:"optin,empty"`
		Read     *Read        `json:"read,omitempty"`
	} `json:"messaging"`
}

type ReceivedMessage struct {
	ID          string             `json:"mid"`
	Text        string             `json:"text,omitempty"`
	Attachments []*Attachment      `json:"attachments,omitempty"`
	Seq         int                `json:"seq"`
	QuickReply  *QuickReplyPayload `json:"quick_reply,omitempty"`
	IsEcho      bool               `json:"is_echo,omitempty"`
	Metadata    *string            `json:"metadata,omitempty"`
}

type QuickReplyPayload struct {
	Payload string
}

type Delivery struct {
	MessageIDS []string `json:"mids"`
	Watermark  int64    `json:"watermark"`
	Seq        int      `json:"seq"`
}

type Postback struct {
	Payload string `json:"payload"`
}

type Optin struct {
	Ref string `json:"ref"`
}

type Read struct {
	Watermark int64 `json:"watermark"`
	Seq       int   `json:"seq"`
}

type MessageEcho struct {
	ReceivedMessage
	AppID int64 `json:"app_id,omitempty"`
}

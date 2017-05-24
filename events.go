package messenger

type upstreamEvent struct {
	Object  string          `json:"object"`
	Entries []*MessageEvent `json:"entry"`
}

// Event represents a Webhook postback event.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type Event struct {
	ID   string `json:"id"`
	Time int64  `json:"time"`
}

// MessageOpts contains information common to all message events.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type MessageOpts struct {
	Sender struct {
		ID string `json:"id"`
	} `json:"sender"`
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Timestamp int64 `json:"timestamp"`
}

// MessageEvent encapsulates common info plus the specific type of callback
// being received.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference#format
type MessageEvent struct {
	Event
	Messaging []struct {
		MessageOpts
		Message  *MessageEcho `json:"message,omitempty"`
		Delivery *Delivery    `json:"delivery,omitempty"`
		Postback *Postback    `json:"postback,omitempty"`
		Optin    *Optin       `json:"optin,omitempty"`
		Read     *Read        `json:"read,omitempty"`
		Referral *Referral    `json:"referral,omitempty"`
	} `json:"messaging"`
}

// ReceivedMessage contains message specific information included with an echo
// callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-echo
type ReceivedMessage struct {
	ID          string             `json:"mid"`
	Text        string             `json:"text,omitempty"`
	Attachments []*Attachment      `json:"attachments,omitempty"`
	Seq         int                `json:"seq"`
	QuickReply  *QuickReplyPayload `json:"quick_reply,omitempty"`
	IsEcho      bool               `json:"is_echo,omitempty"`
	Metadata    *string            `json:"metadata,omitempty"`
}

// QuickReplyPayload contains content specific to a quick reply.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message
type QuickReplyPayload struct {
	Payload string
}

// Delivery contains information specific to a message delivered callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-delivered
type Delivery struct {
	MessageIDS []string `json:"mids"`
	Watermark  int64    `json:"watermark"`
	Seq        int      `json:"seq"`
}

// Postback contains content specific to a postback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message
type Postback struct {
	Payload	 string `json:"payload"`
	Referral *Referral `json:"referral,omitempty"`
}

// Referral contains content specific to a referal.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/postback
type Referral struct {
	Ref	string`json:"referral,omitempty"`
	Source  string`json:"source,omitempty"`
	Type    string`json:"type,omitempty"`
}

// Optin contains information specific to Opt-In callbacks.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/optins
type Optin struct {
	Ref string `json:"ref"`
}

// Read contains data specific to message read callbacks.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-read
type Read struct {
	Watermark int64 `json:"watermark"`
	Seq       int   `json:"seq"`
}

// MessageEcho contains information specific to an echo callback.
// https://developers.facebook.com/docs/messenger-platform/webhook-reference/message-echo
type MessageEcho struct {
	ReceivedMessage
	AppID int64 `json:"app_id,omitempty"`
}

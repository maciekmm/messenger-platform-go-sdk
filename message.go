package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Message struct {
	Text       string     `json:"text,omiempty"`
	Attachment Attachment `json:"attachment,omitempty"`
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
	Recipient
	Message          Message `json:"message"`
	NotificationType `json:"notification_type,omitempty"`
}

type MessageResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}

type rawMessage struct {
	Recipient
	MessageQuery
}

func (m *Messenger) SendMessage(recipient Recipient, mq MessageQuery) (*MessageResponse, error) {
	rm := &rawMessage{
		Recipient:    recipient,
		MessageQuery: mq,
	}
	byt, err := json.Marshal(rm)
	if err != nil {
		return nil, err
	}
	resp, err := m.doRequest("POST", "https://graph.facebook.com/v2.6/me/messages", bytes.NewReader(byt))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid status code")
	}
	decoder := json.NewDecoder(resp.Body)
	response := &MessageResponse{}
	err = decoder.Decode(response)
	return response, err
}

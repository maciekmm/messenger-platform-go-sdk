package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Text       string      `json:"text,omiempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

// Recipient describes the person who will receive the message
// Either ID or PhoneNumber has to be set
type Recipient struct {
	ID          int64  `json:"id,omitempty"`
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
	Recipient        Recipient `json:"recipient"`
	Message          Message   `json:"message"`
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

func (m *Messenger) SendMessage(mq MessageQuery) (*MessageResponse, error) {
	byt, err := json.Marshal(mq)
	if err != nil {
		return nil, err
	}
	resp, err := m.doRequest("POST", GraphAPI+"/v2.6/me/messages", bytes.NewReader(byt))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		er := new(Error)
		json.Unmarshal(read, er)
		return nil, errors.New("Error occured: " + er.Message)
	}
	response := &MessageResponse{}
	err = json.Unmarshal(read, response)
	return response, err
}

func (m *Messenger) SendSimpleMessage(recipient int64, message string) (*MessageResponse, error) {
	return m.SendMessage(MessageQuery{
		Recipient: Recipient{
			ID: recipient,
		},
		Message: Message{
			Text: message,
		},
	})
}

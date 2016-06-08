package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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
		er := new(rawError)
		json.Unmarshal(read, er)
		return nil, errors.New("Error occured: " + er.Error.Message)
	}
	response := &MessageResponse{}
	err = json.Unmarshal(read, response)
	return response, err
}

func (m *Messenger) SendSimpleMessage(recipient string, message string) (*MessageResponse, error) {
	return m.SendMessage(MessageQuery{
		Recipient: Recipient{
			ID: recipient,
		},
		Message: SendMessage{
			Text: message,
		},
	})
}

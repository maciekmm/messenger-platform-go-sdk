package messenger

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSendMessageMarshalling(t *testing.T) {
	//Avoid HTTPS in tests
	GraphAPI = "http://example.com"
	messenger := &Messenger{}

	mockData := &MessageResponse{
		RecipientID: "11213",
		MessageID:   "abagfda",
	}

	body, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err)
	}

	setClient(200, body)

	profile, err := messenger.SendMessage(MessageQuery{})
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(profile, mockData) {
		t.Error("Response is invalid")
	}
}

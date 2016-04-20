package messenger

import (
	"encoding/json"
	"testing"
)

func TestSetWelcomeMessage(t *testing.T) {
	//Avoid HTTPS in tests
	GraphAPI = "http://example.com"
	messenger := &Messenger{
		PageID: "foo",
	}

	mockData := &result{
		Result: "Successfully added new_thread's CTAs",
	}

	body, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err)
	}

	setClient(200, body)

	err = messenger.SetWelcomeMessage(&Message{
		Text: "hello!",
	})
	if err != nil {
		t.Error(err)
	}

	mockData = &result{
		Result: "error!",
	}

	body, err = json.Marshal(mockData)
	if err != nil {
		t.Error(err)
	}
	setClient(200, body)

	err = messenger.SetWelcomeMessage(&Message{})
	if err == nil {
		t.Error("Error should have been thrown!")
	}
}

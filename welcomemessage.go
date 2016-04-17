package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ctaBase struct {
	SettingType string `json:"setting_type"`
	ThreadState string `json:"thread_state"`
}

var welcomeMessage = ctaBase{
	SettingType: "setting_type",
	ThreadState: "new_thread",
}

type cta struct {
	ctaBase
	CallToActions []*Message `json:"call_to_actions"`
}

type Result struct {
	Result string `json:"result"`
}

// SetWelcomeMessage sets the message that is sent first. If message is nil or empty the welcome message is not sent.
func (m *Messenger) SetWelcomeMessage(message *Message) error {
	cta := &cta{
		ctaBase:       welcomeMessage,
		CallToActions: []*Message{message},
	}
	if m.PageID == "" {
		return errors.New("PageID is empty")
	}
	byt, err := json.Marshal(cta)
	if err != nil {
		return err
	}
	resp, err := m.doRequest("POST", fmt.Sprintf("https://graph.facebook.com/v2.6/%s/thread_settings", m.PageID), bytes.NewReader(byt))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Invalid status code")
	}
	decoder := json.NewDecoder(resp.Body)
	result := &Result{}
	err = decoder.Decode(result)
	if err != nil {
		return err
	}
	if result.Result != "Successfully added new_thread's CTAs" {
		return errors.New("Something went wrong with setting thread's welcome message, facebook error: " + result.Result)
	}
	return nil
}

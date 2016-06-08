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
	SettingType: "call_to_actions",
	ThreadState: "new_thread",
}

type ctaMessage struct {
	Message *SendMessage `json:"message"`
}

type cta struct {
	ctaBase
	CallToActions []ctaMessage `json:"call_to_actions"`
}

type result struct {
	Result string `json:"result"`
}

// SetWelcomeMessage sets the message that is sent first. If message is nil or empty the welcome message is not sent.
func (m *Messenger) SetWelcomeMessage(message *SendMessage) error {
	cta := &cta{
		ctaBase:       welcomeMessage,
		CallToActions: []ctaMessage{ctaMessage{Message: message}},
	}
	if m.PageID == "" {
		return errors.New("PageID is empty")
	}
	byt, err := json.Marshal(cta)
	if err != nil {
		return err
	}
	resp, err := m.doRequest("POST", fmt.Sprintf(GraphAPI+"/v2.6/%s/thread_settings", m.PageID), bytes.NewReader(byt))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Invalid status code %d", resp.StatusCode)
	}
	decoder := json.NewDecoder(resp.Body)
	result := &result{}
	err = decoder.Decode(result)
	if err != nil {
		return err
	}
	if result.Result != "Successfully added new_thread's CTAs" {
		return errors.New("Something went wrong with setting thread's welcome message, facebook error: " + result.Result)
	}
	return nil
}

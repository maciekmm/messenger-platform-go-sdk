package messenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MenuItemType string

const (
	MenuItemTypeWebURL   MenuItemType = "web_url"
	MenuItemTypePostback MenuItemType = "postback"
	MenuItemTypeNested   MenuItemType = "nested"
)

type MenuItem struct {
	Type          MenuItemType `json:"type"`
	Title         string       `json:"title"`
	URL           string       `json:"url,omitempty"`
	Payload       string       `json:"payload,omitempty"`
	CallToActions []MenuItem   `json:"call_to_actions,omitempty"`
	// TODO webview_height_ratio
	// TODO messenger_extensions
	// TODO fallback_url
}

type PersistentMenu struct {
	Locale                string     `json:"locale"`
	ComposerInputDisabled bool       `json:"composer_input_disabled"`
	CallToActions         []MenuItem `json:"call_to_actions,omitempty"`
}

type PersistentMenuSettings struct {
	PersistentMenu []PersistentMenu `json:"persistent_menu"`
}

func (m *Messenger) SetNestedPersistentMenu(set *PersistentMenuSettings) error {
	body, err := json.Marshal(set)
	if err != nil {
		return err
	}

	resp, err := m.doRequest("POST", GraphAPI+"/v2.6/me/messenger_profile", bytes.NewReader(body))
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
	return nil
}

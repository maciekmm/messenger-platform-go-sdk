package messenger

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Profile struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_pic"`
}

func (m *Messenger) GetProfile(userID int64) (*Profile, error) {
	resp, err := m.doRequest("GET", fmt.Sprintf("https://graph.facebook.com/v2.6/%d?fields=first_name,last_name,profile_pic", userID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid status code")
	}
	decoder := json.NewDecoder(resp.Body)
	profile := new(Profile)
	return profile, decoder.Decode(profile)
}

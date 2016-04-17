package messenger

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type MessageReceivedHandler func(Event, MessageOpts, ReceivedMessage)
type MessageDeliveredHandler func(Event, MessageOpts, Delivery)
type PostbackHandler func(Event, MessageOpts, Postback)
type AuthenticationHandler func(Event, MessageOpts, *Optin)

type Messenger struct {
	VerifyToken      string
	AppSecret        string
	AccessToken      string
	PageID           string
	MessageReceived  MessageReceivedHandler
	MessageDelivered MessageDeliveredHandler
	Postback         PostbackHandler
	Authentication   AuthenticationHandler
}

func (m *Messenger) Handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		query := req.URL.Query()
		if query.Get("hub.verify_token") != m.VerifyToken {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(query.Get("hub.challenge")))
	} else if req.Method == "POST" {
		read, err := ioutil.ReadAll(req.Body)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		//Message integrity check
		if m.AppSecret != "" {
			mac := hmac.New(sha1.New, []byte(m.AppSecret))
			mac.Write(read)
			if fmt.Sprintf("%x", mac.Sum(nil)) != req.Header.Get("x-hub-signature")[5:] {
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		event := &rawEvent{}
		err = json.Unmarshal(read, event)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		for _, entry := range event.Entries {
			for _, message := range entry.Messaging {
				if message.Delivery != nil {
					if m.MessageDelivered != nil {
						go m.MessageDelivered(entry.Event, message.MessageOpts, *message.Delivery)
					}
				} else if message.Message != nil {
					if m.MessageReceived != nil {
						go m.MessageReceived(entry.Event, message.MessageOpts, *message.Message)
					}
				} else if message.Postback != nil {
					if m.Postback != nil {
						go m.Postback(entry.Event, message.MessageOpts, *message.Postback)
					}
				} else if m.Authentication != nil {
					go m.Authentication(entry.Event, message.MessageOpts, message.Optin)
				}
			}
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"status":"ok"}`))
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (m *Messenger) doRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	query := req.URL.Query()
	query.Set("access_token", m.AccessToken)
	req.URL.RawQuery = query.Encode()
	return http.DefaultClient.Do(req)
}

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

var (
	//GraphAPI specifies host used for API requests
	GraphAPI = "https://graph.facebook.com"
)

// MessageReceivedHandler is called when a new message is received
type MessageReceivedHandler func(Event, MessageOpts, ReceivedMessage)

// MessageDeliveredHandler is called when a message sent has been successfully delivered
type MessageDeliveredHandler func(Event, MessageOpts, Delivery)

// PostbackHandler is called when the postback button has been pressed by recipient
type PostbackHandler func(Event, MessageOpts, Postback)

// AuthenticationHandler is called when a new user joins/authenticates
type AuthenticationHandler func(Event, MessageOpts, *Optin)

// MessageReadHandler is called when a message has been read by recipient
type MessageReadHandler func(Event, MessageOpts, Read)

// MessageEchoHandler is called when a message is sent by your page
type MessageEchoHandler func(Event, MessageOpts, MessageEcho)

// Messenger is the main service which handles all callbacks from facebook
// Events are delivered to handlers if they are specified
type Messenger struct {
	VerifyToken string
	AppSecret   string
	AccessToken string
	PageID      string

	MessageReceived  MessageReceivedHandler
	MessageDelivered MessageDeliveredHandler
	Postback         PostbackHandler
	Authentication   AuthenticationHandler
	MessageRead      MessageReadHandler
	MessageEcho      MessageEchoHandler
}

// Handler is the main HTTP handler for the Messenger service.
// It MUST be attached to some web server in order to receive messages
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
		m.handlePOST(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (m *Messenger) handlePOST(rw http.ResponseWriter, req *http.Request) {
	read, err := ioutil.ReadAll(req.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	//Message integrity check
	if m.AppSecret != "" {
		if len(req.Header.Get("x-hub-signature")) < 6 || !checkIntegrity(m.AppSecret, read, req.Header.Get("x-hub-signature")[5:]) {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	event := &upstreamEvent{}
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
			} else if message.Message != nil && message.Message.IsEcho {
				if m.MessageEcho != nil {
					go m.MessageEcho(entry.Event, message.MessageOpts, *message.Message)
				}
			} else if message.Message != nil {
				if m.MessageReceived != nil {
					go m.MessageReceived(entry.Event, message.MessageOpts, message.Message.ReceivedMessage)
				}
			} else if message.Postback != nil {
				if m.Postback != nil {
					go m.Postback(entry.Event, message.MessageOpts, *message.Postback)
				}
			} else if message.Read != nil {
				if m.MessageRead != nil {
					go m.MessageRead(entry.Event, message.MessageOpts, *message.Read)
				}
			} else if m.Authentication != nil {
				go m.Authentication(entry.Event, message.MessageOpts, message.Optin)
			}
		}
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status":"ok"}`))
}

func checkIntegrity(appSecret string, bytes []byte, expectedSignature string) bool {
	mac := hmac.New(sha1.New, []byte(appSecret))
	mac.Write(bytes)
	if fmt.Sprintf("%x", mac.Sum(nil)) != expectedSignature {
		return false
	}
	return true
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

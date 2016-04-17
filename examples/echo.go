package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maciekmm/messenger-platform-go-sdk"
)

var mess = &messenger.Messenger{
	AccessToken: "EAAYRloZALk2cBAGahz3hMWWVZB0JgQOhvC27IBRgZA3Lm6QDIfIyZChVuJnmRZBfKl6B3ZADSJNUrrgacD04i1No23QMWqPaajUM7bjV2LBL3apM6mSYr2S1A4JIynEWja7T4ZAFAFuQp2APerrxrIeIZCNH06ZCZBEWFntKn5LDU46AZDZD",
}

func main() {
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":5646", nil))
}

func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
	resp, err := mess.SendMessage(messenger.MessageQuery{Recipient: messenger.Recipient{ID: opts.Sender.ID}, Message: messenger.Message{
		Text: msg.Text,
	}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maciekmm/messenger-platform-go-sdk"
)

var mess = &messenger.Messenger{
	AccessToken: "ACCESS_TOKEN",
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

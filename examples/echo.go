package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maciekmm/messenger-platform-go-sdk"
)

var mess = &messenger.Messenger{
	AccessToken: "ACCESS_TOKEN",
	AppSecret:   "APP_SECRET",
}

func main() {
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":5646", nil))
}

func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
	resp, err := mess.SendSimpleMessage(opts.Sender.ID, msg.Text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}

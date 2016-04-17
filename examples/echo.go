package main

import (
	"fmt"

	"github.com/maciekmm/messenger-platform-go-sdk"
)

func init() {
	simpleEcho := func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		resp, err := mess.SendSimpleMessage(opts.Sender.ID, msg.Text)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", resp)
	}
	mess.MessageReceived = simpleEcho
}

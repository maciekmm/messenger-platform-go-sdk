package main

import (
	"fmt"

	"github.com/mgilbir/messenger-platform-go-sdk"
)

func init() {
	simpleEcho := func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		profile, err := mess.GetProfile(opts.Sender.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Hello, %s %s, %s", profile.FirstName, profile.LastName, msg.Text))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", resp)
	}
	mess.MessageReceived = simpleEcho
}

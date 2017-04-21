package main

import (
	"fmt"

	"github.com/mgilbir/messenger-platform-go-sdk"
	"github.com/mgilbir/messenger-platform-go-sdk/template"
)

func init() {
	handler := func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		mq := messenger.MessageQuery{}
		mq.RecipientID(opts.Sender.ID)
		mq.Template(template.GenericTemplate{Title: "abc",
			Buttons: []template.Button{
				template.Button{
					Type:    template.ButtonTypePostback,
					Payload: "test",
					Title:   "abecadło",
				},
			},
		})
		resp, err := mess.SendMessage(mq)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", resp)
	}
	mess.MessageReceived = handler
}

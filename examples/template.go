package main

import (
	"fmt"

	"github.com/maciekmm/messenger-platform-go-sdk"
	"github.com/maciekmm/messenger-platform-go-sdk/template"
)

func init() {
	handler := func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		resp, err := mess.SendMessage(messenger.MessageQuery{Recipient: messenger.Recipient{ID: opts.Sender.ID}, Message: messenger.Message{
			Attachment: &messenger.Attachment{
				Type: messenger.AttachmentTypeTemplate,
				Payload: &template.Payload{
					Elements: []template.Template{
						template.Generic{Title: "abc",
							Buttons: []template.Button{
								template.Button{
									Type:    template.ButtonTypePostback,
									Payload: "test",
									Title:   "abecadło",
								},
							},
						},
					},
				},
			},
		}})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", resp)
	}
	mess.MessageReceived = handler
}

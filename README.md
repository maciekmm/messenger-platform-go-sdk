Messenger Platform Go SDK
=====

[![Build Status](https://travis-ci.org/maciekmm/messenger-platform-go-sdk.svg?branch=master)](https://travis-ci.org/maciekmm/messenger-platform-go-sdk) 
[![Coverage Status](https://coveralls.io/repos/github/maciekmm/messenger-platform-go-sdk/badge.svg?branch=master)](https://coveralls.io/github/maciekmm/messenger-platform-go-sdk?branch=master)

A Go SDK for the [Facebook Messenger Platform](https://developers.facebook.com/docs/messenger-platform).
**Note: Work in Progress, not suitable for production environment yet! Be careful!**

## Installation

```bash
go get gopkg.in/maciekmm/messenger-platform-go-sdk.v4
```

## Usage

The main package has been named `messenger` for convenience. 

Your first step is to create `Messenger` instance.

```go
import "gopkg.in/maciekmm/messenger-platform-go-sdk.v4"

//...

messenger := &messenger.Messenger {
	VerifyToken: "VERIFY_TOKEN/optional",
	AppSecret: "APP_SECRET/optional",
	AccessToken: "PAGE_ACCESS_TOKEN",
	PageID: "PAGE_ID/optional",
}
```

### Parameters
* `VerifyToken` is the token needed for a verification process facebook performs. It's only required once. Optional.
* `AppSecret` is the Application Secret token. It's used for message integrity check. Optional.
* `AccessToken` is required to send messages. You can find this token in your app developer dashboard under `Messenger` tab.
* `PageID` is required for setting welcome message. Optional.

The next step is to hook up the handler to your HTTP server. 

```go
//hook up
http.HandleFunc("/webhook", messenger.Handler)
//start the server
http.ListenAndServe(":5646", nil)
```

The next step is to *subscribe* to an event, to do that you have to hook up your own handler.

```go
messenger.MessageReceived = MessageReceived

//...

func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
//do stuff
}
```

### Sending messages

## Example

Check more examples in [examples folder.](https://github.com/maciekmm/messenger-platform-go-sdk/tree/master/examples)

```go
var mess = &messenger.Messenger{
	AccessToken: "ACCESS_TOKEN",
}

func main() {
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":5646", nil))
}

func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
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
```
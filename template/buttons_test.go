package template

import "testing"

func TestURLButtonCreation(t *testing.T) {
	button := NewWebURLButton("title", "http://github.com/")
	if button.Payload != "" {
		t.Error("URL button's payload is not empty")
	}
	if button.Title != "title" {
		t.Error("URL button's title is not correct")
	}
	if button.URL != "http://github.com/" {
		t.Error("URL button's url is not correct")
	}
}

func TestPostbackButtonCreation(t *testing.T) {
	button := NewPostbackButton("title", "postback")
	if button.URL != "" {
		t.Error("Postback button's payload is not empty")
	}
	if button.Title != "title" {
		t.Error("Postback button's title is not correct")
	}
	if button.Payload != "postback" {
		t.Error("Postback button's url is not correct")
	}
}

func TestPhoneNumberButtonCreation(t *testing.T) {
	button := NewPhoneNumberButton("title", "+1123123123")
	if button.URL != "" {
		t.Error("PhoneNumber button's payload is not empty")
	}
	if button.Title != "title" {
		t.Error("PhoneNumber button's title is not correct")
	}
	if button.Payload != "+1123123123" {
		t.Error("PhoneNumber button's payload is not correct")
	}
}

func TestAccountLinkButtonCreation(t *testing.T) {
	button := NewAccountLinkButton("https://example.com/oauth/authorize")
	if button.Title != "" {
		t.Error("AccountLink button's title is not empty")
	}
	if button.Payload != "" {
		t.Error("AccountLink button's payload is not empty")
	}
	if button.URL != "https://example.com/oauth/authorize" {
		t.Error("AccountLink button's url is incorrect")
	}
}

func TestAccountUnlinkButtonCreation(t *testing.T) {
	button := NewAccountUnlinkButton()
	if button.Title != "" {
		t.Error("AccountUnlink button's title is not empty")
	}
	if button.Payload != "" {
		t.Error("AccountUnlink button's payload is not empty")
	}
	if button.URL != "" {
		t.Error("AccountLink button's url is not empty")
	}
}

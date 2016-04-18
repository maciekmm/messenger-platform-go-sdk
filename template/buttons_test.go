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

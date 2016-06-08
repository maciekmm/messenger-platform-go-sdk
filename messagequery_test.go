package messenger

import (
	"testing"

	"github.com/maciekmm/messenger-platform-go-sdk/template"
)

func TestRecipientID(t *testing.T) {
	mq := MessageQuery{}
	err := mq.RecipientID("123145512251")
	if err != nil || mq.Recipient.ID != "123145512251" {
		t.Error("RecipientID cannot be set.", err)
	}
	mq.Recipient.PhoneNumber = "111-111-111"
	err = mq.RecipientID("2142141241241241")
	if err == nil {
		t.Error("Should not permit setting both RecipientID and PhoneNumber.")
	}
}

func TestPhoneNumber(t *testing.T) {
	mq := MessageQuery{}
	err := mq.RecipientPhoneNumber("111-111-111")
	if err != nil {
		t.Error("PhoneNumber cannot be set.", err)
	}
	mq.Recipient.ID = "23123123131231"
	err = mq.RecipientPhoneNumber("3213-3232-3131")
	if err == nil {
		t.Error("Should not permit setting both ID and PhoneNumber.")
	}
}

func TestNotification(t *testing.T) {
	mq := MessageQuery{}
	mq.Notification(NotificationTypeRegular)
	if mq.NotificationType != NotificationTypeRegular {
		t.Error("NotificationType cannot be set.")
	}
}

func TestText(t *testing.T) {
	mq := MessageQuery{}
	err := mq.Text("Text!")
	if err != nil || mq.Message.Text != "Text!" {
		t.Error("Cannot set text message.")
	}
	err = mq.Template(&template.GenericTemplate{})
	if err != nil {
		t.Error("Cannot set template.", err.Error())
	}
	err = mq.Text("test!")
	if err == nil {
		t.Error("Should not permit specifying both text and template.")
	}
}

func TestResources(t *testing.T) {
	mq := MessageQuery{}
	err := mq.Audio("http://example.com/audio.mp3")
	if err != nil {
		t.Error("Cannot set audio.", err)
	}
	if mq.Message.Attachment.Type != AttachmentTypeAudio {
		t.Error("Invalid Type was set.")
	}
	err = mq.Audio("https://com.example/audio.mp3")
	if err == nil {
		t.Error("Should not permit overriding payloads.")
	}
	mq.Message.Attachment = nil
	err = mq.Video("https://com.example/video.mp4")
	if err != nil {
		t.Error("Cannot set video.", err)
	}
	mq.Message.Attachment = nil
	err = mq.Image("https://com.example/image.jpg")
	if err != nil {
		t.Error("Cannot set image.", err)
	}
}

func TestTemplate(t *testing.T) {
	mq := MessageQuery{}
	err := mq.Template(&template.GenericTemplate{})
	if err != nil {
		t.Error("Cannot set template.", err)
	}
	temp := &template.GenericTemplate{Title: "test"}
	err = mq.Template(temp)
	if err != nil || mq.Message.Attachment.Payload.(*template.Payload).Elements[1] != temp {
		t.Error("Cannot add template or the template does not match.", err)
	}
	err = mq.Template(&template.ButtonTemplate{})
	if err == nil {
		t.Error("Should not allow adding two distinct templates.")
	}
	mq.Message.Attachment = nil
	err = mq.Audio("https://com.example/audio.mp3")
	if err != nil {
		t.Error("Error occured while setting audio as attachment.")
	}
	err = mq.Template(temp)
	if err == nil {
		t.Error("Should not allow for adding template if another attachment is specified.")
	}
}

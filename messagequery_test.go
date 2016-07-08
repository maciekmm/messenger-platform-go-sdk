package messenger

import (
	"math/rand"
	"testing"
	"time"

	"github.com/maciekmm/messenger-platform-go-sdk/template"
)

//seed the rand
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(leng uint) string {
	b := make([]rune, leng)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestRandString(t *testing.T) {
	if len(randString(0)) != 0 {
		t.Error("Invalid length of generated string")
	}
	if len(randString(20)) != 20 {
		t.Error("Invalid length of generated string")
	}
	if len(randString(1000)) != 1000 {
		t.Error("Invalid length of generated string")
	}
}

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

func TestQuickReply(t *testing.T) {
	mq := MessageQuery{}
	err := mq.QuickReply("Valid Title", "Valid Payload")
	if err != nil {
		t.Error("Cannot add quick reply", err)
	}
	//20 characters edge case
	err = mq.QuickReply(randString(20), randString(1000))
	if err != nil {
		t.Error("Cannot add quick reply", err)
	}
	err = mq.QuickReply(randString(100), randString(2000))
	if err == nil {
		t.Error("Can add an invalid quick reply")
	}
	for _, qr := range mq.Message.QuickReplies {
		if qr.ContentType != "text" {
			t.Error("QuickReply content type must be 'text'")
		}
	}
}

func TestMetadata(t *testing.T) {
	mq := MessageQuery{}
	err := mq.Metadata("Correct metadata.")
	if err != nil {
		t.Error("Cannot add metadata", err)
	}
	//edge case
	err = mq.Metadata(randString(1000))
	if err != nil {
		t.Error("Cannot add metadata", err)
	}
	err = mq.Metadata(randString(2000))
	if err == nil {
		t.Error("Can add an invalid metadata")
	}
}

package messenger

import (
	"testing"

	"github.com/mgilbir/messenger-platform-go-sdk/template"
)

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

	//Check if adding valid one works
	err := mq.QuickReply(QuickReply{
		Title:   "Valid",
		Payload: "Valid",
	})
	if err != nil {
		t.Error("Cannot add quick reply", err)
	}
	//Check if automatically populates
	if mq.Message.QuickReplies[0].ContentType != ContentTypeText {
		t.Error("ContentType was not populated")
	}

	//20 characters edge case
	err = mq.QuickReply(QuickReply{
		Title:   randString(20),
		Payload: randString(1000),
	})
	if err != nil {
		t.Error("Cannot add quick reply", err)
	}

	//Title too long
	err = mq.QuickReply(QuickReply{
		Title:   randString(100),
		Payload: randString(100),
	})
	if err == nil {
		t.Error("Can add an invalid quick reply")
	}

	//Payload too long
	err = mq.QuickReply(QuickReply{
		Title:   randString(20),
		Payload: randString(2000),
	})
	if err == nil {
		t.Error("Can add an invalid quick reply")
	}

	//Populate to the limit of 10 messages
	//TODO: Thing whether we should use constant here (and above as well)
	toAdd := 10 - len(mq.Message.QuickReplies)
	for i := 0; i < toAdd; i++ {
		err = mq.QuickReply(QuickReply{
			Title:   "a",
			Payload: "b",
		})
	}

	//Adding a quick message over the limit should yield an error
	err = mq.QuickReply(QuickReply{
		Title:   randString(20),
		Payload: randString(1000),
	})
	if err == nil {
		t.Error("Can add quick reply over the limit", err)
	}

	//Location does not support title or payload
	err = mq.QuickReply(QuickReply{
		ContentType: ContentTypeLocation,
		Title:       "Title",
		Payload:     "Payload",
	})
	if err == nil {
		t.Error("Can add quick reply of type location with title or payload")
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

package messenger

import (
	"testing"

	"github.com/mgilbir/messenger-platform-go-sdk/template"
)

func TestSetGreetingText(t *testing.T) {
	GraphAPI = "http://example.com"
	mess := &Messenger{}

	setClient(200, []byte(`{"result":"Successfully updated greeting"}`))
	err := mess.SetGreetingText(randString(180))
	if err == nil {
		t.Error("Error shouldn't be empty, length of the greeting text exceeds maximum length.")
	}

	err = mess.SetGreetingText(randString(60))
	if err != nil {
		t.Error(err)
	}
	setClient(500, []byte(`{"result":"Error occured while updating greeting"}`))
	err = mess.SetGreetingText(randString(60))
	if err == nil {
		t.Error("Error shouldn't be empty")
	}
	setClient(200, []byte(`{"result":"Error occured while updating greeting"}`))
	err = mess.SetGreetingText(randString(60))
	if err == nil {
		t.Error("Error shouldn't be empty")
	}
}

func TestSetGetStartedButton(t *testing.T) {
	GraphAPI = "http://example.com"
	mess := &Messenger{}

	setClient(200, []byte(`{"result":"Successfully added new_thread's CTAs"}`))
	err := mess.SetGetStartedButton(randString(60))

	if err != nil {
		t.Error(err)
	}

	setClient(500, []byte(`{"result":"Error occured while updating CTAs"}`))
	err = mess.SetGetStartedButton(randString(60))
	if err == nil {
		t.Error("Error shouldn't be empty")
	}

	setClient(200, []byte(`{"result":"Error occured while updating CTAs"}`))
	err = mess.SetGetStartedButton(randString(60))
	if err == nil {
		t.Error("Error shouldn't be empty")
	}
}

func TestDeleteGetStartedButton(t *testing.T) {
	GraphAPI = "http://example.com"
	mess := &Messenger{}

	setClient(200, []byte(`{"result":"Successfully deleted all new_thread's CTAs"}`))
	err := mess.DeleteGetStartedButton()

	if err != nil {
		t.Error(err)
	}

	setClient(500, []byte(`{"result":"Error occured while deleting CTAs"}`))
	err = mess.DeleteGetStartedButton()
	if err == nil {
		t.Error("Error shouldn't be empty")
	}

	setClient(200, []byte(`{"result":"Error occured while deleting CTAs"}`))
	err = mess.DeleteGetStartedButton()
	if err == nil {
		t.Error("Error shouldn't be empty")
	}
}

func TestSetPersistentMenu(t *testing.T) {
	GraphAPI = "http://example.com"
	mess := &Messenger{}
	button := template.NewWebURLButton("test", "http://example.com")

	setClient(200, []byte(`{"result":"Successfully added structured menu CTAs"}`))
	err := mess.SetPersistentMenu([]template.Button{button, button, button, button, button, button})
	if err == nil {
		t.Error("Error shouldn't be empty, number of buttons exceeds current limit")
	}
	err = mess.SetPersistentMenu([]template.Button{button, button, button, button, button})
	if err != nil {
		t.Error(err)
	}
	err = mess.SetPersistentMenu([]template.Button{button, button})
	if err != nil {
		t.Error(err)
	}

	setClient(500, []byte(`{"result":"Error occured while adding structured menu"}`))
	err = mess.SetPersistentMenu([]template.Button{button, button})
	if err == nil {
		t.Error("Error shouldn't be empty")
	}

	setClient(200, []byte(`{"result":"Error occured while adding structured menu"}`))
	err = mess.SetPersistentMenu([]template.Button{button, button})
	if err == nil {
		t.Error("Error shouldn't be empty")
	}
}

func TestDeletePersistentMenu(t *testing.T) {
	GraphAPI = "http://example.com"
	mess := &Messenger{}

	setClient(200, []byte(`{"result":"Successfully deleted structured menu CTAs"}`))
	err := mess.DeletePersistentMenu()

	if err != nil {
		t.Error(err)
	}

	setClient(500, []byte(`{"result":"Error occured while deleting CTAs"}`))
	err = mess.DeletePersistentMenu()
	if err == nil {
		t.Error("Error shouldn't be empty")
	}

	setClient(200, []byte(`{"result":"Error occured while deleting CTAs"}`))
	err = mess.DeletePersistentMenu()
	if err == nil {
		t.Error("Error shouldn't be empty")
	}
}

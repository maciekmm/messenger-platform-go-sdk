package messenger

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetProfile(t *testing.T) {
	//Avoid HTTPS in tests
	GraphAPI = "http://example.com"
	messenger := &Messenger{}

	mockData := &Profile{
		FirstName:      "John",
		LastName:       "Smith",
		ProfilePicture: "https://example.com/",
		Gender:         "male",
		Timezone:       -5,
		Locale:         "en_US",
	}

	body, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err)
	}

	setClient(200, body)

	profile, err := messenger.GetProfile("123")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(profile, mockData) {
		t.Error("Profiles do not match")
	}

	errorData := &rawError{Error: Error{
		Message: "w/e",
	}}
	body, err = json.Marshal(errorData)
	if err != nil {
		t.Error(err)
	}
	setClient(400, body)
	_, err = messenger.GetProfile("123")
	if err.Error() != "Error occured: "+errorData.Error.Message {
		t.Error("Invalid error parsing")
	}
}

func TestGetPSID(t *testing.T) {
	GraphAPI = "http://example.com"
	messenger := &Messenger{}
	setClient(200, []byte(`{"recipient":"userid", "id":"pageid"}`))
	psid, err := messenger.GetPSID("token")
	if err != nil {
		t.Error(err)
	}
	if *psid != "userid" {
		t.Error("Invalid userid")
	}
	setClient(400, []byte(""))
	_, err = messenger.GetPSID("token")
	if err == nil {
		t.Error("Error shouldn't be empty.")
	}
}

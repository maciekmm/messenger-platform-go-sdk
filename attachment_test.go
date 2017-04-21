package messenger

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mgilbir/messenger-platform-go-sdk/template"
)

func TestAttachmentUnmarshalling(t *testing.T) {
	//Location
	testMarshaller(Attachment{
		Type: AttachmentTypeLocation,
		Payload: &Location{
			Coordinates: Coordinates{
				Latitude:  22.2,
				Longitude: 10.1,
			},
		},
	}, t)
	//Resource
	testMarshaller(Attachment{
		Type: AttachmentTypeAudio,
		Payload: &Resource{
			URL: "http://example.com/example.mp3",
		},
	}, t)
	//Template
	att := Attachment{
		Type: AttachmentTypeTemplate,
		Payload: &template.Payload{
			Elements: []template.Template{
				template.GenericTemplate{
					Title: "abc",
				},
			},
		},
	}
	bytes, err := json.Marshal(&att)
	if err != nil {
		t.Error(err)
	}
	dest := Attachment{}
	err = json.Unmarshal(bytes, &dest)
	if err != nil {
		t.Error(err)
	}
	if _, ok := dest.Payload.(json.RawMessage); !ok {
		t.Error("Invalid template attachment unmarshalling")
	}
}

func testMarshaller(att Attachment, t *testing.T) {
	bytes, err := json.Marshal(&att)
	if err != nil {
		t.Error(err)
	}
	dest := Attachment{}
	err = json.Unmarshal(bytes, &dest)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(dest, att) {
		t.Error("Source and destination do not match.", att, dest)
	}
}

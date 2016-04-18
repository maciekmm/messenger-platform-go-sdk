package template

import (
	"encoding/json"
	"testing"
)

func TestTemplateMarshaling(t *testing.T) {
	payload := &Payload{
		Elements: []Template{
			GenericTemplate{Title: "abc",
				Buttons: []Button{
					Button{
						Type:    ButtonTypePostback,
						Payload: "test",
						Title:   "abecadło",
					},
				},
			},
		},
	}
	mock := &rawPayload{
		Type:     TemplateTypeGeneric,
		Elements: payload.Elements,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Error(err)
	}
	mockBytes, err := json.Marshal(mock)
	if err != nil {
		t.Error(err)
	}
	if string(mockBytes) != string(payloadBytes) {
		t.Error("Payloads do not match")
	}

	payload = &Payload{
		Elements: []Template{},
	}
	_, err = json.Marshal(payload)
	if err == nil {
		t.Error("Marshalling error is not thrown")
	}
}

package messenger

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEventUnmarshalJSON(t *testing.T) {
	rawPostbackData := []byte(`{"id":"1234","time":1458692752478}`)
	rawPageData := []byte(`{"id":"1234","time":1458692752478}`)
	postbackEvent := &Event{}
	err := json.Unmarshal(rawPostbackData, postbackEvent)
	if err != nil {
		t.Error(err)
	}
	pageEvent := &Event{}
	err = json.Unmarshal(rawPageData, pageEvent)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(*postbackEvent, *pageEvent) {
		t.Error("Events do not match")
	}
}

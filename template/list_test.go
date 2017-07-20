package template

import "testing"

func TestGenericType(t *testing.T) {
	template := &ListTemplate{}
	if template.Type() != TemplateTypeList {
		t.Error("List template returned invalid type")
	}
	if !template.SupportsButtons() {
		t.Error("List template supports buttons, but reports otherwise")
	}
}

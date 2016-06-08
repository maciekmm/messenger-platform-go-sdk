package template

import "testing"

func TestReceiptType(t *testing.T) {
	template := &ReceiptTemplate{}
	if template.Type() != TemplateTypeReceipt {
		t.Error("Receipt template returned invalid type")
	}
	if template.SupportsButtons() {
		t.Error("Button template is marked as supporting buttons.")
	}
}

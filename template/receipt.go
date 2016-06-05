package template

const TemplateTypeReceipt TemplateType = "receipt"

type ReceiptTemplate struct {
	RecipientName string           `json:"recipient_name"`
	OrderNumber   string           `json:"order_number"`
	Currency      string           `json:"currency"`
	PaymentMethod string           `json:"payment_method"`
	Timestamp     string           `json:"timestamp,omitempty"`
	OrderURL      string           `json:"order_url,omitempty"`
	Elements      []ReceiptElement `json:"elements"`
	Address       Address          `json:"address,omitempty"`
	Summary       Summary          `json:"summary"`
	Adjustments   []Adjustment     `json:"adjustments,omitempty"`
}

func (ReceiptTemplate) Type() TemplateType {
	return TemplateTypeReceipt
}

func (ReceiptTemplate) SupportsButtons() bool {
	return false
}

type ReceiptElement struct {
	Title    string  `json:"title"`
	Subtitle string  `json:"subtitle,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Currency string  `json:"currency,omitempty"`
	ImageURL string  `json:"image_url,omitempty"`
}

type Address struct {
	Street1    string `json:"street_1"`
	Street2    string `json:"street_2,omitempty"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	//??
	State   string `json:"state,omitempty"`
	Country string `json:"country"`
}

type Summary struct {
	Subtotal     float64 `json:"subtotal,omitempty"`
	ShippingCost float64 `json:"shipping_cost,omitempty"`
	TotalTax     float64 `json:"total_text,omitempty"`
	TotalCost    float64 `json:"total_cost"`
}

type Adjustment struct {
	Name   string  `json:"name,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

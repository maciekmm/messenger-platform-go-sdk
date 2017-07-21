package template

const TemplateTypeList TemplateType = "list"

type ListTemplate struct {
	Elements []ListElement `json:"elements"`
	Buttons  []Button      `json:"buttons,omitempty"`
}

func (ListTemplate) Type() TemplateType {
	return TemplateTypeList
}

func (ListTemplate) SupportsButtons() bool {
	return true
}

type ListElement struct {
	Title         string `json:"title"`
	ImageURL      string `json:"image_url,omitempty"`
	Subtitle      string `json:"subtitle,omitempty"`
	DefaultAction `json:"default_action,omitempty"`
	Buttons       []ListButton `json:"buttons,omitempty"`
}

type DefaultAction struct {
	Type                string `json:"type"`
	URL                 string `json:"url",omitempty`
	MessengerExtensions bool   `json:"messenger_extensions",omitempty`
	WebviewHeightRatio  string `json:"webview_height_ratio",omitempty`
	FallbackURL         string `json:"fallback_url",omitempty`
}

type ListButton struct {
	Title               string `json:"title"`
	Type                string `json:"type",omitempty`
	URL                 string `json:"url",omitempty`
	MessengerExtensions bool   `json:"messenger_extensions",omitempty`
	WebviewHeightRatio  string `json:"webview_height_ratio",omitempty`
	FallbackURL         string `json:"fallback_url",omitempty`
}

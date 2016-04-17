package template

// ButtonType defines the behavior of the button in the ButtonTemplate
type ButtonType string

const (
	ButtonTypeWebURL   ButtonType = "web_url"
	ButtonTypePostback ButtonType = "postback"
)

type Button struct {
	Type    ButtonType `json:"type"`
	Title   string     `json:"title"`
	URL     string     `json:"url,omitempty"`
	Payload string     `json:"payload,omitempty"`
}

// NewWebURLButton creates a button used in ButtonTemplate that redirects user to external address upon clicking the URL
func NewWebURLButton(title string, url string) Button {
	return Button{
		Type:  ButtonTypeWebURL,
		Title: title,
		URL:   url,
	}
}

// NewPostbackButton creates a button used in ButtonTemplate that upon clicking sends
func NewPostbackButton(title string, payload string) Button {
	return Button{
		Type:    ButtonTypePostback,
		Title:   title,
		Payload: payload,
	}
}

package template

// ButtonType defines the behavior of the button in the ButtonTemplate
type ButtonType string

const (
	ButtonTypeWebURL        ButtonType = "web_url"
	ButtonTypePostback      ButtonType = "postback"
	ButtonTypePhoneNumber   ButtonType = "phone_number"
	ButtonTypeAccountLink   ButtonType = "account_link"
	ButtonTypeAccountUnlink ButtonType = "account_unlink"
)

type Button struct {
	Type    ButtonType `json:"type"`
	Title   string     `json:"title,omitempty"`
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

// NewPostbackButton creates a button used in ButtonTemplate that upon clicking sends a payload request to the server
func NewPostbackButton(title string, payload string) Button {
	return Button{
		Type:    ButtonTypePostback,
		Title:   title,
		Payload: payload,
	}
}

// NewPhoneNumberButton creates a button used in ButtonTemplate that upon clicking opens a native dialer
func NewPhoneNumberButton(title string, phone string) Button {
	return Button{
		Type:    ButtonTypePhoneNumber,
		Title:   title,
		Payload: phone,
	}
}

// NewAccountLinkButton creates a button used for account linking
// https://developers.facebook.com/docs/messenger-platform/account-linking/authentication
func NewAccountLinkButton(url string) Button {
	return Button{
		Type: ButtonTypeAccountLink,
		URL:  url,
	}
}

// NewAccountUnlinkButton creates a button used for account unlinking
// https://developers.facebook.com/docs/messenger-platform/account-linking/authentication
func NewAccountUnlinkButton() Button {
	return Button{
		Type: ButtonTypeAccountUnlink,
	}
}

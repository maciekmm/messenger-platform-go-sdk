package template

import "fmt"

// ButtonType defines the behavior of the button in the ButtonTemplate
type ButtonType string

const (
	ButtonTypeWebURL        ButtonType = "web_url"
	ButtonTypePostback      ButtonType = "postback"
	ButtonTypePhoneNumber   ButtonType = "phone_number"
	ButtonTypeAccountLink   ButtonType = "account_link"
	ButtonTypeAccountUnlink ButtonType = "account_unlink"
	ButtonTypeShare         ButtonType = "element_share"
)

type Button struct {
	Type          ButtonType `json:"type"`
	Title         string     `json:"title,omitempty"`
	URL           string     `json:"url,omitempty"`
	Payload       string     `json:"payload,omitempty"`
	ShareContents Template   `json:"share_contents,omitempty"`
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

// NewSharebutton creates a button used for sharing
// https://developers.facebook.com/docs/messenger-platform/send-api-reference/share-button
func NewShareButton(shareContent ...Template) (Button, error) {
	b := Button{
		Type: ButtonTypeShare,
	}

	if len(shareContent) > 1 {
		return Button{}, fmt.Errorf("A single Template may be passed. Got %d", len(shareContent))
	}

	if len(shareContent) == 1 {
		sc := shareContent[0]

		switch scc := sc.(type) {
		case GenericTemplate:
			if len(scc.Buttons) > 1 {
				return Button{}, fmt.Errorf("The shareContent has to be a generic template with at most 1 button")
			}
			if len(scc.Buttons) == 1 {
				if scc.Buttons[0].Type != ButtonTypeWebURL {
					return Button{}, fmt.Errorf("The shareContent can have a URL button, but got %T", scc.Buttons[0].Type)
				}
			}
		default:
			return Button{}, fmt.Errorf("The shareContent Template has to be of type generic. Got %T", sc)
		}

		b.ShareContents = sc
	}

	return b, nil
}

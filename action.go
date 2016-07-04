package messenger

type SenderAction string

const (
	SenderActionMarkSeen SenderAction = "mark_seen"
	//SenderActionTypingOn indicator is automatically turned off after 20 seconds
	SenderActionTypingOn  SenderAction = "typing_on"
	SenderActionTypingOff SenderAction = "typing_off"
)

type rawAction struct {
	Recipient Recipient    `json:"recipient"`
	Action    SenderAction `json:"sender_action"`
}

func (m *Messenger) SendAction(recipient Recipient, action SenderAction) error {
	_, err := m.sendCustomMessage(&rawAction{Recipient: recipient, Action: action})
	return err
}

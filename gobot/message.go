package gobot

// TODO: can subtypes just all inherit from Message?

// type Message interface {
// 	ToString() (string, error)
// }

// Message represents an incoming chat message.
type Message struct {
	ID   string
	User *User
	Room string
	Text string
	*Robot
}

// TextMessage represents an incoming chat message.
type TextMessage struct {
	Message
}

// EnterMessage represents an incoming user entrance notification.
type EnterMessage struct {
	Message
}

// LeaveMessage represents an incoming user exit notification.
type LeaveMessage struct {
	Message
}

// TopicMessage represents an incoming topic change notification.
type TopicMessage struct {
	TextMessage
}

// CatchAllMessage represents an unmatched message.
type CatchAllMessage struct {
	Message
}

// Match determines if the message matches the given regex.
func (msg *TextMessage) Match(regex string) error {
	// m.text matches regex
	return nil
}

// ToString returns the message text.
func (msg *Message) ToString() (string, error) {
	return msg.Text, nil
}

func (msg *Message) Send(strings ...string) {
	msg.Robot.Send(strings...)
}

package hal

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
}

// TextMessage represents an incoming chat message.
type TextMessage Message

// EnterMessage represents an incoming user entrance notification.
type EnterMessage Message

// LeaveMessage represents an incoming user exit notification.
type LeaveMessage Message

// TopicMessage represents an incoming topic change notification.
type TopicMessage Message

// CatchAllMessage represents an unmatched message.
type CatchAllMessage Message

// Match determines if the message matches the given regex.
func (msg *Message) Match(regex string) error {
	// m.text matches regex
	return nil
}

// ToString returns the message text.
func (msg *Message) ToString() (string, error) {
	return msg.Text, nil
}

// // Send sends a response to the room
// func (msg *Message) Send(strings ...string) {
// 	msg.Robot.Send(msg, strings...)
// }

// // Reply sends a response to the user who sent the message
// func (msg *Message) Reply(strings ...string) {
// 	msg.Robot.Reply(msg, strings...)
// }

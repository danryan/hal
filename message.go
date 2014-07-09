package hal

// TODO: can subtypes just all inherit from Message?
// type Message interface {
// 	ToString() (string, error)
// }

// Message represents an incoming chat message.
type Message struct {
	ID   string
	User User
	Room string
	Text string
	Type string
}

// Match determines if the message matches the given regex.
func (msg *Message) Match(regex string) error {
	// m.text matches regex
	return nil
}

// ToString returns the message text.
func (msg *Message) String() (string, error) {
	return msg.Text, nil
}

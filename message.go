package hal

// Message represents an incoming chat message.
type Message struct {
	ID   string
	User User
	Room string
	Text string
	Type string
}

// String implements the Stringer interface
func (msg *Message) String() string {
	return msg.Text
}

package gobot

// Handler is an interface for objects to implement in order to respond
// to messages.
type Handler interface {
	Handle(msg *Message)
	// String() string
}

// HandlerFunc is an adapter that allows regular functions to act as handlers.
type HandlerFunc func(*Message)

func (fn HandlerFunc) Handle(msg *Message) {
	fn(msg)
}

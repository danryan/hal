package gobot

// Handler is an interface for objects to implement in order to respond
// to messages.
type Handler interface {
	Handle(res *Response) error
	String() string
}

// HandlerFunc is an adapter that allows regular functions to act as handlers.
type HandlerFunc func(res *Response) error

// Handle method for HandlerFunc
func (fn HandlerFunc) Handle(response *Response) error {
	if err := fn(response); err != nil {
		return err
	}
	return nil
}

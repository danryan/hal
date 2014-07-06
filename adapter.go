package hal

import (
	"fmt"
)

// Adapter interface
type Adapter interface {
	// New() (Adapter, error)
	Run() error
	Stop() error

	Receive(*Message) error
	Send(*Response, ...string) error
	Emote(*Response, ...string) error
	Reply(*Response, ...string) error
	Topic(*Response, ...string) error
	Play(*Response, ...string) error

	String() string
	Name() string
}

type adapter struct {
	name     string
	newFunc  func(*Robot) (Adapter, error)
	sendChan chan *Response
	recvChan chan *Message
	receive  func(Responder) (Message, error)
	send     func(Responder) (Response, error)
}

// Adapters is a map of registered adapters
var Adapters = map[string]adapter{}

// NewAdapter creates a new initialized adapter
func NewAdapter(robot *Robot) (Adapter, error) {
	if _, ok := Adapters[robot.AdapterName]; !ok {
		return nil, fmt.Errorf("%s is not a registered adapter", robot.AdapterName)
	}

	adapter, err := Adapters[robot.AdapterName].newFunc(robot)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}

// RegisterAdapter registers an adapter
func RegisterAdapter(name string, newFunc func(*Robot) (Adapter, error)) {
	Adapters[name] = adapter{
		name:    name,
		newFunc: newFunc,
	}
}

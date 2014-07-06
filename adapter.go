package hal

import (
	"fmt"
	"net/http"
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
}

// Adapters is a map of registered adapters
var Adapters = map[string]adapter{}

// NewAdapter creates a new initialized adapter
func NewAdapter(robot *Robot) (Adapter, error) {
	name := Config.AdapterName
	if _, ok := Adapters[name]; !ok {
		return nil, fmt.Errorf("%s is not a registered adapter", Config.AdapterName)
	}

	adapter, err := Adapters[name].newFunc(robot)
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

// BasicAdapter declares common functions shared by all adapters
type BasicAdapter struct {
	*Robot
}

func (a *BasicAdapter) SetRobot(r *Robot) {
	a.Robot = r
}

func (a *BasicAdapter) preRun() {
	Logger.Infof("Starting %s adapter.", a)
	// TODO: probably not useful for production
	Router.HandleFunc("/hal/adapter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", a)
	})
}

func (a *BasicAdapter) postRun() {
	Logger.Infof("Started %s adapter.", a)
}

func (a *BasicAdapter) run() {
	a.preRun()
	a.postRun()
}

func (a *BasicAdapter) preStop() {
	fmt.Println() // so we don't break up the log formatting :)
	Logger.Infof("Stopping %s adapter.", a)
}

func (a *BasicAdapter) postStop() {
	Logger.Infof("Stopped %s adapter.", a)
}

func (a *BasicAdapter) stop() {
	a.preStop()
	a.postStop()
}

func (a *BasicAdapter) String() string {
	return a.Robot.Adapter.Name()
}

func (a *BasicAdapter) Name() string {
	return "basic"
}

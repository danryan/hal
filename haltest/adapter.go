package haltest

import (
	"github.com/danryan/hal"
)

func init() {
	hal.RegisterAdapter("test", New)
}

type adapter struct {
	hal.BasicAdapter
}

// New returns an initialized adapter
func New(r *hal.Robot) (hal.Adapter, error) {
	a := &adapter{}
	a.SetRobot(r)
	return a, nil
}

func (a *adapter) Run() error                           { return nil }
func (a *adapter) Stop() error                          { return nil }
func (a *adapter) Receive(*hal.Message) error           { return nil }
func (a *adapter) Send(*hal.Response, ...string) error  { return nil }
func (a *adapter) Reply(*hal.Response, ...string) error { return nil }
func (a *adapter) Emote(*hal.Response, ...string) error { return nil }
func (a *adapter) Topic(*hal.Response, ...string) error { return nil }
func (a *adapter) Play(*hal.Response, ...string) error  { return nil }
func (a *adapter) Name() string                         { return "test" }

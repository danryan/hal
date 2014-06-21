package gobot

import (
	"log"
	_ "regexp"
)

// Robot receives messages from an adapter and sends them to listeners
type Robot struct {
	*Config
	Name    string
	Adapter Adapter
	Logger  *log.Logger
	// listeners []*Listener
	// handlers map[string]Handler
	handlers []Handler
	// commands map[string]Command
}

// func (r *Robot) Response() *Response {
// 	return &Response{}
// }

func (r *Robot) Handlers() []Handler {
	return r.handlers
}

// NewRobot returns a new Robot instance
func NewRobot() *Robot {
	config := NewConfig()
	return &Robot{
		Name:    config.Name,
		Adapter: config.Adapter,
		Logger:  config.Logger,
	}
}

func (r *Robot) Handle(handler ...Handler) {
	r.handlers = append(r.handlers, handler...)
}

// func (r *Robot) AddHandlers(handlers ...Handler) {
// 	for _, handler := range handlers {
// 		r.AddHandler(handler)
// 	}
// }

// func (r *Robot) AddHandler(handler Handler) {
// 	r.Handlers = append(r.Handlers, handler)
// }

// func (r *Robot) AddHandler(method string, pattern string, handler HandlerFunc) error {
// 	regex, err := regexp.Compile(pattern)

// 	if err != nil {
// 		return err
// 	}

// 	listener := &Listener{method: method, regex: regex, handler: handler}
// 	r.listeners = append(r.listeners, listener)

// 	return nil
// }

// func (r *Robot) Hear(regex string, handler HandlerFunc) error {
// 	r.AddHandler(HEAR, regex, handler)

// 	// handler := &Handler
// 	// // r.AddHandler()
// 	// listener := &Listener{Robot: r, Regex: regex, Callback: handler}
// 	// r.Listeners = append(r.Listeners, listener)
// 	// append(r.Listeners, &Listener{})
// 	return nil
// }

// func (r *Robot) Respond(regex string, fn func()) error {
// 	return nil
// }

// func (r *Robot) Topic() error {
// 	return nil
// }

// func (r *Robot) Enter() error {
// 	return nil
// }

// func (r *Robot) Leave() error {
// 	return nil
// }

func (r *Robot) Send(strings ...string) error {
	// r.Adapter.Send(env, strings)
	return nil
}

func (r *Robot) Reply() error {
	return nil
}

func (r *Robot) Close() error {
	return nil
}

func (r *Robot) Run() error {
	return nil
}

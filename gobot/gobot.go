package gobot

import (
	_ "errors"
	_ "regexp"
)

const (
	HEAR    = "HEAR"
	RESPOND = "RESPOND"
	TOPIC   = "TOPIC"
	ENTER   = "ENTER"
	LEAVE   = "LEAVE"
)

// New returns a Robot instance.
func New() *Robot {
	return NewRobot()
}

// func AddHandler(method string, pattern string, handler HandlerFunc) (*Listener, error) {
// 	regex, err := regexp.Compile(pattern)

// 	if err != nil {
// 		return nil, err
// 	}

// 	listener := &Listener{method: method, regex: regex, handler: handler}
// 	// r.listeners = append(r.listeners, listener)

// 	return listener, nil
// }

// func Respond(pattern string, handler HandlerFunc) (*Command, error) {
// 	regex, err := regexp.Compile(pattern)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Command{method: RESPOND, pattern: pattern, handler: handler}, nil
// }

// Hear a message
func Hear(pattern string, handler HandlerFunc) *Command {
	return &Command{Method: HEAR, Pattern: pattern, Handler: handler}
}

// Respond creates a new listener for Respond messages
func Respond(pattern string, handler HandlerFunc) *Command {
	return &Command{Method: RESPOND, Pattern: pattern, Handler: handler}
}

// Topic returns a new listener for Topic messages
func Topic(pattern string, handler HandlerFunc) *Command {
	return &Command{Method: TOPIC, Pattern: pattern, Handler: handler}
}

// Enter returns a new listener for Enter messages
func Enter(pattern string, handler HandlerFunc) *Command {
	return &Command{Method: ENTER, Pattern: pattern, Handler: handler}
}

// Leave creates a new listener for Leave messages
func Leave(pattern string, handler HandlerFunc) *Command {
	return &Command{Method: LEAVE, Pattern: pattern, Handler: handler}
}

func Send(msg *Message, strings ...string) error {
	// r.Adapter.Send(msg, strings)
	return nil
}

func Reply(msg *Message, strings ...string) error {
	// r.Adapter.Send(msg, strings)
	return nil
}

func Close() error {
	return nil
}

func Run() error {
	// return errors.New("This is an error")
	return nil
}

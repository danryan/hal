package hal

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
func New() (*Robot, error) {
	return NewRobot()
}

// Hear a message
func Hear(pattern string, handler HandlerFunc) *Listener {
	return &Listener{Method: HEAR, Pattern: pattern, Handler: handler}
}

// Respond creates a new listener for Respond messages
func Respond(pattern string, handler HandlerFunc) *Listener {
	return &Listener{Method: RESPOND, Pattern: pattern, Handler: handler}
}

// Topic returns a new listener for Topic messages
func Topic(pattern string, handler HandlerFunc) *Listener {
	return &Listener{Method: TOPIC, Pattern: pattern, Handler: handler}
}

// Enter returns a new listener for Enter messages
func Enter(handler HandlerFunc) *Listener {
	return &Listener{Method: ENTER, Handler: handler}
}

// Leave creates a new listener for Leave messages
func Leave(handler HandlerFunc) *Listener {
	return &Listener{Method: LEAVE, Handler: handler}
}

// Close shuts down the robot
func Close() error {
	return nil
}

package gobot

import (
	"regexp"
)

// Command struct
type Command struct {
	Method  string
	Pattern string
	Handler func(*Message)
}

// Handle implements the gobot.Handler interface
func (c *Command) Handle(msg *Message) {
	// if !c.match(msg) { return }
	c.Handler(msg)
}

func (c *Command) match(msg *Message) bool {
	text := msg.Text
	matched, _ := regexp.MatchString(c.Pattern, text)

	return matched
}

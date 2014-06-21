package gobot

import (
	"regexp"
)

// Listener receives every message from the chat source and decide whether to act on it.
// type Listener struct {
// 	Robot    *Robot
// 	Regex    string
// 	Callback func(robot *Robot, response *Response)
// 	// Matcher  string
// }
type Listener struct {
	method  string
	regex   *regexp.Regexp
	handler HandlerFunc
}

// type TextListener struct {
// 	Listener
// 	Regex string
// }

// TODO: Match should actually match something
func (l *Listener) Match(m *Message) ([]string, bool) {
	return []string{"true"}, true
	// if match := l.Match
}

// // Call checks whether the listener should
// func (l *Listener) Call(msg *Message) bool {
// 	if match, ok := l.Match(msg); ok {
// 		response := &Response{Robot: l.Robot, Message: msg, Match: match}
// 		go l.Callback(l.Robot, response)
// 		return true
// 	}
// 	return false
// }

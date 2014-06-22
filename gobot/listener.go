package gobot

import (
	// "github.com/davecgh/go-spew/spew"
	"log"
	"regexp"
)

// Listener struct
type Listener struct {
	Method  string
	Pattern string
	Handler func(res *Response) error
	regex   *regexp.Regexp
}

// Handle implements the gobot.Handler interface
func (l *Listener) Handle(res *Response) error {
	robot := res.Robot
	text := res.Message.Text

	if l.Method == RESPOND {
		l.regex = regexp.MustCompile(robot.RespondRegex(l.Pattern))
	} else {
		l.regex = regexp.MustCompile(l.Pattern)
	}
	// regex, err := regexp.Compile(pattern)

	match := l.regex.FindAllStringSubmatch(text, -1)
	if match == nil {
		log.Printf(`/%s/ did not match "%s"`, l.String(), text)
		return nil
	}
	log.Printf(`/%s/ matched "%s"`, l.String(), text)
	res.Match = match

	herr := l.Handler(res)
	if herr != nil {
		return herr
	}
	return nil
}

func (l *Listener) String() string {
	return l.Pattern
}

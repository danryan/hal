package hal

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	// RespondRegexp is used to determine whether a received message should be processed as a response
	respondRegexp = fmt.Sprintf(`^(?:@?(?:%s|%s)[:,]?)\s+(?:(.+))`, Config.Alias, Config.Name)
	// RespondRegexpTemplate expands the RespondRegexp
	respondRegexpTemplate = fmt.Sprintf(`^(?:@?(?:%s|%s)[:,]?)\s+(?:${1})`, Config.Alias, Config.Name)
)

// Listener struct
type Listener struct {
	Method  string
	Pattern string
	Handler func(res *Response) error
	regex   *regexp.Regexp
}

// Handle implements the hal.Handler interface
func (l *Listener) Handle(res *Response) error {
	text := res.Message.Text

	if l.Method == RESPOND {
		l.regex = regexp.MustCompile(strings.Replace(respondRegexpTemplate, "${1}", l.Pattern, 1))
	} else {
		l.regex = regexp.MustCompile(l.Pattern)
	}

	match := l.regex.FindAllStringSubmatch(text, -1)

	if match == nil {
		Logger.Debugf(`/%s/ did not match "%s"`, l.String(), text)
		return nil
	}
	Logger.Debugf(`/%s/ matched "%s"`, l.String(), text)
	// res.Match = match
	res.Match = match[0]

	if err := l.Handler(res); err != nil {
		return err
	}
	return nil
}

func (l *Listener) String() string {
	return l.Pattern
}

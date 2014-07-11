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

type handlerFunc func() handler

// handler is an interface for objects to implement in order to respond to messages.
type handler interface {
	Handle(res *Response) error
}

// fullHandler is an interface for objects that wish to supply their own define methods
type fullHandler interface {
	handler
	Run(*Response) error
	Usage() string
	Pattern() string
	Method() string
}

// Handlers is a map of registered handlers
var Handlers = map[string]handler{}

// Handler declares common functions shared by all handlers
type Handler struct {
	method  string
	pattern string
	usage   string
	run     func(res *Response) error
}

// Match func
func (h *Handler) Match(res *Response) bool {
	if !h.Regexp().MatchString(res.Message.Text) {
		return false
	}
	return true
}

// Handle func
func (h *Handler) Handle(res *Response) error {
	switch {
	// handle the response without matching
	case h.Pattern() == "":
		return h.Run(res)
	// handle the response after finding matches
	case h.Match(res):
		res.Match = h.Regexp().FindAllStringSubmatch(res.Message.Text, -1)[0]
		return h.Run(res)
	// if we don't find a match, return
	default:
		return nil
	}
}

// Pattern func
func (h *Handler) Pattern() string {
	return h.pattern
}

// Usage func
func (h *Handler) Usage() string {
	return h.usage
}

// Method func
func (h *Handler) Method() string {
	return h.method
}

// Run func
func (h *Handler) Run(res *Response) error {
	return h.run(res)
}

// Regexp func
func (h *Handler) Regexp() *regexp.Regexp {
	if h.Method() == RESPOND {
		return regexp.MustCompile(strings.Replace(respondRegexpTemplate, "${1}", h.Pattern(), 1))
	}
	return regexp.MustCompile(h.Pattern())
}

// NewHandler func
func NewHandler(h handler) handler {
	if fh, ok := h.(fullHandler); ok {
		return &Handler{
			pattern: fh.Pattern(),
			usage:   fh.Usage(),
			method:  fh.Method(),
			run:     fh.Run,
		}
	}
	return h
}

// BasicHandler type
type BasicHandler struct {
	Method  string
	Pattern string
	Usage   string
	Run     func(res *Response) error
}

// Handle implements the hal.Handler interface
func (h *BasicHandler) Handle(res *Response) error {
	text := res.Message.Text
	var regex *regexp.Regexp

	if h.Method == RESPOND {
		regex = regexp.MustCompile(strings.Replace(respondRegexpTemplate, "${1}", h.Pattern, 1))
	} else {
		regex = regexp.MustCompile(h.Pattern)
	}

	// assume we handle if no pattern was specified
	if h.Pattern == "" {
		return nil
	}

	if !regex.MatchString(text) {
		return nil
	}

	Logger.Debugf(`%s matched /%s/`, text, regex)
	res.Match = regex.FindAllStringSubmatch(text, -1)[0]
	return h.Run(res)
}

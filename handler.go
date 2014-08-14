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

// FullHandler declares common functions shared by all handlers
type FullHandler struct {
	method  string
	pattern string
	usage   string
	run     func(res *Response) error
}

// Handler type
type Handler struct {
	Method  string
	Pattern string
	Usage   string
	Run     func(res *Response) error
}

// Match func
func (h *FullHandler) Match(res *Response) bool {
	return handlerMatch(h.Regexp(), res.Message.Text)
}

// Match func
func (h *Handler) Match(res *Response) bool {
	return handlerMatch(h.Regexp(), res.Message.Text)
}

func handlerMatch(r *regexp.Regexp, text string) bool {
	if !r.MatchString(text) {
		return false
	}
	return true
}

// Handle func
func (h *FullHandler) Handle(res *Response) error {
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

// Handle func
func (h *Handler) Handle(res *Response) error {
	switch {
	// handle the response without matching
	case h.Pattern == "":
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
func (h *FullHandler) Pattern() string {
	return h.pattern
}

// Usage func
func (h *FullHandler) Usage() string {
	return h.usage
}

// Method func
func (h *FullHandler) Method() string {
	return h.method
}

// Run func
func (h *FullHandler) Run(res *Response) error {
	return h.run(res)
}

// Regexp func
func (h *FullHandler) Regexp() *regexp.Regexp {
	return handlerRegexp(h.Method(), h.Pattern())
}

// Regexp func
func (h *Handler) Regexp() *regexp.Regexp {
	return handlerRegexp(h.Method, h.Pattern)
}

func handlerRegexp(method, pattern string) *regexp.Regexp {
	if method == RESPOND {
		return regexp.MustCompile(strings.Replace(respondRegexpTemplate, "${1}", pattern, 1))
	}
	return regexp.MustCompile(pattern)
}

// NewHandler func
func NewHandler(h handler) handler {
	if fh, ok := h.(fullHandler); ok {
		return &FullHandler{
			pattern: fh.Pattern(),
			usage:   fh.Usage(),
			method:  fh.Method(),
			run:     fh.Run,
		}
	}
	return h
}

// BasicHandler is used to construct handlers that are low complexity and may
// not benefit from creating a custom handler type.
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

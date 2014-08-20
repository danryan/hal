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

// handler is an interface for objects to implement in order to respond to messages.
type handler interface {
	Handle(res *Response) error
}

// Handlers is a map of registered handlers
var Handlers = map[string]handler{}

func handlerMatch(r *regexp.Regexp, text string) bool {
	if !r.MatchString(text) {
		return false
	}
	return true
}

func handlerRegexp(method, pattern string) *regexp.Regexp {
	if method == RESPOND {
		return regexp.MustCompile(strings.Replace(respondRegexpTemplate, "${1}", pattern, 1))
	}
	return regexp.MustCompile(pattern)
}

// NewHandler checks whether h implements the handler interface, wrapping it in a FullHandler
func NewHandler(h interface{}) (handler, error) {
	switch v := h.(type) {
	case fullHandler:
		return &FullHandler{handler: v}, nil
	case handler:
		return v, nil
	default:
		return nil, fmt.Errorf("%v does not implement the handler interface", v)
	}
}

// Handler type
type Handler struct {
	Method  string
	Pattern string
	Usage   string
	Run     func(res *Response) error
}

// Handle func
func (h *Handler) Handle(res *Response) error {
	switch {
	// handle the response without matching
	case h.Pattern == "":
		return h.Run(res)
	// handle the response after finding matches
	case h.match(res):
		res.Match = h.regexp().FindAllStringSubmatch(res.Message.Text, -1)[0]
		return h.Run(res)
	// if we don't find a match, return
	default:
		return nil
	}
}

func (h *Handler) regexp() *regexp.Regexp {
	return handlerRegexp(h.Method, h.Pattern)
}

// Match func
func (h *Handler) match(res *Response) bool {
	return handlerMatch(h.regexp(), res.Message.Text)
}

// fullHandler is an interface for objects that wish to supply their own define methods
type fullHandler interface {
	Run(*Response) error
	Usage() string
	Pattern() string
	Method() string
}

// FullHandler declares common functions shared by all handlers
type FullHandler struct {
	handler fullHandler
}

// Regexp func
func (h *FullHandler) Regexp() *regexp.Regexp {
	return handlerRegexp(h.handler.Method(), h.handler.Pattern())
}

// Match func
func (h *FullHandler) Match(res *Response) bool {
	return handlerMatch(h.Regexp(), res.Message.Text)
}

// Handle func
func (h *FullHandler) Handle(res *Response) error {
	switch {
	// handle the response without matching
	case h.handler.Pattern() == "":
		return h.handler.Run(res)
	// handle the response after finding matches
	case h.Match(res):
		res.Match = h.Regexp().FindAllStringSubmatch(res.Text(), -1)[0]
		return h.handler.Run(res)
	// if we don't find a match, return
	default:
		return nil
	}
}

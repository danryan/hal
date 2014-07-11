package handler

import (
	"github.com/danryan/hal"
)

type echo struct {
	hal.Handler
}

// Echo exports our echo handler
func Echo() *echo { return new(echo) }

func (h *echo) Method() string {
	return hal.HEAR
}

func (h *echo) Usage() string {
	return "echo STRING - echoes STRING"
}

func (h *echo) Pattern() string {
	return `echo (.+)`
}

func (h *echo) Run(res *hal.Response) error {
	return res.Reply(res.Match[1])
}

package handler

import (
	"github.com/danryan/hal"
)

type ping struct {
	hal.Handler
}

// Ping exports
func Ping() *ping { return new(ping) }

func (h *ping) Method() string {
	return hal.RESPOND
}

func (h *ping) Usage() string {
	return `ping - responds with "PONG"`
}

func (h *ping) Pattern() string {
	return `(?:(?:p|P)ing|PING)`
}

func (h *ping) Run(res *hal.Response) error {
	return res.Send("PONG")
}

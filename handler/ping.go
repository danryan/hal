package handler

import (
	"github.com/danryan/hal"
)

type ping struct{}

func (h *ping) Method() string {
	return hal.RESPOND
}

func (h *ping) Usage() string {
	return `ping - responds with "PONG"`
}

func (h *ping) Pattern() string {
	return `(?i)ping`
}

func (h *ping) Run(res *hal.Response) error {
	return res.Send("PONG")
}

// Ping exports
var Ping = &ping{}

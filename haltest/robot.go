package haltest

import (
	"github.com/danryan/hal"
)

type Robot struct {
	hal.Robot
}

type ResponseRecorder struct {
}

func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{}
}

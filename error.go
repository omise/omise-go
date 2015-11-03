package omise

import (
	"errors"
	"strconv"
)

type ErrInternal string

func (e ErrInternal) Error() string {
	return "internal inconsistency: " + string(e)
}

type Error struct {
	Location   string `json:"location"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *Error) String() string {
	if e.StatusCode != 0 {
		return "(" + strconv.Itoa(e.StatusCode) + "/" + e.Code + ") " + e.Message
	} else {
		return "(" + e.Code + ") " + e.Message
	}
}

func (e *Error) Error() string {
	return e.String()
}

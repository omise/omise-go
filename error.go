package omise

import (
	"errors"
	"strconv"
)

var ErrInvalidKey = errors.New("invalid public or secret key.")

// ErrInternal represents internal library error. If you encounter this, it is mostly
// likely due to a bug in the omise-go library itself. Please report it by opening a new
// GitHub issue or contacting support.
type ErrInternal string

func (e ErrInternal) Error() string {
	return "internal inconsistency: " + string(e)
}

type ErrTransport struct {
	Err    error
	Buffer []byte
}

func (e ErrTransport) Error() string {
	return "transport error: " + e.Err.Error() +
		"\n with response body: " + string(e.Buffer)
}

// Error struct represents errors that may be returned from Omise's REST API. You can use
// the Code or the HTTP StatusCode field to test for the exact error condition in your
// code.
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

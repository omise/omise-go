package omise

import (
	"errors"
	"strconv"
)

// ErrInvalidKey represents missing or bad API key errors.
var ErrInvalidKey = errors.New("invalid public or secret key")

// ErrInternal represents internal library error. If you encounter this, it is mostly
// likely due to a bug in the omise-go library itself. Please report it by opening a new
// GitHub issue or contacting support.
type ErrInternal string

func (err ErrInternal) Error() string {
	return "internal inconsistency: " + string(err)
}

// ErrTransport wraps error returned by omise-go internal HTTP transport implementation.
type ErrTransport struct {
	Err    error
	Buffer []byte
}

func (err ErrTransport) Error() string {
	return "transport error: " + err.Err.Error() +
		"\n with response body: " + string(err.Buffer)
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

func (err *Error) String() string {
	if err.StatusCode == 0 {
		return "(" + err.Code + ") " + err.Message
	}

	return "(" + strconv.Itoa(err.StatusCode) + "/" + err.Code + ") " + err.Message
}

func (err *Error) Error() string {
	return err.String()
}

package main

type ErrMissingArg string

func (e ErrMissingArg) Error() string {
	return "requires: " + string(e)
}

package omise

import (
	"io"
)

type File struct {
	Name string
	Body io.Reader
}

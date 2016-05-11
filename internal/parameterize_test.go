package internal

import (
	"io/ioutil"
	"net/url"
	"testing"

	a "github.com/stretchr/testify/assert"
)

type TestRequest struct {
	op *Op `query:"-"`

	A string
	B string
}

func (req *TestRequest) Op() *Op {
	return req.op
}

func TestParameterize_GET(t *testing.T) {
	req := &TestRequest{
		op: &Op{
			Method: "GET",
			Values: url.Values{"c": []string{"gopher"}},
		},

		A: "hello",
		B: "world",
	}

	values, reader, _, e := Parameterize(req)
	a.NoError(t, e)
	a.Nil(t, reader)

	v := url.Values{}
	v.Set("a", "hello")
	v.Set("b", "world")
	v.Set("c", "gopher")
	a.Equal(t, v, values)
}

func TestParameterize_POST(t *testing.T) {
	req := &TestRequest{
		op: &Op{Method: "POST"},
		A:  "hello",
		B:  "world",
	}

	values, reader, _, e := Parameterize(req)
	a.NoError(t, e)
	a.NotNil(t, reader)
	a.Len(t, values, 0)

	buffer, e := ioutil.ReadAll(reader)
	a.NoError(t, e)
	a.Equal(t, "a=hello&b=world", string(buffer))
}

func TestParameterize_POST_Multipart(t *testing.T) {
	req := &TestRequest{
		op: &Op{Method: "POST", Multipart: true},
		A:  "hello",
		B:  "world",
	}

	values, reader, _, e := Parameterize(req)
	a.NoError(t, e)
	a.NotNil(t, reader)
	a.Len(t, values, 0)

	buffer, e := ioutil.ReadAll(reader)
	a.NoError(t, e)

	result := string(buffer)
	a.Contains(t, result, "Content-Disposition: form-data; name=\"a\"")
	a.Contains(t, result, "\r\n\r\nhello\r\n")
}

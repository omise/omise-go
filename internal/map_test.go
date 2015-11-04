package internal

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	a "github.com/stretchr/testify/assert"
)

func TestMapURLValues(t *testing.T) {
	s, b, n := "hello", true, 42

	values := url.Values{}
	check := func(struc interface{}) bool {
		result, e := MapURLValues(struc)
		if !(a.NoError(t, e) && a.True(t, reflect.DeepEqual(values, result))) {
			t.Logf("expected: %#v\nactual: %#v\n", values, result)
			return false
		}

		return true
	}

	t.Log("Basic")
	values.Set("a", "42")
	values.Set("b", "42")
	values.Set("c", "true")
	values.Set("d", "true")
	values.Set("e", s)
	values.Set("f", s)

	check(&struct {
		A int
		B *int
		C bool
		D *bool
		E string
		F *string
	}{n, &n, b, &b, s, &s})

	t.Log("Nil Coalescing")
	values = url.Values{} // should be empty

	check(&struct {
		Nil      *string
		NilAgain string
	}{nil, ""})

	t.Log("Tagged Names")
	values = url.Values{}
	values.Set("zxcv", "hello")
	values.Set("asdf", "world")

	check(&struct {
		A string `query:"zxcv"`
		B string `query:"asdf"`
	}{"hello", "world"})

	t.Log("Alised Types")
	values = url.Values{}
	values.Set("str", "hello")
	values.Set("time", "2")

	type theString string
	check(&struct {
		Str  theString
		Time time.Month
	}{"hello", 2})

	t.Log("Embedded Structs")
	values = url.Values{}
	values.Set("x", "hello")
	values.Set("outside", "world")

	type Inside struct{ X string }
	check(&struct {
		Inside
		Outside string
	}{Inside{"hello"}, "world"})
}

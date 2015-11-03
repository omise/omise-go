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

	t.Log("Basic")
	struct1 := &struct {
		A int
		B *int
		C bool
		D *bool
		E string
		F *string
	}{n, &n, b, &b, s, &s}

	values := url.Values{}
	values.Set("a", "42")
	values.Set("b", "42")
	values.Set("c", "true")
	values.Set("d", "true")
	values.Set("e", s)
	values.Set("f", s)

	result := MapURLValues(struct1)
	if !a.True(t, reflect.DeepEqual(values, result)) {
		t.Logf("expected: %#v\nactual:%#v\n", values, result)
	}

	t.Log("Nil Coalescing")
	struct2 := &struct {
		Nil      *string
		NilAgain string
	}{nil, ""}

	values = url.Values{}

	result = MapURLValues(struct2)
	if !a.True(t, reflect.DeepEqual(values, result)) {
		t.Logf("expected: %#v\nactual:%#v\n", values, result)
	}

	t.Log("Tagged Names")
	struct3 := &struct {
		A string `query:"zxcv"`
		B string `query:"asdf"`
	}{"hello", "world"}

	values = url.Values{}
	values.Set("zxcv", "hello")
	values.Set("asdf", "world")

	result = MapURLValues(struct3)
	if !a.True(t, reflect.DeepEqual(values, result)) {
		t.Logf("expected: %#v\nactual:%#v\n", values, result)
	}

	t.Log("Alised Types")
	type theString string
	struct4 := &struct {
		Str  theString
		Time time.Month
	}{"hello", 2}

	values = url.Values{}
	values.Set("str", "hello")
	values.Set("time", "2")

	result = MapURLValues(struct4)
	if !a.True(t, reflect.DeepEqual(values, result)) {
		t.Logf("expected: %#v\nactual:%#v\n", values, result)
	}
}

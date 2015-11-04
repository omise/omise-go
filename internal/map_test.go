package internal

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	a "github.com/stretchr/testify/assert"
)

func TestMapURLValues_Basic(t *testing.T) {
	s, b, n := "hello", true, 42

	v := url.Values{}
	v.Set("a", "42")
	v.Set("b", "42")
	v.Set("c", "true")
	v.Set("d", "true")
	v.Set("e", s)
	v.Set("f", s)

	check(t, v, &struct {
		A int
		B *int
		C bool
		D *bool
		E string
		F *string
	}{n, &n, b, &b, s, &s})
}

func TestMapURLValues_NilCoalese(t *testing.T) {
	v := url.Values{} // should be empty

	check(t, v, &struct {
		Nil      *string
		NilAgain string
	}{nil, ""})
}

func TestMapURLValues_TaggedNames(t *testing.T) {
	v := url.Values{}
	v.Set("zxcv", "hello")
	v.Set("asdf", "world")

	check(t, v, &struct {
		A string `query:"zxcv"`
		B string `query:"asdf"`
	}{"hello", "world"})
}

func TestMapURLValues_AlisedTypes(t *testing.T) {
	v := url.Values{}
	v.Set("str", "hello")
	v.Set("time", "2")

	type theString string
	check(t, v, &struct {
		Str  theString
		Time time.Month
	}{"hello", 2})
}

func TestMapURLValues_Structs(t *testing.T) {
	now := time.Now()

	v := url.Values{}
	v.Set("x", "hello")
	v.Set("outside", "world")
	v.Set("t", now.Format(time.RFC3339Nano))

	type Inside struct{ X string }
	check(t, v, &struct {
		Inside
		Outside string
		T       time.Time
	}{Inside{"hello"}, "world", now})
}

func check(t *testing.T, values url.Values, struc interface{}) bool {
	result, e := MapURLValues(struc)
	if !(a.NoError(t, e) && a.True(t, reflect.DeepEqual(values, result))) {
		t.Logf("expected: %#v\nactual: %#v\n", values, result)
		return false
	}

	return true
}

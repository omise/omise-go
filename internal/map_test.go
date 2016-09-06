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

func TestMapURLValues_JSONNames(t *testing.T) {
	v := url.Values{}
	v.Set("zxcv", "hello")
	v.Set("asdf", "world")

	check(t, v, &struct {
		A string `json:"zxcv"`
		B string `json:"asdf"`
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

func TestMapURLValues_Zeroes(t *testing.T) {
	v := url.Values{}
	v.Set("n1", "0")
	v.Set("f1", "0.0000")
	v.Set("altf", "0.0000")

	type theString string
	check(t, v, &struct {
		N1 int `query:",sendzero"`
		N2 int
		F1 float32 `query:",sendzero"`
		F2 float32
		F3 float64 `query:"altf,sendzero"`
		F4 float64
	}{0, 0, 0.0, 0.0, 0.0, 0.0})
}

func TestMapURLValues_StringMaps(t *testing.T) {
	v := url.Values{}
	v.Set("filter[a]", "hello")
	v.Set("filter[b]", "world")

	check(t, v, &struct {
		Filter map[string]string
	}{
		map[string]string{"a": "hello", "b": "world"},
	})
}

func TestMapURLValues_Structs(t *testing.T) {
	now := time.Now()

	v := url.Values{}
	v.Set("x", "hello")
	v.Set("named[x]", "world")
	v.Set("nested[inside][x]", "inside")
	v.Set("t", now.Format(time.RFC3339Nano))

	type Embed struct{ X string }
	type Embed2 struct{ Inside Embed }
	check(t, v, &struct {
		Embed
		Named  Embed
		Nested Embed2
		T      time.Time
		TZero  time.Time
	}{Embed{"hello"}, Embed{"world"}, Embed2{Embed{"inside"}}, now, time.Time{}})
}

func check(t *testing.T, values url.Values, struc interface{}) bool {
	result, e := MapURLValues(struc)
	if !(a.NoError(t, e) && a.True(t, reflect.DeepEqual(values, result))) {
		t.Logf("\nexpected: %#v\n  actual: %#v\n", values, result)
		return false
	}

	return true
}

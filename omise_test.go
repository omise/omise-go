package omise_test

import (
	"encoding/json"
	. "github.com/omise/omise-go"
	a "github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

type JSONRoundtripTest struct {
	srcFile string
	value   interface{}
}

func (r JSONRoundtripTest) Test(t *testing.T) {
	t.Log(reflect.TypeOf(r.value).Name())

	inbytes, e := ioutil.ReadFile("testdata/" + r.srcFile)
	if !a.NoError(t, e) {
		return
	}
	if e := json.Unmarshal(inbytes, r.value); !a.NoError(t, e) {
		return
	}

	outbytes := []byte{}
	if outbytes, e = json.Marshal(r.value); !a.NoError(t, e) {
		return
	}
	if e := json.Unmarshal(outbytes, r.value); !a.NoError(t, e) {
		return
	}

	m1, m2 := map[string]interface{}{}, map[string]interface{}{}
	if e := json.Unmarshal(inbytes, &m1); !a.NoError(t, e) {
		return
	} else if e := json.Unmarshal(outbytes, &m2); !a.NoError(t, e) {
		return
	}

	if !assertJSONEquals(t, m1, m2) {
		return
	}
}

func TestJSONRoundtrip(t *testing.T) {
	var testcases = []JSONRoundtripTest{
		{"charge_object.json", &Charge{}},
		{"card_object.json", &Card{}},
	}

	for _, test := range testcases {
		test.Test(t)
	}
}

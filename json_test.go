package omise_test

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	. "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	a "github.com/stretchr/testify/assert"
)

var JSONRoundtripTests = []JSONRoundtripTest{
	{"account_object.json", &Account{}},
	{"balance_object.json", &Balance{}},
	{"bank_account_object.json", &BankAccount{}},
	{"card_object.json", &Card{}},
	{"charge_object.json", &Charge{}},
	{"customer_object.json", &Customer{}},
	{"dispute_object.json", &Dispute{}},
	{"recipient_object.json", &Recipient{}},
	{"refund_object.json", &Refund{}},
	{"token_object.json", &Token{}},
	{"transaction_object.json", &Transaction{}},
	{"transfer_object.json", &Transfer{}},
}

type JSONRoundtripTest struct {
	srcFile string
	value   interface{}
}

func (r JSONRoundtripTest) Test(t *testing.T) {
	t.Log(reflect.ValueOf(r.value).Elem().Type().Name())

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

	if !testutil.AssertJSONEquals(t, m1, m2) {
		return
	}
}

func TestJSONRoundtrip(t *testing.T) {
	for _, test := range JSONRoundtripTests {
		test.Test(t)
	}
}

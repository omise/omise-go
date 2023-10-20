package omise_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	. "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	r "github.com/stretchr/testify/require"
)

var JSONRoundtripTests = []JSONRoundtripTest{
	{"account_object.json", &Account{}},
	{"balance_object.json", &Balance{}},
	{"bank_account_object.json", &BankAccount{}},
	{"card_object.json", &Card{}},
	{"charge_object.json", &Charge{}},
	{"customer_object.json", &Customer{}},
	{"dispute_object.json", &Dispute{}},
	{"document_object.json", &Document{}},
	{"event_object.json", &Event{}},
	{"recipient_object.json", &Recipient{}},
	{"refund_object.json", &Refund{}},
	{"schedule_object.json", &Schedule{}},
	{"token_object.json", &Token{}},
	{"transaction_object.json", &Transaction{}},
	{"transfer_object.json", &Transfer{}},
}

type JSONRoundtripTest struct {
	srcFile string
	value   interface{}
}

func (test JSONRoundtripTest) Test(t *testing.T) {
	t.Log(reflect.ValueOf(test.value).Elem().Type().Name())

	inbytes, err := os.ReadFile("testdata/objects/" + test.srcFile)
	r.NoError(t, err)

	err = json.Unmarshal(inbytes, test.value)
	r.NoError(t, err)

	outbytes, err := json.Marshal(test.value)
	r.NoError(t, err)
	err = json.Unmarshal(outbytes, test.value)
	r.NoError(t, err)

	m1, m2 := map[string]interface{}{}, map[string]interface{}{}
	err = json.Unmarshal(inbytes, &m1)
	r.NoError(t, err)
	err = json.Unmarshal(outbytes, &m2)
	r.NoError(t, err)

	testutil.AssertJSONEquals(t, m1, m2)
}

func TestJSONRoundtrip(t *testing.T) {
	for _, test := range JSONRoundtripTests {
		test.Test(t)
	}
}

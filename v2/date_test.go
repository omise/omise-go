package omise_test

import (
	"encoding/json"
	"testing"

	"github.com/omise/omise-go/v2"
	r "github.com/stretchr/testify/require"
)

func TestDateString_Marshall(t *testing.T) {
	var d omise.DateString = "2020-01-01"

	json, _ := json.Marshal(d)
	r.Equal(t, "\"2020-01-01\"", string(json))
}

func TestDateString_Empty(t *testing.T) {
	var d omise.DateString = ""

	json, _ := json.Marshal(d)
	r.Equal(t, "", string(json))
}

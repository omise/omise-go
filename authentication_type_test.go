package omise_test

import (
	"encoding/json"
	"testing"

	"github.com/omise/omise-go"
	r "github.com/stretchr/testify/require"
)

func TestAuthenticationTypeConstants(t *testing.T) {
	r.Equal(t, omise.AuthenticationType("3DS"), omise.ThreeDS)
	r.Equal(t, omise.AuthenticationType("PASSKEY"), omise.Passkey)
}

func TestAuthenticationTypeJSONMarshal(t *testing.T) {
	testCases := []struct {
		name     string
		authType omise.AuthenticationType
		expected string
	}{
		{
			name:     "ThreeDS",
			authType: omise.ThreeDS,
			expected: `"3DS"`,
		},
		{
			name:     "Passkey",
			authType: omise.Passkey,
			expected: `"PASSKEY"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.authType)
			r.NoError(t, err)
			r.Equal(t, tc.expected, string(b))
		})
	}
}

func TestAuthenticationTypeJSONUnmarshal(t *testing.T) {
	testCases := []struct {
		name     string
		json     string
		expected omise.AuthenticationType
	}{
		{
			name:     "ThreeDS",
			json:     `"3DS"`,
			expected: omise.ThreeDS,
		},
		{
			name:     "Passkey",
			json:     `"PASSKEY"`,
			expected: omise.Passkey,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var authType omise.AuthenticationType
			err := json.Unmarshal([]byte(tc.json), &authType)
			r.NoError(t, err)
			r.Equal(t, tc.expected, authType)
		})
	}
}

func TestAuthenticationTypePointer(t *testing.T) {
	// Test pointer handling for nullable field
	authType := omise.ThreeDS
	ptr := &authType

	// Test marshaling pointer
	b, err := json.Marshal(ptr)
	r.NoError(t, err)
	r.Equal(t, `"3DS"`, string(b))

	// Test unmarshaling to pointer
	var unmarshaledPtr *omise.AuthenticationType
	err = json.Unmarshal(b, &unmarshaledPtr)
	r.NoError(t, err)
	r.NotNil(t, unmarshaledPtr)
	r.Equal(t, omise.ThreeDS, *unmarshaledPtr)

	// Test null pointer
	var nullPtr *omise.AuthenticationType
	b, err = json.Marshal(nullPtr)
	r.NoError(t, err)
	r.Equal(t, "null", string(b))
}

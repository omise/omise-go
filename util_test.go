package omise_test

import (
	a "github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func testKeys() (string, string) {
	return os.Getenv("OMISE_PUBLIC_KEY"),
		os.Getenv("OMISE_SECRET_KEY")
}

func assertJSONEquals(t *testing.T, m1 map[string]interface{}, m2 map[string]interface{}) bool {
	for k, v := range m1 {
		v2, ok := m2[k]
		if !a.True(t, ok, "missing `"+k+"` key") {
			return false
		}

		if vmap, ok := v.(map[string]interface{}); ok {
			if vmap2, ok2 := v2.(map[string]interface{}); !a.True(t, ok2, "mismatched type for `"+k+"` key") {
				return false
			} else if !assertJSONEquals(t, vmap, vmap2) {
				return false
			}
		} // else not map

		if !a.Equal(t, v, v2, "mismatched value for `"+k+"` key") {
			return false
		}
	}

	return true
}

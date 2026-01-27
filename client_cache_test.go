package omise

import (
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestClientCache_ReusesClient(t *testing.T) {
	cache := NewClientCache()

	first, err := cache.Get("pkey_test_a", "skey_test_a")
	r.NoError(t, err)

	second, err := cache.Get("pkey_test_a", "skey_test_a")
	r.NoError(t, err)

	r.Same(t, first, second)
}

func TestClientCache_SeparatesKeys(t *testing.T) {
	cache := NewClientCache()

	a, err := cache.Get("pkey_test_a", "skey_test_a")
	r.NoError(t, err)

	b, err := cache.Get("pkey_test_b", "skey_test_b")
	r.NoError(t, err)

	r.NotSame(t, a, b)
}

func TestClientCache_InvalidKey(t *testing.T) {
	cache := NewClientCache()

	_, err := cache.Get("invalid", "skey_test_a")
	r.Error(t, err)
}

func TestClientCache_Delete(t *testing.T) {
	cache := NewClientCache()

	first, err := cache.Get("pkey_test_del", "skey_test_del")
	r.NoError(t, err)

	cache.Delete("pkey_test_del", "skey_test_del")

	second, err := cache.Get("pkey_test_del", "skey_test_del")
	r.NoError(t, err)

	r.NotSame(t, first, second)
}

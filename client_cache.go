package omise

import "sync"

// ClientCache provides an optional, concurrency-safe cache of Clients keyed by
// their (pkey, skey) pair. It is additive and does not change the behavior of
// NewClient; callers opt in when they want reuse without mutating shared
// clients.
type ClientCache struct {
	cache sync.Map
}

// NewClientCache constructs an empty cache.
func NewClientCache() *ClientCache {
	return &ClientCache{}
}

// Get returns a cached Client for the provided key pair or creates one with
// NewClient and stores it for future reuse.
func (c *ClientCache) Get(pkey, skey string) (*Client, error) {
	key := pkey + "|" + skey

	if existing, ok := c.cache.Load(key); ok {
		return existing.(*Client), nil
	}

	client, err := NewClient(pkey, skey)
	if err != nil {
		return nil, err
	}

	actual, _ := c.cache.LoadOrStore(key, client)
	return actual.(*Client), nil
}

// Delete removes a cached Client for the provided key pair, allowing a fresh
// instance to be created on the next Get.
func (c *ClientCache) Delete(pkey, skey string) {
	key := pkey + "|" + skey
	c.cache.Delete(key)
}

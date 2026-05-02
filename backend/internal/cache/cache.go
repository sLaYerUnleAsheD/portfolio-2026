// Package cache provides a simple, thread-safe, in-memory key-value cache
// with time-to-live (TTL) expiry. Designed for caching expensive API
// responses like the daily art-of-the-day image.
package cache

import (
	"sync"
	"time"
)

// entry represents a single cached item with its value and expiry time.
type entry struct {
	value     interface{}
	expiresAt time.Time
}

// Cache is a concurrent-safe in-memory store with per-key TTL expiry.
type Cache struct {
	mu      sync.RWMutex
	items   map[string]entry
	defaultTTL time.Duration
}

// New creates a new Cache with the given default TTL.
// Items expire after the TTL has elapsed since they were set.
func New(defaultTTL time.Duration) *Cache {
	return &Cache{
		items:      make(map[string]entry),
		defaultTTL: defaultTTL,
	}
}

// Get retrieves a value from the cache. Returns the value and true if the
// key exists and has not expired; otherwise returns nil and false.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e, exists := c.items[key]
	if !exists {
		return nil, false
	}

	// Check if the entry has expired
	if time.Now().After(e.expiresAt) {
		// Expired — the cleanup happens lazily on the next Set or explicit call
		return nil, false
	}

	return e.value, true
}

// Set stores a value in the cache with the default TTL.
func (c *Cache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, c.defaultTTL)
}

// SetWithTTL stores a value in the cache with a custom TTL.
func (c *Cache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = entry{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}

// Delete removes a key from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// Cleanup removes all expired entries from the cache.
// Call this periodically if memory pressure is a concern.
func (c *Cache) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, e := range c.items {
		if now.After(e.expiresAt) {
			delete(c.items, key)
		}
	}
}

package activecache

import (
	"container/list"
	"time"
)

// Get function gets the value of the key from the cache if no value is found it returns []byte
// Note:it does not panic if key is not found
func (c *activeTTLCache) Get(key []byte) ([]byte, time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element := c.data.Get(key); element != nil {
		entry := element.Value.(*entry)
		if entry.expiry.IsZero() || entry.expiry.After(time.Now()) {
			c.eviction.MoveToFront(entry.frequency)
			return entry.value, entry.ttl
		} else {
			c.removeEntry(element)
		}
	}

	return nil, time.Duration(0)
}

// Set function sets the value and TTL with reference with the key in the cache.
func (c *activeTTLCache) Set(key, value []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expiry time.Time
	if ttl > 0 {
		expiry = time.Now().Add(ttl)
	}

	if element := c.data.Get(key); element != nil {
		entry := element.Value.(*entry)
		entry.value = value
		entry.expiry = expiry
		c.eviction.MoveToFront(entry.frequency)
	} else {
		element := c.eviction.PushFront(&entry{key, value, expiry, ttl, nil})
		c.data.Set(key, element)
		element.Value.(*entry).frequency = element
	}

	c.evictExpired()
}

// evictExpired actively removes expired entries from the catch.
func (c *activeTTLCache) evictExpired() {
	now := time.Now()
	for {
		element := c.eviction.Back()
		if element == nil {
			break
		}

		entry := element.Value.(*entry)
		if entry.expiry.IsZero() || entry.expiry.After(now) {
			break
		}

		c.removeEntry(element)
	}
}

func (c *activeTTLCache) removeEntry(element *list.Element) {
	entry := element.Value.(*entry)
	// remove element from custom hashmap
	c.data.Remove(entry.key)
	//remove element from list
	c.eviction.Remove(element)
}

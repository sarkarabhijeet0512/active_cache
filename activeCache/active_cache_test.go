package activecache

import (
	"testing"
	"time"
)

func BenchmarkActiveTTLCache(b *testing.B) {
	cache := NewCache()
	key := []byte("key")
	value := []byte("value")
	ttl := time.Second

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set(key, value, ttl)
		cache.Get(key)
	}
}

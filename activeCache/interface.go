package activecache

import (
	customHashTable "audigent_task/custom_hash_table"
	"audigent_task/custom_hash_table/hasher"
	"container/list"
	"time"
)

type Cache interface {
	// Set will store the key value pair with a given TTL.
	Set(key, value []byte, ttl time.Duration)

	// Get returns the value stored using `key`.
	//
	// If the key is not present value will be set to nil.
	Get(key []byte) (value []byte, ttl time.Duration)
}

func NewCache() Cache {
	return &activeTTLCache{
		data:     customHashTable.NewHashTable(bucketSize, loadFactor, hasher.Djb2),
		eviction: list.New(),
	}
}

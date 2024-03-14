package activecache

import (
	customHashTable "audigent_task/custom_hash_table"
	"container/list"
	"sync"
	"time"
)

var bucketSize = 4

// %load at which the bucket is resized/doubled and rehashed
var loadFactor = 80

type (
	activeTTLCache struct {
		mu       sync.Mutex
		data     *customHashTable.HashTable
		eviction *list.List
	}

	entry struct {
		key       []byte
		value     []byte
		expiry    time.Time
		ttl       time.Duration
		frequency *list.Element
	}
)

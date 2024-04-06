package customhashtable

import (
	"bytes"
	"container/list"
)

type KeyVal struct {
	Key   []byte
	Value *list.Element
}

type HashTable struct {
	BucketSize int
	FilledSize int
	Bucket     [][]KeyVal
	LoadFactor int
	HashFunc   func([]byte) int
}

// NewHashTable function creates a new hash table with the specified parameters and returns a pointer to it.
//
// Parameters:
//
// BucketSize: Initializes the hash table with the specified number of buckets, determining the initial capacity.
//
// Bucket: Initializes the array of buckets where key-value pairs will be stored. Each bucket is a slice of KeyVal structs.
//
// FilledSize: Initializes the count of key-value pairs currently stored in the hash table.
//
// LoadFactor: Specifies the threshold at which the hash table will be resized to accommodate more key-value pairs. The load factor is a ratio of the number of stored key-value pairs to the number of buckets.
//
// HashFunc: Sets the hash function used to calculate the index for storing and retrieving key-value pairs in the hash table.
func NewHashTable(BucketSize int, LoadFactor int, HashFunc func([]byte) int) *HashTable {
	return &HashTable{
		BucketSize: BucketSize,
		Bucket:     make([][]KeyVal, BucketSize),
		FilledSize: 0,
		LoadFactor: LoadFactor,
		HashFunc:   HashFunc,
	}
}

// _hash calculates the hash value for a given key and returns the corresponding bucket index.
func (ht *HashTable) _hash(Key []byte) int {
	return ht.HashFunc(Key) % ht.BucketSize
}

// Get retrieves the value associated with the given key from the hash table.
func (ht *HashTable) Get(Key []byte) *list.Element {
	hash := ht._hash(Key)

	for _, v := range ht.Bucket[hash] {
		if bytes.Equal(v.Key, Key) {
			return v.Value
		}
	}
	return nil
}

// Set inserts a new key-value pair into the hash table or updates the value if the key already exists.
func (ht *HashTable) Set(Key []byte, Value *list.Element) {
	hash := ht._hash(Key)

	for i, v := range ht.Bucket[hash] {
		if bytes.Equal(v.Key, Key) {
			ht.Bucket[hash][i].Value = Value
			return
		}
	}

	ht.FilledSize++
	ht.Bucket[hash] = append(ht.Bucket[hash], KeyVal{Key: Key, Value: Value})

	// Resize if load factor exceeded
	if ht.FilledSize*100/ht.BucketSize >= ht.LoadFactor {
		ht.resizeAndRehash()
	}
}

// resizeAndRehash resizes the hash table and rehashes all existing key-value pairs when the load factor is exceeded.

func (ht *HashTable) resizeAndRehash() {
	ht.BucketSize *= 2
	newBucket := make([][]KeyVal, ht.BucketSize)

	for _, bucket := range ht.Bucket {
		for _, kv := range bucket {
			hash := ht._hash(kv.Key)
			newBucket[hash] = append(newBucket[hash], kv)
		}
	}
	ht.Bucket = newBucket
}

// Remove deletes the key-value pair associated with the given key from the hash table.
//
// Parameters:
// - Key: The key of the pair to be removed.
//
// Returns:
// - True if the key was found and removed successfully, false otherwise.
func (ht *HashTable) Remove(Key []byte) bool {
	hash := ht._hash(Key)

	for i, v := range ht.Bucket[hash] {
		if bytes.Equal(v.Key, Key) {
			ht.Bucket[hash] = append(ht.Bucket[hash][:i], ht.Bucket[hash][i+1:]...)
			ht.FilledSize--
			return true
		}
	}
	return false
}

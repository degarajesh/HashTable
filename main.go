package main

import (
	"fmt"
	"hash/fnv"
)

// Entry struct represents a single entry in a hash table
type Entry struct {
	Key   string
	Value string
}

// HashTable struct represents a hash table
type HashTable struct {
	Buckets []Entry
	Size    int
}

// NewHashTable creates a new hash table of given size
func NewHashTable(size int) *HashTable {
	return &HashTable{
		Buckets: make([]Entry, size),
		Size:    size,
	}
}

// Hash computes a hash value for a given string
func (h *HashTable) Hash(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32()) % h.Size
}

// Insert adds a new entry to the hash table
func (h *HashTable) Insert(key string, value string) {
	index := h.Hash(key)
	h.Buckets[index] = Entry{Key: key, Value: value}
}

// Delete removes an entry from the hash table
func (h *HashTable) Delete(key string) {
	index := h.Hash(key)
	h.Buckets[index] = Entry{}
}

// Get retrieves a value from the hash table by its key
func (h *HashTable) Get(key string) (string, bool) {
	index := h.Hash(key)
	if h.Buckets[index].Key == "" {
		return "", false
	}
	return h.Buckets[index].Value, true
}

func main() {
	hashTable := NewHashTable(100)
	hashTable.Insert("key1", "value1")
	hashTable.Insert("key2", "value2")

	value, exists := hashTable.Get("key1")
	if exists {
		fmt.Println(value)
	}

	hashTable.Delete("key1")

	value, exists = hashTable.Get("key1")
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Value not found")
	}
}

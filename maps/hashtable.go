package maps

import (
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/chiranjeevipavurala/gocollections/lists"
	"github.com/chiranjeevipavurala/gocollections/sets"
)

// HashTable is a thread-safe implementation of a hash table.
// Unlike HashMap, HashTable provides thread-safe operations for concurrent access.
// Note: This implementation allows zero values for both keys and values.
// It is the responsibility of the user to handle zero values appropriately for their use case.
// For example, if zero values are not desired, the user should implement their own checks
// before calling Put, PutIfAbsent, or Replace operations.
type HashTable[K comparable, V comparable] struct {
	items map[K]V
	mu    sync.RWMutex
}

// NewHashTable creates a new HashTable.
func NewHashTable[K comparable, V comparable]() *HashTable[K, V] {
	return &HashTable[K, V]{
		items: make(map[K]V),
	}
}

// Put associates the specified value with the specified key in this map.
func (ht *HashTable[K, V]) Put(key K, value V) V {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	oldValue := ht.items[key]
	ht.items[key] = value
	return oldValue
}

// Get returns the value associated with the specified key.
func (ht *HashTable[K, V]) Get(key K) *V {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	if value, exists := ht.items[key]; exists {
		return &value
	}
	return nil
}

// Remove removes the mapping for a key from this map if it is present.
func (ht *HashTable[K, V]) Remove(key K) V {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	oldValue := ht.items[key]
	delete(ht.items, key)
	return oldValue
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (ht *HashTable[K, V]) ContainsKey(key K) bool {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	_, exists := ht.items[key]
	return exists
}

// ContainsValue returns true if this map maps one or more keys to the specified value.
func (ht *HashTable[K, V]) ContainsValue(value V) bool {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	for _, v := range ht.items {
		if v == value {
			return true
		}
	}
	return false
}

// Size returns the number of key-value mappings in this map.
func (ht *HashTable[K, V]) Size() int {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	return len(ht.items)
}

// IsEmpty returns true if this map contains no key-value mappings.
func (ht *HashTable[K, V]) IsEmpty() bool {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	return len(ht.items) == 0
}

// Clear removes all of the mappings from this map.
func (ht *HashTable[K, V]) Clear() {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	ht.items = make(map[K]V)
}

// EntrySet returns a set of all entries in the map.
func (ht *HashTable[K, V]) EntrySet() collections.Set[collections.MapEntry[K, V]] {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	entries := make([]collections.MapEntry[K, V], 0, len(ht.items))
	for k, v := range ht.items {
		entries = append(entries, collections.NewHashMapEntry(k, v))
	}
	set := sets.NewHashSet[collections.MapEntry[K, V]]()
	for _, entry := range entries {
		set.Add(entry)
	}
	return set
}

// KeySet returns a set of all keys in the map.
func (ht *HashTable[K, V]) KeySet() collections.Set[K] {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	keys := make([]K, 0, len(ht.items))
	for k := range ht.items {
		keys = append(keys, k)
	}
	set := sets.NewHashSet[K]()
	for _, key := range keys {
		set.Add(key)
	}
	return set
}

// Values returns a collection of all values in the map.
func (ht *HashTable[K, V]) Values() collections.Collection[V] {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	values := make([]V, 0, len(ht.items))
	for _, v := range ht.items {
		values = append(values, v)
	}
	list := lists.NewArrayList[V]()
	for _, value := range values {
		list.Add(value)
	}
	return list
}

// Equals returns true if this map equals the given map.
func (ht *HashTable[K, V]) Equals(other any) bool {
	if other == nil {
		return false
	}
	otherMap, ok := other.(collections.Map[K, V])
	if !ok {
		return false
	}

	ht.mu.RLock()
	defer ht.mu.RUnlock()

	if ht.Size() != otherMap.Size() {
		return false
	}

	for k, v := range ht.items {
		otherValue := otherMap.Get(k)
		if otherValue == nil || *otherValue != v {
			return false
		}
	}
	return true
}

// PutIfAbsent associates the specified value with the specified key in this map if the key is not already associated with a value.
func (ht *HashTable[K, V]) PutIfAbsent(key K, value V) V {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	if existingValue, exists := ht.items[key]; exists {
		return existingValue
	}
	ht.items[key] = value
	var zeroValue V
	return zeroValue
}

// Replace replaces the entry for the specified key only if it is currently mapped to some value.
func (ht *HashTable[K, V]) Replace(key K, value V) V {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	if oldValue, exists := ht.items[key]; exists {
		ht.items[key] = value
		return oldValue
	}
	var zeroValue V
	return zeroValue
}

// ReplaceKeyWithValue replaces the entry for the specified key only if currently mapped to the given old value.
func (ht *HashTable[K, V]) ReplaceKeyWithValue(key K, oldValue V, newValue V) bool {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	if currentValue, exists := ht.items[key]; exists && currentValue == oldValue {
		ht.items[key] = newValue
		return true
	}
	return false
}

// RemoveKeyWithValue removes the entry for the specified key only if it is currently mapped to the specified value.
func (ht *HashTable[K, V]) RemoveKeyWithValue(key K, value V) bool {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	if currentValue, exists := ht.items[key]; exists && currentValue == value {
		delete(ht.items, key)
		return true
	}
	return false
}

// PutAll copies all of the mappings from the specified map to this map.
func (ht *HashTable[K, V]) PutAll(m collections.Map[K, V]) {
	if m == nil {
		return
	}

	ht.mu.Lock()
	defer ht.mu.Unlock()

	for _, entry := range m.EntrySet().ToArray() {
		ht.items[entry.GetKey()] = entry.GetValue()
	}
}

// HasKey returns true if this map contains a mapping for the specified key.
func (ht *HashTable[K, V]) HasKey(key K) bool {
	return ht.ContainsKey(key)
}

// HasValue returns true if this map maps one or more keys to the specified value.
func (ht *HashTable[K, V]) HasValue(value V) bool {
	return ht.ContainsValue(value)
}

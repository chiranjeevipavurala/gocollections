package maps

import (
	"errors"
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/chiranjeevipavurala/gocollections/sets"
)

// node represents a key-value pair in the linked list.
type node[K comparable, V comparable] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

// LinkedHashMap is a Map that maintains insertion order.
type LinkedHashMap[K comparable, V comparable] struct {
	head  *node[K, V]
	tail  *node[K, V]
	items map[K]*node[K, V]
	mu    sync.RWMutex
}

// NewLinkedHashMap creates a new LinkedHashMap.
func NewLinkedHashMap[K comparable, V comparable]() collections.Map[K, V] {
	return &LinkedHashMap[K, V]{
		items: make(map[K]*node[K, V]),
	}
}

// Clear removes all mappings from this map.
func (lhm *LinkedHashMap[K, V]) Clear() {
	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	lhm.head = nil
	lhm.tail = nil
	lhm.items = make(map[K]*node[K, V])
}

// HasKey returns true if this map contains a mapping for the specified key.
func (lhm *LinkedHashMap[K, V]) HasKey(key K) bool {
	lhm.mu.RLock()
	defer lhm.mu.RUnlock()

	_, exists := lhm.items[key]
	return exists
}

// HasValue returns true if this map maps one or more keys to the specified value.
func (lhm *LinkedHashMap[K, V]) HasValue(value V) bool {
	lhm.mu.RLock()
	defer lhm.mu.RUnlock()

	for current := lhm.head; current != nil; current = current.next {
		if current.value == value {
			return true
		}
	}
	return false
}

// EntrySet returns a Set view of the mappings contained in this map.
func (lhm *LinkedHashMap[K, V]) EntrySet() collections.Set[collections.MapEntry[K, V]] {
	// Collect entries first
	entries := make([]collections.MapEntry[K, V], 0)
	lhm.mu.RLock()
	for current := lhm.head; current != nil; current = current.next {
		entries = append(entries, collections.NewHashMapEntry(current.key, current.value))
	}
	lhm.mu.RUnlock()

	// Create set without holding lock
	set := sets.NewHashSet[collections.MapEntry[K, V]]()
	if hashSet, ok := set.(*sets.HashSet[collections.MapEntry[K, V]]); ok {
		hashSet.AddAllBatch(entries)
	}
	return set
}

// Equals returns true if this map is equal to the specified object.
func (lhm *LinkedHashMap[K, V]) Equals(obj any) bool {
	if obj == nil {
		return false
	}
	mapObj, ok := obj.(collections.Map[K, V])
	if !ok {
		return false
	}

	// Get size before acquiring lock
	if lhm.Size() != mapObj.Size() {
		return false
	}

	// Get all our entries before acquiring lock
	entries := make([]struct {
		key   K
		value V
	}, 0)
	lhm.mu.RLock()
	for current := lhm.head; current != nil; current = current.next {
		entries = append(entries, struct {
			key   K
			value V
		}{current.key, current.value})
	}
	lhm.mu.RUnlock()

	// Check values without holding our lock
	for _, entry := range entries {
		value := mapObj.Get(entry.key)
		if value == nil || *value != entry.value {
			return false
		}
	}
	return true
}

// Get returns the value to which the specified key is mapped.
func (lhm *LinkedHashMap[K, V]) Get(key K) *V {
	lhm.mu.RLock()
	defer lhm.mu.RUnlock()

	if node, exists := lhm.items[key]; exists {
		return &node.value
	}
	return nil
}

// IsEmpty returns true if this map contains no key-value mappings.
func (lhm *LinkedHashMap[K, V]) IsEmpty() bool {
	lhm.mu.RLock()
	defer lhm.mu.RUnlock()

	return len(lhm.items) == 0
}

// KeySet returns a Set view of the keys contained in this map.
func (lhm *LinkedHashMap[K, V]) KeySet() collections.Set[K] {
	// Collect keys first
	keys := make([]K, 0)
	lhm.mu.RLock()
	for current := lhm.head; current != nil; current = current.next {
		keys = append(keys, current.key)
	}
	lhm.mu.RUnlock()

	// Create set without holding lock
	set := sets.NewHashSet[K]()
	if hashSet, ok := set.(*sets.HashSet[K]); ok {
		hashSet.AddAllBatch(keys)
	}
	return set
}

// Put associates the specified value with the specified key in this map.
func (lhm *LinkedHashMap[K, V]) Put(key K, value V) V {
	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	if existingNode, exists := lhm.items[key]; exists {
		oldValue := existingNode.value
		existingNode.value = value
		return oldValue
	}

	newNode := &node[K, V]{
		key:   key,
		value: value,
	}
	lhm.items[key] = newNode

	if lhm.head == nil {
		lhm.head = newNode
		lhm.tail = newNode
	} else {
		newNode.prev = lhm.tail
		lhm.tail.next = newNode
		lhm.tail = newNode
	}

	var zero V
	return zero
}

// PutAll copies all of the mappings from the specified map to this map.
func (lhm *LinkedHashMap[K, V]) PutAll(m collections.Map[K, V]) {
	if m == nil {
		return
	}

	// If the source map is also a LinkedHashMap, use ForEachEntry to preserve order
	if sourceLinkedMap, ok := m.(*LinkedHashMap[K, V]); ok {
		lhm.mu.Lock()
		defer lhm.mu.Unlock()

		sourceLinkedMap.ForEachEntry(func(key K, value V) {
			if existingNode, exists := lhm.items[key]; exists {
				existingNode.value = value
			} else {
				newNode := &node[K, V]{
					key:   key,
					value: value,
				}
				lhm.items[key] = newNode
				if lhm.head == nil {
					lhm.head = newNode
					lhm.tail = newNode
				} else {
					newNode.prev = lhm.tail
					lhm.tail.next = newNode
					lhm.tail = newNode
				}
			}
		})
		return
	}

	// For other map types, use the original approach
	// Get entries before acquiring lock
	entries := m.EntrySet().ToArray()

	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	for _, entry := range entries {
		key := entry.GetKey()
		value := entry.GetValue()

		if existingNode, exists := lhm.items[key]; exists {
			existingNode.value = value
		} else {
			newNode := &node[K, V]{
				key:   key,
				value: value,
			}
			lhm.items[key] = newNode
			if lhm.head == nil {
				lhm.head = newNode
				lhm.tail = newNode
			} else {
				newNode.prev = lhm.tail
				lhm.tail.next = newNode
				lhm.tail = newNode
			}
		}
	}
}

// PutIfAbsent associates the specified value with the specified key in this map if the key is not already associated with a value.
func (lhm *LinkedHashMap[K, V]) PutIfAbsent(key K, value V) V {
	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	if existingNode, exists := lhm.items[key]; exists {
		return existingNode.value
	}

	newNode := &node[K, V]{
		key:   key,
		value: value,
	}
	lhm.items[key] = newNode

	if lhm.head == nil {
		lhm.head = newNode
		lhm.tail = newNode
	} else {
		newNode.prev = lhm.tail
		lhm.tail.next = newNode
		lhm.tail = newNode
	}

	var zero V
	return zero
}

// Remove removes the mapping for a key from this map if it is present.
func (lhm *LinkedHashMap[K, V]) Remove(key K) V {
	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	if existingNode, exists := lhm.items[key]; exists {
		if existingNode.prev != nil {
			existingNode.prev.next = existingNode.next
		} else {
			lhm.head = existingNode.next
		}
		if existingNode.next != nil {
			existingNode.next.prev = existingNode.prev
		} else {
			lhm.tail = existingNode.prev
		}
		delete(lhm.items, key)
		return existingNode.value
	}

	var zero V
	return zero
}

// RemoveKeyWithValue removes the entry for the specified key only if it is currently mapped to the specified value.
func (lhm *LinkedHashMap[K, V]) RemoveKeyWithValue(key K, value V) bool {
	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	if existingNode, exists := lhm.items[key]; exists && existingNode.value == value {
		if existingNode.prev != nil {
			existingNode.prev.next = existingNode.next
		} else {
			lhm.head = existingNode.next
		}
		if existingNode.next != nil {
			existingNode.next.prev = existingNode.prev
		} else {
			lhm.tail = existingNode.prev
		}
		delete(lhm.items, key)
		return true
	}
	return false
}

// Replace replaces the entry for the specified key only if it is currently mapped to some value.
func (lhm *LinkedHashMap[K, V]) Replace(key K, value V) V {
	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	if existingNode, exists := lhm.items[key]; exists {
		oldValue := existingNode.value
		existingNode.value = value
		return oldValue
	}

	var zero V
	return zero
}

// ReplaceKeyWithValue replaces the entry for the specified key only if currently mapped to the given old value.
func (lhm *LinkedHashMap[K, V]) ReplaceKeyWithValue(key K, oldValue V, newValue V) bool {
	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	if existingNode, exists := lhm.items[key]; exists && existingNode.value == oldValue {
		existingNode.value = newValue
		return true
	}
	return false
}

// Size returns the number of key-value mappings in this map.
func (lhm *LinkedHashMap[K, V]) Size() int {
	lhm.mu.RLock()
	defer lhm.mu.RUnlock()

	return len(lhm.items)
}

// Values returns a Collection view of the values contained in this map.
func (lhm *LinkedHashMap[K, V]) Values() collections.Collection[V] {
	// Collect values first
	values := make([]V, 0)
	lhm.mu.RLock()
	for current := lhm.head; current != nil; current = current.next {
		values = append(values, current.value)
	}
	lhm.mu.RUnlock()

	// Create set without holding lock
	set := sets.NewHashSet[V]()
	if hashSet, ok := set.(*sets.HashSet[V]); ok {
		hashSet.AddAllBatch(values)
	}
	return set
}

// GetOrDefault returns the value to which the specified key is mapped, or defaultValue if this map contains no mapping for the key.
func (lhm *LinkedHashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	lhm.mu.RLock()
	defer lhm.mu.RUnlock()

	if existingNode, exists := lhm.items[key]; exists {
		return existingNode.value
	}
	return defaultValue
}

// ForEachEntry performs the given action for each entry in this map.
func (lhm *LinkedHashMap[K, V]) ForEachEntry(action func(key K, value V)) {
	// Collect entries first
	entries := make([]struct {
		key   K
		value V
	}, 0)
	lhm.mu.RLock()
	for current := lhm.head; current != nil; current = current.next {
		entries = append(entries, struct {
			key   K
			value V
		}{current.key, current.value})
	}
	lhm.mu.RUnlock()

	// Execute callback without holding lock
	for _, entry := range entries {
		action(entry.key, entry.value)
	}
}

// ComputeIfAbsent computes a value for the specified key if the key is not already associated with a value.
func (lhm *LinkedHashMap[K, V]) ComputeIfAbsent(key K, mappingFunction func(K) V) (V, error) {
	if mappingFunction == nil {
		var zero V
		return zero, errors.New("mapping function cannot be nil")
	}

	lhm.mu.Lock()
	defer lhm.mu.Unlock()

	if existingNode, exists := lhm.items[key]; exists {
		return existingNode.value, nil
	}

	value := mappingFunction(key)
	newNode := &node[K, V]{
		key:   key,
		value: value,
	}
	lhm.items[key] = newNode

	if lhm.head == nil {
		lhm.head = newNode
		lhm.tail = newNode
	} else {
		newNode.prev = lhm.tail
		lhm.tail.next = newNode
		lhm.tail = newNode
	}

	return value, nil
}

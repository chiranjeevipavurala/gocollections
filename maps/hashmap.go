package maps

import (
	"errors"
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
	"github.com/chiranjeevipavurala/gocollections/sets"
)

// DefaultCapacity is the default initial capacity for HashMap
const DefaultCapacity = 16

// LoadFactor is the factor at which the map will be resized
const LoadFactor = 0.75

type HashMap[K comparable, V comparable] struct {
	entries map[K]V
	mu      sync.RWMutex
}

func NewHashMap[K comparable, V comparable]() collections.Map[K, V] {
	return &HashMap[K, V]{
		entries: make(map[K]V, DefaultCapacity),
		mu:      sync.RWMutex{},
	}
}

func NewHashMapWithCapacity[K comparable, V comparable](capacity int) collections.Map[K, V] {
	if capacity <= 0 {
		capacity = DefaultCapacity
	}
	return &HashMap[K, V]{
		entries: make(map[K]V, capacity),
		mu:      sync.RWMutex{},
	}
}

func (h *HashMap[K, V]) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.entries = make(map[K]V, DefaultCapacity)
}

func (h *HashMap[K, V]) ContainsKey(key K) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	_, ok := h.entries[key]
	return ok
}

func (h *HashMap[K, V]) ContainsValue(value V) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, val := range h.entries {
		if val == value {
			return true
		}
	}
	return false
}

func (h *HashMap[K, V]) EntrySet() collections.Set[collections.MapEntry[K, V]] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	set := sets.NewHashSet[collections.MapEntry[K, V]]()
	for key, val := range h.entries {
		entry := collections.NewHashMapEntry(key, val)
		_ = set.Add(entry)
	}
	return set
}

func (h *HashMap[K, V]) Equals(obj any) bool {
	if obj == nil {
		return false
	}
	mapObj, ok := obj.(collections.Map[K, V])
	if !ok {
		return false
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.entries) != mapObj.Size() {
		return false
	}

	// Create a copy of entries to avoid holding the lock during comparison
	entries := make(map[K]V, len(h.entries))
	for k, v := range h.entries {
		entries[k] = v
	}

	// Release the lock before comparing entries
	h.mu.RUnlock()
	defer h.mu.RLock()

	for k, v := range entries {
		if !mapObj.HasKey(k) || *mapObj.Get(k) != v {
			return false
		}
	}
	return true
}

func (h *HashMap[K, V]) Get(key K) *V {
	h.mu.RLock()
	defer h.mu.RUnlock()

	val, ok := h.entries[key]
	if !ok {
		return nil
	}
	return &val
}

func (h *HashMap[K, V]) IsEmpty() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.entries) == 0
}

func (h *HashMap[K, V]) KeySet() collections.Set[K] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	set := sets.NewHashSet[K]()
	for key := range h.entries {
		_ = set.Add(key)
	}
	return set
}

func (h *HashMap[K, V]) Put(key K, value V) V {
	h.mu.Lock()
	defer h.mu.Unlock()

	oldValue := h.entries[key]
	h.entries[key] = value
	return oldValue
}

func (h *HashMap[K, V]) PutAll(m collections.Map[K, V]) {
	if m == nil {
		return
	}

	// Get all entries first to minimize lock time
	entries := m.EntrySet().ToArray()

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, entry := range entries {
		h.entries[entry.GetKey()] = entry.GetValue()
	}
}

func (h *HashMap[K, V]) Remove(key K) V {
	h.mu.Lock()
	defer h.mu.Unlock()

	oldValue := h.entries[key]
	delete(h.entries, key)
	return oldValue
}

func (h *HashMap[K, V]) RemoveKeyWithValue(key K, value V) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	val, ok := h.entries[key]
	if !ok || val != value {
		return false
	}
	delete(h.entries, key)
	return true
}

func (h *HashMap[K, V]) Replace(key K, value V) V {
	h.mu.Lock()
	defer h.mu.Unlock()

	oldValue := h.entries[key]
	if _, ok := h.entries[key]; ok {
		h.entries[key] = value
	}
	return oldValue
}

func (h *HashMap[K, V]) ReplaceKeyWithValue(key K, oldValue V, newValue V) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	val, ok := h.entries[key]
	if !ok || val != oldValue {
		return false
	}
	h.entries[key] = newValue
	return true
}

func (h *HashMap[K, V]) Size() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.entries)
}

func (h *HashMap[K, V]) Values() collections.Collection[V] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	collection := sets.NewHashSet[V]()
	for _, val := range h.entries {
		_ = collection.Add(val)
	}
	return collection
}

func (h *HashMap[K, V]) PutIfAbsent(key K, value V) V {
	h.mu.Lock()
	defer h.mu.Unlock()

	if oldValue, ok := h.entries[key]; ok {
		return oldValue
	}
	h.entries[key] = value
	var zero V
	return zero
}

func (h *HashMap[K, V]) HasKey(key K) bool {
	return h.ContainsKey(key)
}

func (h *HashMap[K, V]) HasValue(value V) bool {
	return h.ContainsValue(value)
}

// GetOrDefault returns the value to which the specified key is mapped, or defaultValue if not mapped
func (h *HashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	h.mu.RLock()
	defer h.mu.RUnlock()

	value, ok := h.entries[key]
	if !ok {
		return defaultValue
	}
	return value
}

// ForEachEntry performs the given action for each entry
func (h *HashMap[K, V]) ForEachEntry(action func(key K, value V)) {
	h.mu.RLock()
	// Create a copy of entries to avoid holding the lock during action execution
	entries := make(map[K]V, len(h.entries))
	for k, v := range h.entries {
		entries[k] = v
	}
	h.mu.RUnlock()

	for k, v := range entries {
		action(k, v)
	}
}

// ComputeIfAbsent computes a value if key is not already associated with a value
func (h *HashMap[K, V]) ComputeIfAbsent(key K, mappingFunction func(K) V) (V, error) {
	if mappingFunction == nil {
		return *new(V), errors.New(string(errcodes.NullPointerError))
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	value, exists := h.entries[key]
	if !exists {
		newValue := mappingFunction(key)
		h.entries[key] = newValue
		return newValue, nil
	}
	return value, nil
}

// PutAllBatch performs a batch put operation for better performance
func (h *HashMap[K, V]) PutAllBatch(entries []collections.MapEntry[K, V]) error {
	if entries == nil {
		return errors.New(string(errcodes.NullPointerError))
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, entry := range entries {
		h.entries[entry.GetKey()] = entry.GetValue()
	}
	return nil
}

// RemoveAllBatch performs a batch remove operation for better performance
func (h *HashMap[K, V]) RemoveAllBatch(keys []K) error {
	if keys == nil {
		return errors.New(string(errcodes.NullPointerError))
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, key := range keys {
		delete(h.entries, key)
	}
	return nil
}

// Clone creates a deep copy of the map
func (h *HashMap[K, V]) Clone() collections.Map[K, V] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	newMap := make(map[K]V, len(h.entries))
	for k, v := range h.entries {
		newMap[k] = v
	}
	return &HashMap[K, V]{
		entries: newMap,
		mu:      sync.RWMutex{},
	}
}

// RemoveIf removes all entries that satisfy the given predicate
func (h *HashMap[K, V]) RemoveIf(predicate func(K, V) bool) bool {
	if predicate == nil {
		return false
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	modified := false
	for k, v := range h.entries {
		if predicate(k, v) {
			delete(h.entries, k)
			modified = true
		}
	}
	return modified
}

// ReplaceAll replaces each entry's value with the result of applying the operator
func (h *HashMap[K, V]) ReplaceAll(operator func(K, V) V) {
	if operator == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for k, v := range h.entries {
		h.entries[k] = operator(k, v)
	}
}

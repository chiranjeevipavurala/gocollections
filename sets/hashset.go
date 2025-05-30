package sets

import (
	"errors"
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
)

type HashSet[E comparable] struct {
	values map[E]bool
	mu     sync.RWMutex
}

func NewHashSet[E comparable]() collections.Set[E] {
	return &HashSet[E]{
		values: make(map[E]bool),
	}
}

func NewHashSetWithCapacity[E comparable](initialCapacity int) collections.Set[E] {
	if initialCapacity < 0 {
		initialCapacity = 0
	}
	return &HashSet[E]{
		values: make(map[E]bool, initialCapacity),
	}
}

func NewHashSetFromCollection[E comparable](collection collections.Collection[E]) collections.Set[E] {
	if collection == nil {
		return NewHashSet[E]()
	}
	set := NewHashSet[E]()
	set.AddAll(collection)
	return set
}

func (h *HashSet[E]) Add(element E) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, exists := h.values[element]; exists {
		return false
	}
	h.values[element] = true
	return true
}

func (h *HashSet[E]) AddAll(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	return h.AddAllBatch(collection.ToArray())
}

func (h *HashSet[E]) AddAllBatch(elements []E) bool {
	if elements == nil {
		return false
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	modified := false
	for _, val := range elements {
		if _, exists := h.values[val]; !exists {
			h.values[val] = true
			modified = true
		}
	}
	return modified
}

func (h *HashSet[E]) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.values = make(map[E]bool)
}

func (h *HashSet[E]) Contains(element E) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	_, exists := h.values[element]
	return exists
}

func (h *HashSet[E]) ContainsAll(collection collections.Collection[E]) (bool, error) {
	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	elements := collection.ToArray()
	if len(elements) == 0 {
		return false, nil
	}
	for _, val := range elements {
		if _, exists := h.values[val]; !exists {
			return false, nil
		}
	}
	return true, nil
}

func (h *HashSet[E]) Equals(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	elements := collection.ToArray()
	if len(h.values) != len(elements) {
		return false
	}
	for _, val := range elements {
		if _, exists := h.values[val]; !exists {
			return false
		}
	}
	return true
}

func (h *HashSet[E]) IsEmpty() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.values) == 0
}

func (h *HashSet[E]) Iterator() collections.Iterator[E] {
	h.mu.RLock()
	defer h.mu.RUnlock()
	values := make([]E, 0, len(h.values))
	for val := range h.values {
		values = append(values, val)
	}
	return &HashSetIterator[E]{
		values: values,
		index:  0,
	}
}

func (h *HashSet[E]) Remove(element E) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, exists := h.values[element]; exists {
		delete(h.values, element)
		return true
	}
	return false
}

func (h *HashSet[E]) RemoveAll(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	return h.RemoveAllBatch(collection.ToArray())
}

func (h *HashSet[E]) RemoveAllBatch(elements []E) bool {
	if elements == nil {
		return false
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	modified := false
	for _, val := range elements {
		if _, exists := h.values[val]; exists {
			delete(h.values, val)
			modified = true
		}
	}
	return modified
}

func (h *HashSet[E]) RetainAll(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	elements := collection.ToArray()
	if len(elements) == 0 {
		if len(h.values) > 0 {
			h.values = make(map[E]bool)
			return true
		}
		return false
	}

	// Create a set of elements to retain
	retainSet := make(map[E]bool, len(elements))
	for _, val := range elements {
		retainSet[val] = true
	}

	// Check if any elements need to be removed
	modified := false
	for val := range h.values {
		if !retainSet[val] {
			delete(h.values, val)
			modified = true
		}
	}
	return modified
}

func (h *HashSet[E]) Size() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.values)
}

func (h *HashSet[E]) ToArray() []E {
	h.mu.RLock()
	defer h.mu.RUnlock()
	result := make([]E, 0, len(h.values))
	for val := range h.values {
		result = append(result, val)
	}
	return result
}

func (h *HashSet[E]) RemoveIf(predicate func(E) bool) bool {
	if predicate == nil {
		return false
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	modified := false
	for val := range h.values {
		if predicate(val) {
			delete(h.values, val)
			modified = true
		}
	}
	return modified
}

func (h *HashSet[E]) ForEach(action func(E)) {
	if action == nil {
		return
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	for val := range h.values {
		action(val)
	}
}

func (h *HashSet[E]) Clone() collections.Set[E] {
	h.mu.RLock()
	defer h.mu.RUnlock()
	newSet := &HashSet[E]{
		values: make(map[E]bool, len(h.values)),
	}
	for val := range h.values {
		newSet.values[val] = true
	}
	return newSet
}

type HashSetIterator[E comparable] struct {
	values []E
	index  int
}

func (h *HashSetIterator[E]) HasNext() bool {
	return h.index < len(h.values)
}

func (h *HashSetIterator[E]) Next() (*E, error) {
	if h.index >= len(h.values) {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	val := h.values[h.index]
	h.index++
	return &val, nil
}

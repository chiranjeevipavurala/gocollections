package sets

import (
	"errors"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/chiranjeevipavurala/gocollections/errorcodes"
)

type HashSet[E comparable] struct {
	values map[E]bool
}

func NewHashSet[E comparable]() collections.Set[E] {
	return &HashSet[E]{
		values: make(map[E]bool),
	}
}

func (h *HashSet[E]) Add(element E) bool {
	if h.Contains(element) {
		return false
	}
	h.values[element] = true
	return true
}

func (h *HashSet[E]) AddAll(collection collections.Collection[E]) bool {
	for _, val := range collection.ToArray() {
		_ = h.Add(val)
	}
	return true
}

func (h *HashSet[E]) Clear() {
	h.values = make(map[E]bool)
}

func (h *HashSet[E]) Contains(element E) bool {
	_, ok := h.values[element]
	return ok
}

func (h *HashSet[E]) ContainsAll(collection collections.Collection[E]) (bool, error) {
	if collection == nil {
		return false, errors.New(string(errorcodes.NullPointerError))
	}
	elements := collection.ToArray()
	if len(elements) == 0 {
		return false, nil
	}
	for _, val := range elements {
		if !h.Contains(val) {
			return false, nil
		}
	}
	return true, nil
}

func (h *HashSet[E]) Equals(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	elements := collection.ToArray()
	if len(h.values) != len(elements) {
		return false
	}
	for _, val := range elements {
		if !h.Contains(val) {
			return false
		}
	}
	return true
}

func (h *HashSet[E]) IsEmpty() bool {
	return len(h.values) == 0
}

func (h *HashSet[E]) Iterator() collections.Iterator[E] {
	return NewHashSetIterator[E](h)
}

func (h *HashSet[E]) Remove(element E) bool {
	if h.Contains(element) {
		delete(h.values, element)
		return true
	}
	return false
}

func (h *HashSet[E]) RemoveAll(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	elements := collection.ToArray()
	if len(elements) == 0 {
		return false
	}
	for _, val := range elements {
		_ = h.Remove(val)
	}
	return true
}

func (h *HashSet[E]) Size() int {
	return len(h.values)
}

func (h *HashSet[E]) ToArray() []E {
	var result []E
	for val := range h.values {
		result = append(result, val)
	}
	return result
}

type HashSetIterator[E comparable] struct {
	values []E
	index  int
}

func NewHashSetIterator[E comparable](h *HashSet[E]) collections.Iterator[E] {
	values := h.ToArray()
	return &HashSetIterator[E]{
		values: values,
		index:  0,
	}
}

func (h *HashSetIterator[E]) HasNext() bool {
	return h.index < len(h.values)
}

func (h *HashSetIterator[E]) Next() (*E, error) {
	if h.index >= len(h.values) {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	val := h.values[h.index]
	h.index++
	return &val, nil
}

package sets

import (
	"sync"
)

type HashSet[T comparable] struct {
	set  map[T]bool
	lock sync.RWMutex
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		set: make(map[T]bool),
	}
}

func (h *HashSet[T]) Add(element T) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.set[element] = true
}

func (h *HashSet[T]) Remove(element T) {
	h.lock.Lock()
	defer h.lock.Unlock()

	delete(h.set, element)
}

func (h *HashSet[T]) Contains(element T) bool {
	h.lock.RLock()
	defer h.lock.RUnlock()

	_, ok := h.set[element]
	return ok
}

func (h *HashSet[T]) Size() int {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return len(h.set)
}

func (h *HashSet[T]) IsEmpty() bool {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return len(h.set) == 0
}

func (h *HashSet[T]) Clear() {
	h.lock.RLock()
	defer h.lock.RUnlock()

	h.set = make(map[T]bool)

}
func (h *HashSet[T]) Iterator() Iterator[T] {
	keys := make([]T, 0)
	for key := range h.set {
		keys = append(keys, key)
	}
	return NewIterator(keys)
}

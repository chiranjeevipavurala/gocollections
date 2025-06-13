package sets

import (
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
)

// node represents an element in the linked list.
type node[E comparable] struct {
	value E
	prev  *node[E]
	next  *node[E]
}

// LinkedHashSet is a Set that maintains insertion order.
type LinkedHashSet[E comparable] struct {
	head  *node[E]
	tail  *node[E]
	items map[E]*node[E]
	mu    sync.RWMutex
}

// NewLinkedHashSet creates a new LinkedHashSet.
func NewLinkedHashSet[E comparable]() *LinkedHashSet[E] {
	return &LinkedHashSet[E]{
		items: make(map[E]*node[E]),
	}
}

// Add adds the specified element to this set if it is not already present.
func (lhs *LinkedHashSet[E]) Add(element E) bool {
	lhs.mu.Lock()
	defer lhs.mu.Unlock()

	if _, exists := lhs.items[element]; exists {
		return false
	}
	newNode := &node[E]{value: element}
	lhs.items[element] = newNode
	if lhs.head == nil {
		lhs.head = newNode
		lhs.tail = newNode
	} else {
		newNode.prev = lhs.tail
		lhs.tail.next = newNode
		lhs.tail = newNode
	}
	return true
}

// AddAll adds all elements from the specified collection to this set.
func (lhs *LinkedHashSet[E]) AddAll(collection collections.Collection[E]) bool {
	lhs.mu.Lock()
	defer lhs.mu.Unlock()

	modified := false
	for _, elem := range collection.ToArray() {
		if _, exists := lhs.items[elem]; !exists {
			newNode := &node[E]{value: elem}
			lhs.items[elem] = newNode
			if lhs.head == nil {
				lhs.head = newNode
				lhs.tail = newNode
			} else {
				newNode.prev = lhs.tail
				lhs.tail.next = newNode
				lhs.tail = newNode
			}
			modified = true
		}
	}
	return modified
}

// Clear removes all elements from this set.
func (lhs *LinkedHashSet[E]) Clear() {
	lhs.mu.Lock()
	defer lhs.mu.Unlock()

	lhs.head = nil
	lhs.tail = nil
	lhs.items = make(map[E]*node[E])
}

// Contains returns true if this set contains the specified element.
func (lhs *LinkedHashSet[E]) Contains(element E) bool {
	lhs.mu.RLock()
	defer lhs.mu.RUnlock()

	_, exists := lhs.items[element]
	return exists
}

// ContainsAll returns true if this set contains all elements from the specified collection.
func (lhs *LinkedHashSet[E]) ContainsAll(collection collections.Collection[E]) (bool, error) {
	lhs.mu.RLock()
	defer lhs.mu.RUnlock()

	for _, elem := range collection.ToArray() {
		if _, exists := lhs.items[elem]; !exists {
			return false, nil
		}
	}
	return true, nil
}

// Equals returns true if this set equals the specified collection.
func (lhs *LinkedHashSet[E]) Equals(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}

	lhs.mu.RLock()
	defer lhs.mu.RUnlock()

	if lhs.Size() != collection.Size() {
		return false
	}

	// Convert collection to array to avoid potential concurrent modifications
	collectionArray := collection.ToArray()
	for _, elem := range collectionArray {
		if _, exists := lhs.items[elem]; !exists {
			return false
		}
	}
	return true
}

// IsEmpty returns true if this set contains no elements.
func (lhs *LinkedHashSet[E]) IsEmpty() bool {
	lhs.mu.RLock()
	defer lhs.mu.RUnlock()

	return len(lhs.items) == 0
}

// Iterator returns an iterator over the elements in this set.
func (lhs *LinkedHashSet[E]) Iterator() collections.Iterator[E] {
	lhs.mu.RLock()
	defer lhs.mu.RUnlock()

	return &linkedHashSetIterator[E]{current: lhs.head}
}

// Remove removes the specified element from this set if it is present.
func (lhs *LinkedHashSet[E]) Remove(element E) bool {
	lhs.mu.Lock()
	defer lhs.mu.Unlock()

	node, exists := lhs.items[element]
	if !exists {
		return false
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lhs.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lhs.tail = node.prev
	}
	delete(lhs.items, element)
	return true
}

// RemoveAll removes all elements from this set that are also contained in the specified collection.
func (lhs *LinkedHashSet[E]) RemoveAll(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}

	lhs.mu.Lock()
	defer lhs.mu.Unlock()

	modified := false
	for _, elem := range collection.ToArray() {
		if node, exists := lhs.items[elem]; exists {
			if node.prev != nil {
				node.prev.next = node.next
			} else {
				lhs.head = node.next
			}
			if node.next != nil {
				node.next.prev = node.prev
			} else {
				lhs.tail = node.prev
			}
			delete(lhs.items, elem)
			modified = true
		}
	}
	return modified
}

// Size returns the number of elements in this set.
func (lhs *LinkedHashSet[E]) Size() int {
	lhs.mu.RLock()
	defer lhs.mu.RUnlock()

	return len(lhs.items)
}

// ToArray returns a slice containing all elements in this set.
func (lhs *LinkedHashSet[E]) ToArray() []E {
	lhs.mu.RLock()
	defer lhs.mu.RUnlock()

	result := make([]E, 0, len(lhs.items))
	for current := lhs.head; current != nil; current = current.next {
		result = append(result, current.value)
	}
	return result
}

// linkedHashSetIterator is an iterator for LinkedHashSet.
type linkedHashSetIterator[E comparable] struct {
	current *node[E]
}

// HasNext returns true if the iteration has more elements.
func (it *linkedHashSetIterator[E]) HasNext() bool {
	return it.current != nil
}

// Next returns the next element in the iteration.
func (it *linkedHashSetIterator[E]) Next() (*E, error) {
	if it.current == nil {
		return nil, collections.ErrNoSuchElement
	}
	value := it.current.value
	it.current = it.current.next
	return &value, nil
}

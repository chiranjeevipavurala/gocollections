package lists

import (
	"errors"
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
)

// Stack represents a LIFO (Last-In-First-Out) stack of elements.
// It is implemented using an ArrayList and provides thread-safe operations.
type Stack[E comparable] struct {
	list  ArrayList[E]
	mutex sync.RWMutex
}

// NewStack creates and returns a new empty Stack.
func NewStack[E comparable]() *Stack[E] {
	return &Stack[E]{
		list: *NewArrayList[E](),
	}
}

// Push adds an element to the top of the stack.
// Returns true if the element was successfully added.
func (s *Stack[E]) Push(element E) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.list.Add(element)
}

// Pop removes and returns the element at the top of the stack.
// Returns an error if the stack is empty.
func (s *Stack[E]) Pop() (*E, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.list.Size() == 0 {
		return nil, errors.New(string(errcodes.EmptyStackError))
	}

	val, _ := s.list.RemoveAtIndex(s.list.Size() - 1)
	return val, nil
}

// Peek returns the element at the top of the stack without removing it.
// Returns an error if the stack is empty.
func (s *Stack[E]) Peek() (*E, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.list.Size() == 0 {
		return nil, errors.New(string(errcodes.EmptyStackError))
	}

	// Get cannot fail here because:
	// 1. We've checked the stack is not empty
	// 2. We're using a mutex for thread safety
	// 3. We're always getting the last element (Size() - 1)
	val, _ := s.list.Get(s.list.Size() - 1)
	return val, nil
}

// IsEmpty returns true if the stack contains no elements.
func (s *Stack[E]) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.list.Size() == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[E]) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.list.Size()
}

// Search returns the 1-based position where an element is on the stack.
// The top element is at position 1, the next element is at position 2, and so on.
// Returns -1 if the element is not found.
func (s *Stack[E]) Search(val E) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	index := s.list.LastIndexOf(val)
	if index == -1 {
		return -1
	}
	return s.list.Size() - index
}

// Clear removes all elements from the stack.
func (s *Stack[E]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.list.Clear()
}

// Contains returns true if the stack contains the specified element.
func (s *Stack[E]) Contains(element E) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.list.Contains(element)
}

// ToArray returns a slice containing all elements in the stack in LIFO order.
func (s *Stack[E]) ToArray() []E {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.list.ToArray()
}

// Clone creates and returns a copy of the stack.
func (s *Stack[E]) Clone() *Stack[E] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	newStack := NewStack[E]()
	clonedList := s.list.Clone()
	newStack.list = *clonedList.(*ArrayList[E])
	return newStack
}

// AddAll adds all elements from the specified collection to the stack.
// Returns true if the stack was modified as a result of the call.
func (s *Stack[E]) AddAll(collection collections.Collection[E]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if collection == nil {
		return false
	}
	return s.list.AddAll(collection)
}

// RemoveAll removes all elements from the stack that are also contained in the specified collection.
// Returns true if the stack was modified as a result of the call.
func (s *Stack[E]) RemoveAll(collection collections.Collection[E]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.list.RemoveAll(collection)
}

// RetainAll retains only the elements in the stack that are contained in the specified collection.
// Returns true if the stack was modified as a result of the call.
func (s *Stack[E]) RetainAll(collection collections.Collection[E]) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	result, _ := s.list.RetainAll(collection)
	return result
}

// Equals compares the stack with the specified collection for equality.
// Returns true if the stack and the collection contain the same elements in the same order.
func (s *Stack[E]) Equals(collection collections.Collection[E]) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.list.Equals(collection)
}

// ForEach performs the given action for each element of the stack.
func (s *Stack[E]) ForEach(consumer func(E)) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.list.ForEach(consumer)
}

// Filter returns a new stack containing only the elements that match the given predicate.
func (s *Stack[E]) Filter(predicate func(E) bool) *Stack[E] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	newStack := NewStack[E]()
	s.list.ForEach(func(element E) {
		if predicate(element) {
			newStack.Push(element)
		}
	})
	return newStack
}

// FindFirst returns the first element that matches the given predicate.
// Returns nil if no element matches or if the stack is empty.
func (s *Stack[E]) FindFirst(predicate func(E) bool) *E {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	element, _ := s.list.FindFirst(predicate)
	return element
}

// FindAll returns a slice containing all elements that match the given predicate.
func (s *Stack[E]) FindAll(predicate func(E) bool) []E {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.list.FindAll(predicate)
}

// RemoveIf removes all elements that match the given predicate.
// Returns true if any elements were removed.
func (s *Stack[E]) RemoveIf(predicate func(E) bool) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.list.RemoveIf(predicate)
}

// ReplaceAll replaces each element with the result of applying the given operator.
func (s *Stack[E]) ReplaceAll(operator func(E) E) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.list.ReplaceAll(operator)
}

// Iterator returns an iterator over the elements in this collection.
func (s *Stack[E]) Iterator() collections.Iterator[E] {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.list.Iterator()
}

// Add adds the specified element to this collection.
// Returns true if the element was successfully added.
func (s *Stack[E]) Add(element E) bool {
	return s.Push(element)
}

// Remove removes a single instance of the specified element from this collection.
// Returns true if the element was removed.
func (s *Stack[E]) Remove(element E) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Get the array and find the index
	values := s.list.ToArray()
	index := -1
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] == element {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	// RemoveAtIndex cannot fail here because:
	// 1. We've found a valid index
	// 2. We're using a mutex for thread safety
	// 3. The index is guaranteed to be within bounds
	_, _ = s.list.RemoveAtIndex(index)
	return true
}

// ContainsAll returns true if this collection contains all elements from the specified collection.
func (s *Stack[E]) ContainsAll(collection collections.Collection[E]) (bool, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}

	elements := collection.ToArray()
	if len(elements) == 0 {
		return true, nil
	}

	for _, element := range elements {
		if !s.list.Contains(element) {
			return false, nil
		}
	}
	return true, nil
}

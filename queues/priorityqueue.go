package queues

import (
	"cmp"
	"errors"
	"reflect"
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
)

// DefaultCapacity is the default initial capacity for PriorityQueue
const DefaultCapacity = 11

// GrowthFactor is the factor by which the capacity is increased when needed
const GrowthFactor = 1.5

// PriorityQueue is an unbounded priority queue based on a priority heap
type PriorityQueue[E comparable] struct {
	elements   []E
	comparator collections.Comparator[E]
	mu         sync.RWMutex // For thread safety
}

// NewPriorityQueue creates a new priority queue with the given comparator
func NewPriorityQueue[E comparable](comparator collections.Comparator[E]) collections.Queue[E] {
	if comparator == nil {
		return nil
	}
	return &PriorityQueue[E]{
		elements:   make([]E, 0, DefaultCapacity),
		comparator: comparator,
		mu:         sync.RWMutex{},
	}
}

// NewPriorityQueueWithCapacity creates a new priority queue with the given comparator and initial capacity
func NewPriorityQueueWithCapacity[E comparable](initialCapacity int, comparator collections.Comparator[E]) collections.Queue[E] {
	if comparator == nil {
		return nil
	}
	if initialCapacity < 1 {
		initialCapacity = DefaultCapacity
	}

	return &PriorityQueue[E]{
		elements:   make([]E, 0, initialCapacity),
		comparator: comparator,
		mu:         sync.RWMutex{},
	}
}

// Add inserts the specified element into this queue
func (pq *PriorityQueue[E]) Add(element E) bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.ensureCapacity(len(pq.elements) + 1)
	pq.elements = append(pq.elements, element)
	pq.siftUp(len(pq.elements) - 1)
	return true
}

// Poll retrieves and removes the head of this queue, or returns nil if this queue is empty
func (pq *PriorityQueue[E]) Poll() (*E, error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.elements) == 0 {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}

	result := pq.elements[0]
	lastIndex := len(pq.elements) - 1

	if lastIndex > 0 {
		// Move the last element to the root and restore heap property
		pq.elements[0] = pq.elements[lastIndex]
		pq.elements = pq.elements[:lastIndex]
		pq.siftDown(0)
	} else {
		pq.elements = pq.elements[:0]
	}

	return &result, nil
}

// Peek retrieves, but does not remove, the head of this queue, or returns nil if this queue is empty
func (pq *PriorityQueue[E]) Peek() (*E, error) {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	if len(pq.elements) == 0 {
		return nil, nil
	}

	return &pq.elements[0], nil
}

// Size returns the number of elements in this collection
func (pq *PriorityQueue[E]) Size() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return len(pq.elements)
}

// IsEmpty returns true if this collection contains no elements
func (pq *PriorityQueue[E]) IsEmpty() bool {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return len(pq.elements) == 0
}

// Clear removes all of the elements from this priority queue
func (pq *PriorityQueue[E]) Clear() {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.elements = make([]E, 0, DefaultCapacity)
}

// Contains returns true if this collection contains the specified element
func (pq *PriorityQueue[E]) Contains(element E) bool {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	for i := 0; i < len(pq.elements); i++ {
		if pq.elements[i] == element {
			return true
		}
	}
	return false
}

// Iterator returns an iterator over the elements in this collection
func (pq *PriorityQueue[E]) Iterator() collections.Iterator[E] {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	// Make a copy of the elements to avoid concurrent modification issues
	elements := make([]E, len(pq.elements))
	copy(elements, pq.elements)

	return &priorityQueueIterator[E]{
		elements: elements,
		position: 0,
	}
}

// ToArray returns an array containing all of the elements in this collection
func (pq *PriorityQueue[E]) ToArray() []E {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	result := make([]E, len(pq.elements))
	copy(result, pq.elements)
	return result
}

// Remove removes the specified element from this queue if it is present
func (pq *PriorityQueue[E]) Remove(element E) bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	index := -1
	for i := 0; i < len(pq.elements); i++ {
		if pq.elements[i] == element {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	// Move the last element to the removed position
	lastIndex := len(pq.elements) - 1
	if index != lastIndex {
		pq.elements[index] = pq.elements[lastIndex]
		// Restore heap property
		pq.siftUp(index)
		pq.siftDown(index)
	}
	// Trim the slice to remove unused capacity
	pq.elements = pq.elements[:lastIndex]

	return true
}

// AddAll adds all of the elements in the specified collection to this collection
func (pq *PriorityQueue[E]) AddAll(collection collections.Collection[E]) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	// Get elements first to minimize lock time
	elements := collection.ToArray()

	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.ensureCapacity(len(pq.elements) + len(elements))

	// Add all elements first
	for _, element := range elements {
		pq.elements = append(pq.elements, element)
	}

	// Heapify the entire array at once (more efficient than sifting up each element)
	pq.heapify()

	return true
}

// RemoveAll removes all of this collection's elements that are also contained in the specified collection
func (pq *PriorityQueue[E]) RemoveAll(collection collections.Collection[E]) bool {
	if collection == nil || collection.IsEmpty() {
		return false
	}

	// Get elements first to minimize lock time
	elements := collection.ToArray()

	pq.mu.Lock()
	defer pq.mu.Unlock()

	modified := false
	for _, element := range elements {
		if pq.removeElement(element) {
			modified = true
		}
	}

	return modified
}

// ContainsAll returns true if this collection contains all of the elements in the specified collection
func (pq *PriorityQueue[E]) ContainsAll(collection collections.Collection[E]) (bool, error) {
	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}

	// Get elements first to minimize lock time
	elements := collection.ToArray()

	pq.mu.RLock()
	defer pq.mu.RUnlock()

	for _, element := range elements {
		if !pq.containsElement(element) {
			return false, nil
		}
	}

	return true, nil
}

// Equals compares the specified object with this collection for equality
func (pq *PriorityQueue[E]) Equals(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}

	// Get elements first to minimize lock time
	elements := collection.ToArray()

	pq.mu.RLock()
	defer pq.mu.RUnlock()

	if len(pq.elements) != len(elements) {
		return false
	}

	// A PriorityQueue is equal to another collection if they contain the same elements
	// regardless of order
	for _, element := range elements {
		if !pq.containsElement(element) {
			return false
		}
	}

	return true
}

// Helper methods for maintaining the heap property

// parent returns the index of the parent of the node at the given index
func (pq *PriorityQueue[E]) parent(index int) int {
	return (index - 1) / 2
}

// leftChild returns the index of the left child of the node at the given index
func (pq *PriorityQueue[E]) leftChild(index int) int {
	return 2*index + 1
}

// rightChild returns the index of the right child of the node at the given index
func (pq *PriorityQueue[E]) rightChild(index int) int {
	return 2*index + 2
}

// siftUp moves the element at the given index up the heap until the heap property is restored
func (pq *PriorityQueue[E]) siftUp(index int) {
	element := pq.elements[index]
	for index > 0 {
		parentIndex := pq.parent(index)
		if pq.comparator.Compare(element, pq.elements[parentIndex]) >= 0 {
			break
		}
		pq.elements[index] = pq.elements[parentIndex]
		index = parentIndex
	}
	pq.elements[index] = element
}

// siftDown moves the element at the given index down the heap until the heap property is restored
func (pq *PriorityQueue[E]) siftDown(index int) {
	element := pq.elements[index]
	halfSize := len(pq.elements) >> 1

	for index < halfSize {
		leftIndex := pq.leftChild(index)
		rightIndex := leftIndex + 1
		smallest := leftIndex

		if rightIndex < len(pq.elements) && pq.comparator.Compare(pq.elements[rightIndex], pq.elements[leftIndex]) < 0 {
			smallest = rightIndex
		}

		if pq.comparator.Compare(element, pq.elements[smallest]) <= 0 {
			break
		}

		pq.elements[index] = pq.elements[smallest]
		index = smallest
	}
	pq.elements[index] = element
}

// heapify converts the array into a heap
func (pq *PriorityQueue[E]) heapify() {
	for i := (len(pq.elements) >> 1) - 1; i >= 0; i-- {
		pq.siftDown(i)
	}
}

// ensureCapacity ensures that the underlying array has at least the specified capacity
func (pq *PriorityQueue[E]) ensureCapacity(minCapacity int) {
	if minCapacity <= cap(pq.elements) {
		return
	}

	// Grow the array to at least double the current capacity
	newCapacity := max(minCapacity, int(float64(cap(pq.elements))*GrowthFactor))
	newElements := make([]E, len(pq.elements), newCapacity)
	copy(newElements, pq.elements)
	pq.elements = newElements
}

// removeElement removes a single instance of the specified element
func (pq *PriorityQueue[E]) removeElement(element E) bool {
	for i := 0; i < len(pq.elements); i++ {
		if pq.elements[i] == element {
			lastIndex := len(pq.elements) - 1
			if i < lastIndex {
				pq.elements[i] = pq.elements[lastIndex]
				if i > 0 && pq.comparator.Compare(pq.elements[i], pq.elements[pq.parent(i)]) < 0 {
					pq.siftUp(i)
				} else {
					pq.siftDown(i)
				}
			}
			pq.elements = pq.elements[:lastIndex]
			return true
		}
	}
	return false
}

// containsElement checks if the element exists in the queue
func (pq *PriorityQueue[E]) containsElement(element E) bool {
	for i := 0; i < len(pq.elements); i++ {
		if pq.elements[i] == element {
			return true
		}
	}
	return false
}

// priorityQueueIterator provides iteration over the elements in a PriorityQueue
type priorityQueueIterator[E comparable] struct {
	elements []E
	position int
}

// HasNext returns true if the iteration has more elements
func (it *priorityQueueIterator[E]) HasNext() bool {
	return it.position < len(it.elements)
}

// Next returns the next element in the iteration
func (it *priorityQueueIterator[E]) Next() (*E, error) {
	if !it.HasNext() {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}

	value := it.elements[it.position]
	it.position++

	return &value, nil
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetComparator returns the comparator used to order the elements in this queue
func (pq *PriorityQueue[E]) GetComparator() collections.Comparator[E] {
	return pq.comparator
}

// RemoveIf removes all elements that satisfy the given predicate
func (pq *PriorityQueue[E]) RemoveIf(predicate func(E) bool) bool {
	if predicate == nil {
		return false
	}

	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.elements) == 0 {
		return false
	}

	// First pass: mark elements to remove
	toRemove := make([]bool, len(pq.elements))
	removed := false
	for i := 0; i < len(pq.elements); i++ {
		if predicate(pq.elements[i]) {
			toRemove[i] = true
			removed = true
		}
	}

	if !removed {
		return false
	}

	// Second pass: remove marked elements and rebuild heap
	newSize := 0
	for i := 0; i < len(pq.elements); i++ {
		if !toRemove[i] {
			pq.elements[newSize] = pq.elements[i]
			newSize++
		}
	}

	// Update size and rebuild heap
	pq.elements = pq.elements[:newSize]
	for i := newSize/2 - 1; i >= 0; i-- {
		pq.siftDown(i)
	}

	return true
}

// ForEach performs the given action for each element
func (pq *PriorityQueue[E]) ForEach(action func(E)) {
	if action == nil {
		return
	}

	pq.mu.RLock()
	defer pq.mu.RUnlock()

	for i := 0; i < len(pq.elements); i++ {
		action(pq.elements[i])
	}
}

// ToArrayWithType returns an array containing all elements in this queue
// with the specified runtime type
func (pq *PriorityQueue[E]) ToArrayWithType(arrayType reflect.Type) (interface{}, error) {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	if arrayType.Kind() != reflect.Array && arrayType.Kind() != reflect.Slice {
		return nil, errors.New("type must be an array or slice")
	}

	// Create a new array of the specified type
	result := reflect.MakeSlice(arrayType, len(pq.elements), len(pq.elements)).Interface()
	resultSlice := reflect.ValueOf(result)

	// Copy elements
	for i := 0; i < len(pq.elements); i++ {
		resultSlice.Index(i).Set(reflect.ValueOf(pq.elements[i]))
	}

	return result, nil
}

// NewPriorityQueueFromCollection creates a new priority queue containing the elements
// of the specified collection
func NewPriorityQueueFromCollection[E comparable](collection collections.Collection[E], comparator collections.Comparator[E]) collections.Queue[E] {
	if comparator == nil {
		return nil
	}
	if collection == nil {
		return NewPriorityQueue[E](comparator)
	}

	elements := collection.ToArray()
	pq := &PriorityQueue[E]{
		elements:   make([]E, len(elements)),
		comparator: comparator,
		mu:         sync.RWMutex{},
	}

	copy(pq.elements, elements)
	pq.heapify() // Convert to heap

	return pq
}

// IntComparator implements Comparator for any ordered type
type IntComparator[E cmp.Ordered] struct{}

func (c *IntComparator[E]) Compare(a, b E) int {
	return cmp.Compare(a, b)
}

// NewPriorityQueueFromSortedSet creates a new priority queue from a sorted set
func NewPriorityQueueFromSortedSet[E cmp.Ordered](sortedSet collections.SortedSet[E]) collections.Queue[E] {
	if sortedSet == nil || sortedSet.IsEmpty() {
		return NewPriorityQueue[E](&IntComparator[E]{})
	}

	// Get elements from the sorted set
	elements := sortedSet.ToArray()
	priorityQueue := NewPriorityQueueWithCapacity[E](len(elements), &IntComparator[E]{})

	// Add all elements to the priority queue
	for _, element := range elements {
		priorityQueue.Add(element)
	}

	return priorityQueue
}

// Element retrieves, but does not remove, the head of this queue
func (pq *PriorityQueue[E]) Element() (*E, error) {
	result, _ := pq.Peek()
	if result == nil {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	return result, nil
}

// Offer inserts the specified element into this priority queue
func (pq *PriorityQueue[E]) Offer(element E) bool {
	return pq.Add(element)
}

// RemoveHead retrieves and removes the head of this queue
func (pq *PriorityQueue[E]) RemoveHead() (*E, error) {
	// Poll cannot fail here because:
	// 1. We're using a mutex for thread safety
	// 2. The empty check is handled by Poll
	return pq.Poll()
}

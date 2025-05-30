package lists

import (
	"errors"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/chiranjeevipavurala/gocollections/collections"
	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
)

// DefaultCapacity is the default initial capacity for ArrayList
const DefaultCapacity = 10

// GrowthFactor is the factor by which the capacity is increased when needed
const GrowthFactor = 1.5

// MaxArraySize is the maximum size of array to allocate
const MaxArraySize = 1<<31 - 1

// MinCapacity is the minimum capacity for ArrayList
const MinCapacity = 4

type ArrayList[E comparable] struct {
	values []E
	mu     sync.RWMutex // For thread safety
}

// calculateNewCapacity calculates the new capacity based on the current capacity and required size
func calculateNewCapacity(currentCapacity, requiredSize int) int {
	if requiredSize > MaxArraySize {
		return MaxArraySize
	}

	if currentCapacity <= 0 {
		return MinCapacity
	}

	newCapacity := int(float64(currentCapacity) * GrowthFactor)
	if newCapacity < requiredSize {
		newCapacity = requiredSize
	}

	if newCapacity > MaxArraySize {
		newCapacity = MaxArraySize
	}

	return newCapacity
}

// ensureCapacity ensures that the list can hold at least the specified number of elements
func (a *ArrayList[E]) ensureCapacity(minCapacity int) {
	if minCapacity > cap(a.values) {
		newCapacity := calculateNewCapacity(cap(a.values), minCapacity)
		newValues := make([]E, len(a.values), newCapacity)
		copy(newValues, a.values)
		a.values = newValues
	}
}

func NewArrayList[E comparable]() *ArrayList[E] {
	values := make([]E, 0, DefaultCapacity)
	return &ArrayList[E]{
		values: values,
		mu:     sync.RWMutex{},
	}
}

func NewArrayListWithInitialCapacity[E comparable](capacity int) *ArrayList[E] {
	if capacity <= 0 || capacity < MinCapacity {
		capacity = DefaultCapacity
	}
	if capacity > MaxArraySize {
		capacity = MaxArraySize
	}
	values := make([]E, 0, capacity)
	return &ArrayList[E]{
		values: values,
		mu:     sync.RWMutex{},
	}
}

func NewArrayListWithInitialCollection[E comparable](values []E) *ArrayList[E] {
	if values == nil {
		values = make([]E, 0, DefaultCapacity)
	}
	return &ArrayList[E]{
		values: values,
		mu:     sync.RWMutex{},
	}
}

func (a *ArrayList[E]) AddAtIndex(index int, element E) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if index > len(a.values) || index < 0 {
		return errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	// Ensure capacity for the new element
	a.ensureCapacity(len(a.values) + 1)

	// Shift elements to make space for the new element
	a.values = append(a.values[:index+1], a.values[index:]...)
	a.values[index] = element

	return nil
}

func (a *ArrayList[E]) Add(element E) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.ensureCapacity(len(a.values) + 1)
	a.values = append(a.values, element)
	return true
}

func (a *ArrayList[E]) AddAllAtIndex(index int, elements collections.Collection[E]) (bool, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if index > len(a.values) || index < 0 {
		return false, errors.New(string(errcodes.IndexOutOfBoundsError))
	}
	if elements == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}

	elementsArray := elements.ToArray()
	if len(elementsArray) == 0 {
		return false, nil
	}

	// Ensure capacity for all new elements
	a.ensureCapacity(len(a.values) + len(elementsArray))

	// Shift elements to make space for new elements
	a.values = append(a.values[:index], append(elementsArray, a.values[index:]...)...)

	return true, nil
}

func (a *ArrayList[E]) AddAll(collection collections.Collection[E]) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if collection == nil {
		return false
	}

	elements := collection.ToArray()
	if len(elements) == 0 {
		return false
	}

	// Ensure capacity for all new elements
	a.ensureCapacity(len(a.values) + len(elements))
	a.values = append(a.values, elements...)
	return true
}

// AddFirst adds the specified element at the beginning of the list.
// This method is thread-safe.
// Time complexity: O(n) due to shifting elements
func (a *ArrayList[E]) AddFirst(element E) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Ensure capacity for the new element
	a.ensureCapacity(len(a.values) + 1)

	// Shift elements to make space for the new element
	a.values = append([]E{element}, a.values...)
}

// AddLast adds the specified element at the end of the list.
// This method is thread-safe.
// Time complexity: O(1) amortized
func (a *ArrayList[E]) AddLast(element E) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Ensure capacity for the new element
	a.ensureCapacity(len(a.values) + 1)
	a.values = append(a.values, element)
}

func (a *ArrayList[E]) Clear() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.values = make([]E, 0, DefaultCapacity)
}

func (a *ArrayList[E]) Contains(element E) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for _, val := range a.values {
		if val == element {
			return true
		}
	}
	return false
}

func (a *ArrayList[E]) ContainsAll(collection collections.Collection[E]) (bool, error) {
	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}
	elements := collection.ToArray()
	if len(elements) == 0 {
		return true, nil
	}
	for _, val := range elements {
		if !a.Contains(val) {
			return false, nil
		}
	}
	return true, nil
}

func (a *ArrayList[E]) CopyOf(collection collections.Collection[E]) collections.List[E] {
	a.mu.RLock()
	defer a.mu.RUnlock()

	values := make([]E, 0)
	if collection == nil {
		return nil
	}
	elements := collection.ToArray()
	if len(elements) == 0 {
		return &ArrayList[E]{
			values: values,
		}
	}
	values = make([]E, len(elements))
	copy(values, elements)
	return &ArrayList[E]{
		values: values,
	}
}

func (a *ArrayList[E]) Equals(collection collections.Collection[E]) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if collection == nil {
		return false
	}
	elements := collection.ToArray()
	if len(a.values) != len(elements) {
		return false
	}
	for i, val := range elements {
		if a.values[i] != val {
			return false
		}
	}
	return true
}

func (a *ArrayList[E]) Get(index int) (*E, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if index >= len(a.values) || index < 0 {
		return nil, errors.New(string(errcodes.IndexOutOfBoundsError))
	}
	return &a.values[index], nil
}

func (a *ArrayList[E]) GetFirst() (*E, error) {
	if len(a.values) == 0 {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	return &a.values[0], nil
}

func (a *ArrayList[E]) GetLast() (*E, error) {
	if len(a.values) == 0 {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	return &a.values[len(a.values)-1], nil
}

// IndexOf returns the index of the first occurrence of the specified element,
// or -1 if the element is not found.
// This method is thread-safe.
// Time complexity: O(n)
func (a *ArrayList[E]) IndexOf(element E) int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for i, val := range a.values {
		if val == element {
			return i
		}
	}
	return -1
}

func (a *ArrayList[E]) IsEmpty() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return len(a.values) == 0
}

func (a *ArrayList[E]) Iterator() collections.Iterator[E] {
	iterator := NewArrayListIterator(a)
	return iterator
}

// LastIndexOf returns the index of the last occurrence of the specified element,
// or -1 if the element is not found.
// This method is thread-safe.
// Time complexity: O(n)
func (a *ArrayList[E]) LastIndexOf(element E) int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for i := len(a.values) - 1; i >= 0; i-- {
		if a.values[i] == element {
			return i
		}
	}
	return -1
}

func (a *ArrayList[E]) Remove(element E) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	index := -1
	for i, val := range a.values {
		if val == element {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	// Shift elements to fill the gap
	copy(a.values[index:], a.values[index+1:])
	a.values = a.values[:len(a.values)-1]

	return true
}

func (a *ArrayList[E]) RemoveAtIndex(index int) (*E, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if index >= len(a.values) || index < 0 {
		return nil, errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	element := a.values[index]

	// Shift elements to fill the gap
	copy(a.values[index:], a.values[index+1:])
	a.values = a.values[:len(a.values)-1]

	return &element, nil
}

func (a *ArrayList[E]) RemoveAll(collection collections.Collection[E]) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if collection == nil {
		return false
	}

	elements := collection.ToArray()
	if len(elements) == 0 {
		return false
	}

	// Create a map for O(1) lookups
	toRemove := make(map[E]struct{}, len(elements))
	for _, elem := range elements {
		toRemove[elem] = struct{}{}
	}

	// Keep track of write position and original length
	writePos := 0
	originalLength := len(a.values)
	for _, val := range a.values {
		if _, shouldRemove := toRemove[val]; !shouldRemove {
			a.values[writePos] = val
			writePos++
		}
	}

	// Trim the slice
	a.values = a.values[:writePos]
	return writePos < originalLength
}

func (a *ArrayList[E]) RemoveFirst() (*E, error) {
	if len(a.values) == 0 {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	return a.RemoveAtIndex(0)
}

func (a *ArrayList[E]) RemoveLast() (*E, error) {
	if len(a.values) == 0 {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	return a.RemoveAtIndex(len(a.values) - 1)
}

func (a *ArrayList[E]) Set(index int, element E) (*E, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if index >= len(a.values) || index < 0 {
		return nil, errors.New(string(errcodes.IndexOutOfBoundsError))
	}
	oldValue := a.values[index]
	a.values[index] = element
	return &oldValue, nil
}

func (a *ArrayList[E]) Size() int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return len(a.values)
}

func (a *ArrayList[E]) SubList(fromIndex int, toIndex int) (collections.List[E], error) {
	if fromIndex < 0 || toIndex > len(a.values) {
		return nil, errors.New(string(errcodes.IndexOutOfBoundsError))
	}
	if fromIndex > toIndex {
		return nil, errors.New(string(errcodes.IllegalArgumentError))
	}
	values := make([]E, toIndex-fromIndex)
	for i := fromIndex; i < toIndex; i++ {
		values[i-fromIndex] = a.values[i]
	}
	return &ArrayList[E]{
		values: values,
	}, nil
}

func (a *ArrayList[E]) Reversed() collections.Collection[E] {
	a.mu.Lock()
	defer a.mu.Unlock()

	for i, j := 0, len(a.values)-1; i < j; i, j = i+1, j-1 {
		a.values[i], a.values[j] = a.values[j], a.values[i]
	}
	return a
}

func (a *ArrayList[E]) Sort(comparator collections.Comparator[E]) {
	a.mu.Lock()
	defer a.mu.Unlock()

	sort.Slice(a.values, func(i, j int) bool {
		return comparator.Compare(a.values[i], a.values[j]) < 0
	})
}

// ToArray returns a slice containing all elements in the list.
// This method is thread-safe.
func (a *ArrayList[E]) ToArray() []E {
	a.mu.RLock()
	defer a.mu.RUnlock()

	result := make([]E, len(a.values))
	copy(result, a.values)
	return result
}

type ArrayListIterator[E comparable] struct {
	values   []E
	index    int
	mu       sync.RWMutex
	modCount int // For fail-fast iteration
}

func NewArrayListIterator[E comparable](a *ArrayList[E]) collections.Iterator[E] {
	a.mu.RLock()
	defer a.mu.RUnlock()

	values := make([]E, len(a.values))
	copy(values, a.values)
	return &ArrayListIterator[E]{
		values:   values,
		index:    0,
		mu:       sync.RWMutex{},
		modCount: 0,
	}
}

func (a *ArrayListIterator[E]) HasNext() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.index < len(a.values)
}

func (a *ArrayListIterator[E]) Next() (*E, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.index >= len(a.values) {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	val := a.values[a.index]
	a.index++
	return &val, nil
}

// TrimToSize reduces the capacity to the current size
func (a *ArrayList[E]) TrimToSize() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.values) < cap(a.values) {
		newValues := make([]E, len(a.values))
		copy(newValues, a.values)
		a.values = newValues
	}
}

// EnsureCapacity ensures that the list can hold at least the specified number of elements
func (a *ArrayList[E]) EnsureCapacity(minCapacity int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if minCapacity > cap(a.values) {
		newCapacity := calculateNewCapacity(cap(a.values), minCapacity)
		newValues := make([]E, len(a.values), newCapacity)
		copy(newValues, a.values)
		a.values = newValues
	}
}

// Capacity returns the current capacity of the list
func (a *ArrayList[E]) Capacity() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return cap(a.values)
}

// RetainAll keeps only elements that are in the specified collection
func (a *ArrayList[E]) RetainAll(collection collections.Collection[E]) (bool, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}

	elements := collection.ToArray()
	if len(elements) == 0 {
		a.values = a.values[:0]
		return true, nil
	}

	// Create a map for O(1) lookups
	toRetain := make(map[E]struct{}, len(elements))
	for _, elem := range elements {
		toRetain[elem] = struct{}{}
	}

	// Keep track of write position
	writePos := 0
	for _, val := range a.values {
		if _, shouldRetain := toRetain[val]; shouldRetain {
			a.values[writePos] = val
			writePos++
		}
	}

	// Trim the slice
	a.values = a.values[:writePos]
	return writePos > 0, nil
}

// Clone creates a deep copy of the list
func (a *ArrayList[E]) Clone() collections.List[E] {
	a.mu.RLock()
	defer a.mu.RUnlock()

	newValues := make([]E, len(a.values))
	copy(newValues, a.values)
	return &ArrayList[E]{
		values: newValues,
	}
}

// Shuffle randomly permutes the list
func (a *ArrayList[E]) Shuffle() {
	a.mu.Lock()
	defer a.mu.Unlock()

	rand.Seed(time.Now().UnixNano())
	for i := len(a.values) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a.values[i], a.values[j] = a.values[j], a.values[i]
	}
}

// FindFirst finds the first element matching the predicate
func (a *ArrayList[E]) FindFirst(predicate func(E) bool) (*E, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for _, val := range a.values {
		if predicate(val) {
			return &val, nil
		}
	}
	return nil, errors.New(string(errcodes.NoSuchElementError))
}

// FindAll finds all elements matching the predicate
func (a *ArrayList[E]) FindAll(predicate func(E) bool) []E {
	a.mu.RLock()
	defer a.mu.RUnlock()

	result := make([]E, 0)
	for _, val := range a.values {
		if predicate(val) {
			result = append(result, val)
		}
	}
	return result
}

// Filter creates a new list with elements matching the predicate
func (a *ArrayList[E]) Filter(predicate func(E) bool) collections.List[E] {
	a.mu.RLock()
	defer a.mu.RUnlock()

	result := NewArrayList[E]()
	for _, val := range a.values {
		if predicate(val) {
			result.Add(val)
		}
	}
	return result
}

// ForEach performs an action on each element
func (a *ArrayList[E]) ForEach(action func(E)) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for _, val := range a.values {
		action(val)
	}
}

// AddAllFirst adds all elements from the collection at the beginning
func (a *ArrayList[E]) AddAllFirst(collection collections.Collection[E]) (bool, error) {
	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	elements := collection.ToArray()
	if len(elements) == 0 {
		return false, nil
	}

	// Ensure capacity for all new elements
	a.ensureCapacity(len(a.values) + len(elements))

	// Create new slice with elements at the beginning
	newValues := make([]E, len(a.values)+len(elements))
	copy(newValues, elements)
	copy(newValues[len(elements):], a.values)
	a.values = newValues
	return true, nil
}

// AddAllLast adds all elements from the collection at the end
func (a *ArrayList[E]) AddAllLast(collection collections.Collection[E]) (bool, error) {
	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	elements := collection.ToArray()
	if len(elements) == 0 {
		return false, nil
	}

	// Ensure capacity for all new elements
	a.ensureCapacity(len(a.values) + len(elements))
	a.values = append(a.values, elements...)
	return true, nil
}

// RemoveIf removes elements that match the predicate
func (a *ArrayList[E]) RemoveIf(predicate func(E) bool) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	newValues := make([]E, 0, len(a.values))
	modified := false

	for _, val := range a.values {
		if !predicate(val) {
			newValues = append(newValues, val)
		} else {
			modified = true
		}
	}

	a.values = newValues
	return modified
}

// ReplaceAll replaces each element with the result of applying the operator
func (a *ArrayList[E]) ReplaceAll(operator func(E) E) {
	a.mu.Lock()
	defer a.mu.Unlock()

	for i := range a.values {
		a.values[i] = operator(a.values[i])
	}
}

// RemoveRange removes elements in the specified range
func (a *ArrayList[E]) RemoveRange(fromIndex, toIndex int) error {
	if fromIndex < 0 || toIndex > len(a.values) || fromIndex > toIndex {
		return errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	newValues := make([]E, len(a.values)-(toIndex-fromIndex))
	copy(newValues, a.values[:fromIndex])
	copy(newValues[fromIndex:], a.values[toIndex:])
	a.values = newValues
	return nil
}

// FastContains uses a map for O(1) lookups when the collection is large.
// The threshold for using map-based lookup is configurable.
// This method is thread-safe.
// Time complexity: O(1) with map, O(n) with linear search
func (a *ArrayList[E]) FastContains(element E) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	const mapThreshold = 1000 // Configurable threshold for map-based lookup
	if len(a.values) > mapThreshold {
		valueMap := make(map[E]struct{}, len(a.values))
		for _, val := range a.values {
			valueMap[val] = struct{}{}
		}
		_, exists := valueMap[element]
		return exists
	}

	// Use linear search for small collections
	for _, val := range a.values {
		if val == element {
			return true
		}
	}
	return false
}

// FastRemoveAll uses a map for O(n) removal of multiple elements
func (a *ArrayList[E]) FastRemoveAll(elements []E) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(elements) == 0 {
		return false
	}

	// Create a map for O(1) lookups
	toRemove := make(map[E]struct{}, len(elements))
	for _, elem := range elements {
		toRemove[elem] = struct{}{}
	}

	// Keep track of write position and original length
	writePos := 0
	originalLength := len(a.values)
	for _, val := range a.values {
		if _, shouldRemove := toRemove[val]; !shouldRemove {
			a.values[writePos] = val
			writePos++
		}
	}

	// Trim the slice
	a.values = a.values[:writePos]
	return writePos < originalLength
}

// FastRetainAll uses a map for O(n) retention of multiple elements
func (a *ArrayList[E]) FastRetainAll(elements []E) (bool, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if elements == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}
	if len(a.values) == 0 {
		return false, nil
	}

	if len(elements) == 0 {
		a.values = a.values[:0]
		return true, nil
	}

	// Create a map for O(1) lookups
	toRetain := make(map[E]struct{}, len(elements))
	for _, elem := range elements {
		toRetain[elem] = struct{}{}
	}

	// Keep track of write position and original length
	writePos := 0
	originalLength := len(a.values)
	for _, val := range a.values {
		if _, shouldRetain := toRetain[val]; shouldRetain {
			a.values[writePos] = val
			writePos++
		}
	}

	// Trim the slice
	a.values = a.values[:writePos]
	return writePos < originalLength, nil
}

// FastIndexOf uses binary search for sorted lists
func (a *ArrayList[E]) FastIndexOf(element E, isSorted bool, comparator collections.Comparator[E]) int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if isSorted && comparator != nil {
		// Binary search for sorted lists
		left, right := 0, len(a.values)-1
		firstOccurrence := -1
		for left <= right {
			mid := (left + right) / 2
			if a.values[mid] == element {
				firstOccurrence = mid
				right = mid - 1 // Continue searching left for earlier occurrence
			} else if comparator.Compare(a.values[mid], element) < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return firstOccurrence
	}

	// Linear search for unsorted lists
	for i, val := range a.values {
		if val == element {
			return i
		}
	}
	return -1
}

// FastLastIndexOf uses reverse linear search for better performance
func (a *ArrayList[E]) FastLastIndexOf(element E) int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	for i := len(a.values) - 1; i >= 0; i-- {
		if a.values[i] == element {
			return i
		}
	}
	return -1
}

// FastSubList uses slice operations for better performance
func (a *ArrayList[E]) FastSubList(fromIndex, toIndex int) (collections.List[E], error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if fromIndex > toIndex {
		return nil, errors.New(string(errcodes.IllegalArgumentError))
	}

	if fromIndex < 0 || toIndex > len(a.values) {
		return nil, errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	subValues := make([]E, toIndex-fromIndex)
	copy(subValues, a.values[fromIndex:toIndex])
	return &ArrayList[E]{
		values: subValues,
	}, nil
}

func (a *ArrayList[E]) checkIndex(index int) error {
	if index < 0 || index >= len(a.values) {
		return errors.New(string(errcodes.IndexOutOfBoundsError))
	}
	return nil
}

func (a *ArrayList[E]) RemoveAllBatch(elements []E) bool {
	if elements == nil {
		return false
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if len(elements) == 0 {
		return false
	}

	// Create a map for O(1) lookups
	toRemove := make(map[E]struct{}, len(elements))
	for _, elem := range elements {
		toRemove[elem] = struct{}{}
	}

	// Keep track of write position and original length
	writePos := 0
	originalLength := len(a.values)
	for _, val := range a.values {
		if _, shouldRemove := toRemove[val]; !shouldRemove {
			a.values[writePos] = val
			writePos++
		}
	}

	// Trim the slice
	a.values = a.values[:writePos]
	return writePos < originalLength
}

func (a *ArrayList[E]) BinarySearch(element E, comparator collections.Comparator[E]) (int, error) {
	if comparator == nil {
		return -1, errors.New(string(errcodes.NullPointerError))
	}

	a.mu.RLock()
	defer a.mu.RUnlock()

	// Binary search for sorted lists
	left, right := 0, len(a.values)-1
	firstOccurrence := -1
	for left <= right {
		mid := (left + right) / 2
		if a.values[mid] == element {
			firstOccurrence = mid
			right = mid - 1 // Continue searching left for earlier occurrence
		} else if comparator.Compare(a.values[mid], element) < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return firstOccurrence, nil
}

func (a *ArrayList[E]) BinarySearchFromIndex(element E, fromIndex, toIndex int, comparator collections.Comparator[E]) (int, error) {
	if fromIndex < 0 || toIndex > len(a.values) || fromIndex > toIndex {
		return -1, errors.New(string(errcodes.IndexOutOfBoundsError))
	}
	if comparator == nil {
		return -1, errors.New(string(errcodes.NullPointerError))
	}

	// Binary search for sorted lists
	left, right := fromIndex, toIndex-1
	firstOccurrence := -1
	for left <= right {
		mid := (left + right) / 2
		if a.values[mid] == element {
			firstOccurrence = mid
			right = mid - 1 // Continue searching left for earlier occurrence
		} else if comparator.Compare(a.values[mid], element) < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return firstOccurrence, nil
}

// AddAllBatch adds all elements from the given slice to the list.
// This method is thread-safe.
// Time complexity: O(n) where n is the length of the elements slice
func (a *ArrayList[E]) AddAllBatch(elements []E) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if elements == nil || len(elements) == 0 {
		return false
	}

	// Ensure capacity for all new elements
	a.ensureCapacity(len(a.values) + len(elements))
	a.values = append(a.values, elements...)
	return true
}

// Package collections provides interfaces and implementations for various data structures.
package collections

// Iterable represents a collection that can be iterated over.
type Iterable[E any] interface {
	// Iterator returns an iterator over the elements in this collection.
	Iterator() Iterator[E]
}

// Collection represents a group of elements.
type Collection[E any] interface {
	Iterable[E]
	// Add adds the specified element to this collection.
	Add(element E) bool
	// AddAll adds all elements from the specified collection to this collection.
	AddAll(collection Collection[E]) bool
	// Clear removes all elements from this collection.
	Clear()
	// Contains returns true if this collection contains the specified element.
	Contains(element E) bool
	// ContainsAll returns true if this collection contains all elements from the specified collection.
	ContainsAll(collection Collection[E]) (bool, error)
	// Equal returns true if this collection is equal to the specified collection.
	Equals(collection Collection[E]) bool
	// Empty returns true if this collection contains no elements.
	IsEmpty() bool
	// Iterator returns an iterator over the elements in this collection.
	Iterator() Iterator[E]
	// Remove removes a single instance of the specified element from this collection.
	Remove(element E) bool
	// RemoveAll removes all elements from this collection that are also contained in the specified collection.
	RemoveAll(collection Collection[E]) bool
	// Size returns the number of elements in this collection.
	Size() int
	// ToArray returns a slice containing all elements in this collection.
	ToArray() []E
}

// SequencedCollection represents a collection with a defined encounter order.
type SequencedCollection[E any] interface {
	Collection[E]
	// AddFirst adds the specified element at the beginning of this collection.
	AddFirst(val E)
	// AddLast adds the specified element at the end of this collection.
	AddLast(val E)
	// GetFirst returns the first element in this collection.
	GetFirst() (*E, error)
	// GetLast returns the last element in this collection.
	GetLast() (*E, error)
	// RemoveFirst removes and returns the first element from this collection.
	RemoveFirst() (*E, error)
	// RemoveLast removes and returns the last element from this collection.
	RemoveLast() (*E, error)
	// Reversed returns a view of this collection in reverse order.
	Reversed() Collection[E]
}

// List represents an ordered collection of elements.
type List[E any] interface {
	SequencedCollection[E]
	// AddAtIndex inserts the specified element at the specified position in this list.
	AddAtIndex(index int, element E) error
	// AddAllAtIndex inserts all elements from the specified collection at the specified position in this list.
	AddAllAtIndex(index int, elements Collection[E]) (bool, error)
	// CopyOf returns a copy of the specified collection as a List.
	CopyOf(collection Collection[E]) List[E]
	// Get returns the element at the specified position in this list.
	Get(index int) (*E, error)
	// IndexOf returns the index of the first occurrence of the specified element in this list.
	IndexOf(element E) int
	// LastIndexOf returns the index of the last occurrence of the specified element in this list.
	LastIndexOf(element E) int
	// RemoveAtIndex removes the element at the specified position in this list.
	RemoveAtIndex(index int) (*E, error)
	// Set replaces the element at the specified position in this list with the specified element.
	Set(index int, element E) (*E, error)
	// Sort sorts this list according to the order induced by the specified comparator.
	Sort(comparator Comparator[E])
	// SubList returns a view of the portion of this list between the specified fromIndex and toIndex.
	SubList(fromIndex int, toIndex int) (List[E], error)
}

// Queue represents a collection designed for holding elements prior to processing.
type Queue[E any] interface {
	Collection[E]
	// Element retrieves, but does not remove, the head of this queue.
	Element() (*E, error)
	// Offer inserts the specified element into this queue.
	Offer(val E) bool
	// Peek retrieves, but does not remove, the head of this queue.
	Peek() (*E, error)
	// Poll retrieves and removes the head of this queue.
	Poll() (*E, error)
	// RemoveHead retrieves and removes the head of this queue.
	RemoveHead() (*E, error)
}

// BlockingQueue represents a queue that additionally supports operations that wait for the queue to become non-empty when retrieving an element.
type BlockingQueue[E any] interface {
	Queue[E]
	// Put inserts the specified element into this queue, waiting if necessary for space to become available.
	Put(element E) error
	// Take retrieves and removes the head of this queue, waiting if necessary until an element becomes available.
	Take() (*E, error)
	// RemainingCapacity returns the number of additional elements that this queue can ideally accept without blocking.
	RemainingCapacity() int
	// DrainTo removes all available elements from this queue and adds them to the given collection.
	DrainTo(c Collection[E]) int
}

// Deque represents a linear collection that supports element insertion and removal at both ends.
type Deque[E any] interface {
	Queue[E]
	SequencedCollection[E]
	// DescendingIterator returns an iterator over the elements in this deque in reverse sequential order.
	DescendingIterator() Iterator[E]
	// OfferFirst inserts the specified element at the front of this deque.
	OfferFirst(val E) bool
	// OfferLast inserts the specified element at the end of this deque.
	OfferLast(val E) bool
	// PeekFirst retrieves, but does not remove, the first element of this deque.
	PeekFirst() (*E, error)
	// PeekLast retrieves, but does not remove, the last element of this deque.
	PeekLast() (*E, error)
	// PollFirst retrieves and removes the first element of this deque.
	PollFirst() (*E, error)
	// PollLast retrieves and removes the last element of this deque.
	PollLast() (*E, error)
	// Pop removes and returns the first element of this deque.
	Pop() (*E, error)
	// Push pushes an element onto the stack represented by this deque.
	Push(val E)
	// RemoveFirstOccurrence removes the first occurrence of the specified element from this deque.
	RemoveFirstOccurrence(val E) bool
	// RemoveLastOccurrence removes the last occurrence of the specified element from this deque.
	RemoveLastOccurrence(val E) bool
}

// Iterator represents an iterator over a collection.
type Iterator[E any] interface {
	// HasNext returns true if the iteration has more elements.
	HasNext() bool
	// Next returns the next element in the iteration.
	Next() (*E, error)
}

// Comparator represents a function that compares two elements.
type Comparator[E any] interface {
	// Compare compares its two arguments for order.
	Compare(a, b E) int
}

// Set represents a collection that contains no duplicate elements.
type Set[E any] interface {
	Collection[E]
}

// SortedSet represents a Set that maintains its elements in ascending order.
type SortedSet[E any] interface {
	Set[E]
	// Comparator returns the comparator used to order the elements in this set.
	Comparator() Comparator[E]
	// First returns the first (lowest) element currently in this set.
	First() (*E, error)
	// Last returns the last (highest) element currently in this set.
	Last() (*E, error)
	// HeadSet returns a view of the portion of this set whose elements are strictly less than toElement.
	HeadSet(toElement E) (SortedSet[E], error)
	// TailSet returns a view of the portion of this set whose elements are greater than or equal to fromElement.
	TailSet(fromElement E) (SortedSet[E], error)
	// SubSet returns a view of the portion of this set whose elements range from fromElement to toElement.
	SubSet(fromElement E, toElement E) (SortedSet[E], error)
}

// NavigableSet represents a SortedSet extended with navigation methods.
type NavigableSet[E any] interface {
	SortedSet[E]
	// Ceiling returns the least element in this set greater than or equal to the given element.
	Ceiling(e E) (*E, error)
	// Floor returns the greatest element in this set less than or equal to the given element.
	Floor(e E) (*E, error)
	// Higher returns the least element in this set strictly greater than the given element.
	Higher(e E) (*E, error)
	// Lower returns the greatest element in this set strictly less than the given element.
	Lower(e E) (*E, error)
	// PollFirst retrieves and removes the first (lowest) element.
	PollFirst() (*E, error)
	// PollLast retrieves and removes the last (highest) element.
	PollLast() (*E, error)
	// DescendingSet returns a reverse order view of the elements contained in this set.
	DescendingSet() NavigableSet[E]
	// DescendingIterator returns an iterator over the elements in this set in descending order.
	DescendingIterator() Iterator[E]
}

// MapEntry represents a key-value pair in a Map.
type MapEntry[K any, V any] interface {
	// GetKey returns the key corresponding to this entry.
	GetKey() K
	// GetValue returns the value corresponding to this entry.
	GetValue() V
}

// Map represents an object that maps keys to values.
type Map[K any, V any] interface {
	// Clear removes all mappings from this map.
	Clear()
	// HasKey returns true if this map contains a mapping for the specified key.
	HasKey(key K) bool
	// HasValue returns true if this map maps one or more keys to the specified value.
	HasValue(value V) bool
	// EntrySet returns a Set view of the mappings contained in this map.
	EntrySet() Set[MapEntry[K, V]]
	// Equal returns true if this map is equal to the specified object.
	Equals(obj any) bool
	// Get returns the value to which the specified key is mapped.
	Get(key K) *V
	// Empty returns true if this map contains no key-value mappings.
	IsEmpty() bool
	// KeySet returns a Set view of the keys contained in this map.
	KeySet() Set[K]
	// Put associates the specified value with the specified key in this map.
	Put(key K, value V) V
	// PutAll copies all of the mappings from the specified map to this map.
	PutAll(Map[K, V])
	// PutIfAbsent associates the specified value with the specified key in this map if the key is not already associated with a value.
	PutIfAbsent(key K, value V) V
	// Remove removes the mapping for a key from this map if it is present.
	Remove(key K) V
	// RemoveKeyWithValue removes the entry for the specified key only if it is currently mapped to the specified value.
	RemoveKeyWithValue(key K, value V) bool
	// Replace replaces the entry for the specified key only if it is currently mapped to some value.
	Replace(key K, value V) V
	// ReplaceKeyWithValue replaces the entry for the specified key only if currently mapped to the given old value.
	ReplaceKeyWithValue(key K, oldValue V, newValue V) bool
	// Size returns the number of key-value mappings in this map.
	Size() int
	// Values returns a Collection view of the values contained in this map.
	Values() Collection[V]
}

// SortedMap represents a Map that maintains its entries in ascending order.
type SortedMap[K any, V any] interface {
	Map[K, V]
	// Comparator returns the comparator used to order the keys in this map.
	Comparator() Comparator[K]
	// FirstKey returns the first (lowest) key currently in this map.
	FirstKey() (K, error)
	// LastKey returns the last (highest) key currently in this map.
	LastKey() (K, error)
	// HeadMap returns a view of the portion of this map whose keys are strictly less than toKey.
	HeadMap(toKey K) (SortedMap[K, V], error)
	// TailMap returns a view of the portion of this map whose keys are greater than or equal to fromKey.
	TailMap(fromKey K) (SortedMap[K, V], error)
	// SubMap returns a view of the portion of this map whose keys range from fromKey to toKey.
	SubMap(fromKey K, toKey K) (SortedMap[K, V], error)
}

// NavigableMap represents a SortedMap extended with navigation methods.
type NavigableMap[K any, V any] interface {
	SortedMap[K, V]
	// CeilingEntry returns a key-value mapping associated with the least key greater than or equal to the given key.
	CeilingEntry(key K) (MapEntry[K, V], error)
	// CeilingKey returns the least key greater than or equal to the given key.
	CeilingKey(key K) (K, error)
	// DescendingKeySet returns a reverse order NavigableSet view of the keys contained in this map.
	DescendingKeySet() NavigableSet[K]
	// DescendingMap returns a reverse order view of the mappings contained in this map.
	DescendingMap() NavigableMap[K, V]
	// FirstEntry returns a key-value mapping associated with the least key in this map.
	FirstEntry() (MapEntry[K, V], error)
	// FloorEntry returns a key-value mapping associated with the greatest key less than or equal to the given key.
	FloorEntry(key K) (MapEntry[K, V], error)
	// FloorKey returns the greatest key less than or equal to the given key.
	FloorKey(key K) (K, error)
	// HigherEntry returns a key-value mapping associated with the least key strictly greater than the given key.
	HigherEntry(key K) (MapEntry[K, V], error)
	// HigherKey returns the least key strictly greater than the given key.
	HigherKey(key K) (K, error)
	// LastEntry returns a key-value mapping associated with the greatest key in this map.
	LastEntry() (MapEntry[K, V], error)
	// LowerEntry returns a key-value mapping associated with the greatest key strictly less than the given key.
	LowerEntry(key K) (MapEntry[K, V], error)
	// LowerKey returns the greatest key strictly less than the given key.
	LowerKey(key K) (K, error)
	// NavigableKeySet returns a NavigableSet view of the keys contained in this map.
	NavigableKeySet() NavigableSet[K]
	// PollFirstEntry removes and returns a key-value mapping associated with the least key in this map.
	PollFirstEntry() (MapEntry[K, V], error)
	// PollLastEntry removes and returns a key-value mapping associated with the greatest key in this map.
	PollLastEntry() (MapEntry[K, V], error)
}

// ConcurrentMap represents a Map that supports concurrent access.
type ConcurrentMap[K any, V any] interface {
	Map[K, V]
	// GetOrDefault returns the value to which the specified key is mapped, or defaultValue if this map contains no mapping for the key.
	GetOrDefault(key K, defaultValue V) V
	// ForEachEntry performs the given action for each entry in this map.
	ForEachEntry(action func(key K, value V))
	// ComputeIfAbsent computes a value for the specified key if the key is not already associated with a value.
	ComputeIfAbsent(key K, mappingFunction func(K) V) (V, error)
}

// HashMapEntry represents a key-value pair in a HashMap.
type HashMapEntry[K comparable, V comparable] struct {
	Key   K
	Value V
}

// NewHashMapEntry creates a new HashMapEntry with the specified key and value.
func NewHashMapEntry[K comparable, V comparable](key K, value V) *HashMapEntry[K, V] {
	return &HashMapEntry[K, V]{
		Key:   key,
		Value: value,
	}
}

// GetKey returns the key corresponding to this entry.
func (h *HashMapEntry[K, V]) GetKey() K {
	return h.Key
}

// GetValue returns the value corresponding to this entry.
func (h *HashMapEntry[K, V]) GetValue() V {
	return h.Value
}

// Equal returns true if this entry is equal to the specified object.
func (h *HashMapEntry[K, V]) Equals(obj any) bool {
	if obj == nil {
		return false
	}
	entry, ok := obj.(*HashMapEntry[K, V])
	if !ok {
		return false
	}
	return h.Key == entry.Key && h.Value == entry.Value
}

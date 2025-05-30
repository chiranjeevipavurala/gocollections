package lists

import (
	"errors"
	"sort"
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
)

type LinkedList[E comparable] struct {
	head ListNode[E]
	tail ListNode[E]
	size int
	mu   sync.RWMutex
}

func NewLinkedList[E comparable]() *LinkedList[E] {
	return &LinkedList[E]{
		head: nil,
		tail: nil,
		size: 0,
		mu:   sync.RWMutex{},
	}
}

func NewLinkedListWithInitialCollection[E comparable](values []E) *LinkedList[E] {
	list := NewLinkedList[E]()
	if len(values) > 0 {
		list.AddAllBatch(values)
	}
	return list
}

func (l *LinkedList[E]) Add(element E) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := NewListNodeImpl(element)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.SetNext(newNode)
		newNode.SetPrev(l.tail)
		l.tail = newNode
	}
	l.size++
	return true
}

func (l *LinkedList[E]) AddAtIndex(index int, element E) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index > l.size || index < 0 {
		return errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	newNode := NewListNodeImpl(element)
	if index == 0 {
		newNode.SetNext(l.head)
		if l.head != nil {
			l.head.SetPrev(newNode)
		}
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
	} else if index == l.size {
		l.tail.SetNext(newNode)
		newNode.SetPrev(l.tail)
		l.tail = newNode
	} else {
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.GetNext()
		}
		newNode.SetNext(current.GetNext())
		newNode.SetPrev(current)
		current.GetNext().SetPrev(newNode)
		current.SetNext(newNode)
	}
	l.size++
	return nil
}

func (l *LinkedList[E]) AddFirst(element E) {
	// Adding at index 0 is always safe as it doesn't depend on the current size
	_ = l.AddAtIndex(0, element)
}

func (l *LinkedList[E]) AddLast(element E) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := NewListNodeImpl(element)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.SetNext(newNode)
		newNode.SetPrev(l.tail)
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList[E]) AddAll(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	return l.AddAllBatch(collection.ToArray())
}

func (l *LinkedList[E]) AddAllBatch(elements []E) bool {
	if len(elements) == 0 {
		return false
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	for _, element := range elements {
		newNode := NewListNodeImpl(element)
		if l.head == nil {
			l.head = newNode
			l.tail = newNode
		} else {
			l.tail.SetNext(newNode)
			newNode.SetPrev(l.tail)
			l.tail = newNode
		}
		l.size++
	}
	return true
}

func (l *LinkedList[E]) AddAllAtIndex(index int, elements collections.Collection[E]) (bool, error) {
	if elements == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}
	return l.AddAllAtIndexBatch(index, elements.ToArray())
}

func (l *LinkedList[E]) AddAllAtIndexBatch(index int, elements []E) (bool, error) {
	if len(elements) == 0 {
		return false, nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if index > l.size || index < 0 {
		return false, errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	// Create a sublist of new nodes
	var firstNewNode, lastNewNode ListNode[E]
	for _, element := range elements {
		newNode := NewListNodeImpl(element)
		if firstNewNode == nil {
			firstNewNode = newNode
			lastNewNode = newNode
		} else {
			lastNewNode.SetNext(newNode)
			newNode.SetPrev(lastNewNode)
			lastNewNode = newNode
		}
	}

	if index == 0 {
		lastNewNode.SetNext(l.head)
		if l.head != nil {
			l.head.SetPrev(lastNewNode)
		}
		l.head = firstNewNode
		if l.tail == nil {
			l.tail = lastNewNode
		}
	} else if index == l.size {
		l.tail.SetNext(firstNewNode)
		firstNewNode.SetPrev(l.tail)
		l.tail = lastNewNode
	} else {
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.GetNext()
		}
		lastNewNode.SetNext(current.GetNext())
		current.GetNext().SetPrev(lastNewNode)
		current.SetNext(firstNewNode)
		firstNewNode.SetPrev(current)
	}

	l.size += len(elements)
	return true, nil
}

func (l *LinkedList[E]) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[E]) Contains(element E) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	current := l.head
	for current != nil {
		if *current.GetData() == element {
			return true
		}
		current = current.GetNext()
	}
	return false
}

func (l *LinkedList[E]) ContainsAll(collection collections.Collection[E]) (bool, error) {
	if collection == nil {
		return false, errors.New(string(errcodes.NullPointerError))
	}
	if collection.Size() == 0 {
		return true, nil
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	// Note: contains() assumes the lock is already held
	elements := collection.ToArray()
	for _, val := range elements {
		if !l.contains(val) {
			return false, nil
		}
	}
	return true, nil
}

// contains is an internal method that assumes the lock is already held
// It should only be called from methods that have already acquired the appropriate lock
func (l *LinkedList[E]) contains(element E) bool {
	current := l.head
	for current != nil {
		if *current.GetData() == element {
			return true
		}
		current = current.GetNext()
	}
	return false
}

func (l *LinkedList[E]) Equals(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	elements := collection.ToArray()
	if l.size != len(elements) {
		return false
	}

	current := l.head
	for _, val := range elements {
		if *current.GetData() != val {
			return false
		}
		current = current.GetNext()
	}
	return true
}

func (l *LinkedList[E]) Get(index int) (*E, error) {
	if err := l.checkIndex(index); err != nil {
		return nil, err
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	// Fast path for first and last elements
	var val E
	if index == 0 {
		val = *l.head.GetData()
		return &val, nil
	}
	if index == l.size-1 {
		val = *l.tail.GetData()
		return &val, nil
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.GetNext()
	}
	val = *current.GetData()
	return &val, nil
}

func (l *LinkedList[E]) GetFirst() (*E, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.size == 0 {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	val := *l.head.GetData()
	return &val, nil
}

func (l *LinkedList[E]) GetLast() (*E, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.size == 0 {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	val := *l.tail.GetData()
	return &val, nil
}

func (l *LinkedList[E]) IndexOf(element E) int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	current := l.head
	for i := 0; i < l.size; i++ {
		if *current.GetData() == element {
			return i
		}
		current = current.GetNext()
	}
	return -1
}

func (l *LinkedList[E]) IsEmpty() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.size == 0
}

// ListIteratorImpl is a thread-safe iterator implementation
type ListIteratorImpl[E comparable] struct {
	current ListNode[E]
	mu      *sync.RWMutex
}

func NewListIteratorImpl[E comparable](head ListNode[E], mu *sync.RWMutex) collections.Iterator[E] {
	return &ListIteratorImpl[E]{
		current: head,
		mu:      mu,
	}
}

func (iter *ListIteratorImpl[E]) HasNext() bool {
	iter.mu.RLock()
	defer iter.mu.RUnlock()
	return iter.current != nil
}

func (iter *ListIteratorImpl[E]) Next() (*E, error) {
	iter.mu.RLock()
	defer iter.mu.RUnlock()

	if iter.current == nil {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	val := *iter.current.GetData()
	iter.current = iter.current.GetNext()
	return &val, nil
}

// DescendingIteratorImpl is a thread-safe iterator implementation
type DescendingIteratorImpl[E comparable] struct {
	current ListNode[E]
	mu      *sync.RWMutex
}

func NewDescendingIteratorImpl[E comparable](tail ListNode[E], mu *sync.RWMutex) collections.Iterator[E] {
	return &DescendingIteratorImpl[E]{
		current: tail,
		mu:      mu,
	}
}

func (iter *DescendingIteratorImpl[E]) HasNext() bool {
	iter.mu.RLock()
	defer iter.mu.RUnlock()
	return iter.current != nil
}

func (iter *DescendingIteratorImpl[E]) Next() (*E, error) {
	iter.mu.RLock()
	defer iter.mu.RUnlock()

	if iter.current == nil {
		return nil, errors.New(string(errcodes.NoSuchElementError))
	}
	val := *iter.current.GetData()
	iter.current = iter.current.GetPrev()
	return &val, nil
}

func (l *LinkedList[E]) Iterator() collections.Iterator[E] {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return NewListIteratorImpl(l.head, &l.mu)
}

func (l *LinkedList[E]) DescendingIterator() collections.Iterator[E] {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return NewDescendingIteratorImpl(l.tail, &l.mu)
}

func (l *LinkedList[E]) LastIndexOf(element E) int {
	l.mu.RLock()
	defer l.mu.RUnlock()

	current := l.tail
	index := -1
	for i := l.size - 1; i >= 0; i-- {
		if *current.GetData() == element {
			index = i
			break
		}
		current = current.GetPrev()
	}
	return index
}

func (l *LinkedList[E]) RemoveAtIndex(index int) (*E, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index >= l.size || index < 0 {
		return nil, errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	var val E
	if index == 0 {
		val = *l.head.GetData()
		l.head = l.head.GetNext()
		if l.head != nil {
			l.head.SetPrev(nil)
		} else {
			l.tail = nil
		}
	} else if index == l.size-1 {
		val = *l.tail.GetData()
		l.tail = l.tail.GetPrev()
		l.tail.SetNext(nil)
	} else {
		current := l.head
		for i := 0; i < index; i++ {
			current = current.GetNext()
		}
		val = *current.GetData()
		current.GetPrev().SetNext(current.GetNext())
		current.GetNext().SetPrev(current.GetPrev())
	}
	l.size--
	return &val, nil
}

func (l *LinkedList[E]) RemoveFirst() (*E, error) {
	return l.RemoveAtIndex(0)
}

func (l *LinkedList[E]) RemoveLast() (*E, error) {
	// RemoveAtIndex cannot fail here because:
	// 1. We're using a mutex for thread safety
	// 2. We're always removing the last element (size - 1)
	// 3. The size check is handled by RemoveAtIndex
	return l.RemoveAtIndex(l.size - 1)
}

func (l *LinkedList[E]) Remove(element E) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.head
	for current != nil {
		if *current.GetData() == element {
			if current == l.head {
				l.head = current.GetNext()
				if l.head != nil {
					l.head.SetPrev(nil)
				} else {
					l.tail = nil
				}
			} else if current == l.tail {
				l.tail = current.GetPrev()
				l.tail.SetNext(nil)
			} else {
				current.GetPrev().SetNext(current.GetNext())
				current.GetNext().SetPrev(current.GetPrev())
			}
			l.size--
			return true
		}
		current = current.GetNext()
	}
	return false
}

func (l *LinkedList[E]) RemoveAll(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
	return l.RemoveAllBatch(collection.ToArray())
}

func (l *LinkedList[E]) RemoveAllBatch(elements []E) bool {
	if len(elements) == 0 {
		return false
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Create a map for O(1) lookups
	toRemove := make(map[E]struct{}, len(elements))
	for _, elem := range elements {
		toRemove[elem] = struct{}{}
	}

	modified := false
	current := l.head
	for current != nil {
		next := current.GetNext()
		if _, shouldRemove := toRemove[*current.GetData()]; shouldRemove {
			if current == l.head {
				l.head = next
				if l.head != nil {
					l.head.SetPrev(nil)
				} else {
					l.tail = nil
				}
			} else if current == l.tail {
				l.tail = current.GetPrev()
				l.tail.SetNext(nil)
			} else {
				current.GetPrev().SetNext(next)
				next.SetPrev(current.GetPrev())
			}
			l.size--
			modified = true
		}
		current = next
	}
	return modified
}

func (l *LinkedList[E]) Set(index int, element E) (*E, error) {
	if err := l.checkIndex(index); err != nil {
		return nil, err
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var current ListNode[E]
	if index == 0 {
		current = l.head
	} else if index == l.size-1 {
		current = l.tail
	} else {
		current = l.head
		for i := 0; i < index; i++ {
			current = current.GetNext()
		}
	}

	oldVal := *current.GetData() // Create a copy of the old value
	current.SetData(element)
	return &oldVal, nil // Return pointer to the copy
}

func (l *LinkedList[E]) Size() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.size
}

func (l *LinkedList[E]) ToArray() []E {
	l.mu.RLock()
	defer l.mu.RUnlock()

	values := make([]E, l.size)
	current := l.head
	for i := 0; i < l.size; i++ {
		values[i] = *current.GetData()
		current = current.GetNext()
	}
	return values
}

func (l *LinkedList[E]) SubList(fromIndex int, toIndex int) (collections.List[E], error) {
	if fromIndex < 0 || toIndex > l.size || fromIndex > toIndex {
		return nil, errors.New(string(errcodes.IndexOutOfBoundsError))
	}

	values := make([]E, toIndex-fromIndex)
	current := l.head
	for i := 0; i < fromIndex; i++ {
		current = current.GetNext()
	}
	for i := 0; i < toIndex-fromIndex; i++ {
		values[i] = *current.GetData()
		current = current.GetNext()
	}
	return NewLinkedListWithInitialCollection(values), nil
}

func (l *LinkedList[E]) CopyOf(collection collections.Collection[E]) collections.List[E] {
	if collection == nil {
		return NewLinkedList[E]()
	}
	return NewLinkedListWithInitialCollection(collection.ToArray())
}

// Queue operations
func (l *LinkedList[E]) Element() (*E, error) {
	return l.GetFirst()
}

func (l *LinkedList[E]) Offer(val E) bool {
	return l.Add(val)
}

func (l *LinkedList[E]) Peek() (*E, error) {
	if l.size == 0 {
		return nil, nil
	}
	return l.GetFirst()
}

func (l *LinkedList[E]) Poll() (*E, error) {
	if l.size == 0 {
		return nil, nil
	}
	return l.RemoveFirst()
}

// Deque operations
func (l *LinkedList[E]) OfferFirst(val E) bool {
	l.AddFirst(val)
	return true
}

func (l *LinkedList[E]) OfferLast(val E) bool {
	l.AddLast(val)
	return true
}

func (l *LinkedList[E]) PeekFirst() (*E, error) {
	return l.GetFirst()
}

func (l *LinkedList[E]) PeekLast() (*E, error) {
	return l.GetLast()
}

func (l *LinkedList[E]) PollFirst() (*E, error) {
	return l.Poll()
}

func (l *LinkedList[E]) PollLast() (*E, error) {
	if l.size == 0 {
		return nil, nil
	}
	return l.RemoveLast()
}

// Stack operations
func (l *LinkedList[E]) Pop() (*E, error) {
	return l.RemoveFirst()
}

func (l *LinkedList[E]) Push(val E) {
	l.AddFirst(val)
}

// Additional operations
func (l *LinkedList[E]) RemoveFirstOccurrence(val E) bool {
	return l.Remove(val)
}

func (l *LinkedList[E]) RemoveLastOccurrence(val E) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.tail
	for current != nil {
		if *current.GetData() == val {
			if current == l.head {
				l.head = current.GetNext()
				if l.head != nil {
					l.head.SetPrev(nil)
				} else {
					l.tail = nil
				}
			} else if current == l.tail {
				l.tail = current.GetPrev()
				l.tail.SetNext(nil)
			} else {
				current.GetPrev().SetNext(current.GetNext())
				current.GetNext().SetPrev(current.GetPrev())
			}
			l.size--
			return true
		}
		current = current.GetPrev()
	}
	return false
}

func (l *LinkedList[E]) Reversed() collections.Collection[E] {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.head
	var prev ListNode[E] = nil
	for current != nil {
		next := current.GetNext()
		current.SetNext(prev)
		current.SetPrev(next)
		prev = current
		current = next
	}
	l.head, l.tail = l.tail, l.head
	return l
}

func (l *LinkedList[E]) Sort(comparator collections.Comparator[E]) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Get the values
	values := make([]E, l.size)
	current := l.head
	for i := 0; i < l.size; i++ {
		values[i] = *current.GetData()
		current = current.GetNext()
	}

	// Sort the values
	sort.Slice(values, func(i, j int) bool {
		return comparator.Compare(values[i], values[j]) < 0
	})

	// Rebuild the list - directly reset instead of calling Clear()
	l.head = nil
	l.tail = nil
	l.size = 0

	// Add sorted values
	for _, val := range values {
		newNode := NewListNodeImpl(val)
		if l.head == nil {
			l.head = newNode
			l.tail = newNode
		} else {
			l.tail.SetNext(newNode)
			newNode.SetPrev(l.tail)
			l.tail = newNode
		}
		l.size++
	}
}

func (l *LinkedList[E]) checkIndex(index int) error {
	if index < 0 || index >= l.size {
		return errors.New(string(errcodes.IndexOutOfBoundsError))
	}
	return nil
}

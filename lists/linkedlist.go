package lists

import (
	"errors"
	"sort"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/chiranjeevipavurala/gocollections/errorcodes"
)

type LinkedList[E comparable] struct {
	head ListNode[E]
	size int
}

func NewLinkedList[E comparable]() collections.List[E] {
	return &LinkedList[E]{
		head: nil,
		size: 0,
	}
}
func NewLinkedListWithInitialCollection[E comparable](values []E) collections.List[E] {
	list := NewLinkedList[E]()
	for _, val := range values {
		list.Add(val)
	}
	return list
}
func (l *LinkedList[E]) Add(element E) bool {
	newNode := NewListNodeImpl[E](element)
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.GetNext() != nil {
			current = current.GetNext()
		}
		current.SetNext(newNode)
		newNode.SetPrev(current)
	}
	l.size = l.size + 1
	return true
}

func (l *LinkedList[E]) AddAtIndex(index int, element E) error {
	if index > l.size || index < 0 {
		return errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	newNode := NewListNodeImpl[E](element)
	if index == 0 {
		newNode.SetNext(l.head)
		l.head = newNode
	} else {
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.GetNext()
		}
		newNode.SetNext(current.GetNext())
		newNode.SetPrev(current)
		current.SetNext(newNode)
	}
	l.size = l.size + 1
	return nil
}
func (l *LinkedList[E]) AddFirst(element E) {
	_ = l.AddAtIndex(0, element)
}
func (l *LinkedList[E]) AddLast(element E) {
	_ = l.AddAtIndex(l.size, element)
}

func (l *LinkedList[E]) AddAll(collection collections.Collection[E]) bool {
	elements := collection.ToArray()
	for _, val := range elements {
		_ = l.Add(val)
	}
	return true
}

func (l *LinkedList[E]) AddAllAtIndex(index int, elements []E) (bool, error) {
	if index > l.size || index < 0 {
		return false, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	for i, val := range elements {
		_ = l.AddAtIndex(index+i, val)
	}
	return true, nil
}
func (l *LinkedList[E]) Clear() {
	l.head = nil
	l.size = 0
}
func (l *LinkedList[E]) Contains(element E) bool {
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
		return false, errors.New(string(errorcodes.NullPointerError))
	}
	if collection.Size() == 0 {
		return false, nil
	}
	elements := collection.ToArray()
	for _, val := range elements {
		if !l.Contains(val) {
			return false, nil
		}
	}
	return true, nil
}

func (l *LinkedList[E]) Equals(collection collections.Collection[E]) bool {
	if collection == nil {
		return false
	}
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
	if index >= l.size || index < 0 {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	current := l.head
	for i := 0; i < index; i++ {
		current = current.GetNext()
	}
	return current.GetData(), nil
}
func (l *LinkedList[E]) GetFirst() (*E, error) {
	if l.size == 0 {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	return l.head.GetData(), nil
}
func (l *LinkedList[E]) GetLast() (*E, error) {
	if l.size == 0 {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	current := l.head
	for i := 0; i < l.size-1; i++ {
		current = current.GetNext()
	}
	return current.GetData(), nil
}

func (l *LinkedList[E]) IndexOf(element E) int {
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
	return l.size == 0
}
func (l *LinkedList[E]) Iterator() collections.Iterator[E] {
	return NewListIteratorImpl[E](l.head)
}
func (l *LinkedList[E]) DescendingIterator() collections.Iterator[E] {
	return NewDescendingIteratorImpl[E](l.head)
}
func (l *LinkedList[E]) LastIndexOf(element E) int {
	current := l.head
	index := -1
	for i := 0; i < l.size; i++ {
		if *current.GetData() == element {
			index = i
		}
		current = current.GetNext()
	}
	return index
}

func (l *LinkedList[E]) RemoveAtIndex(index int) (*E, error) {
	if index >= l.size || index < 0 {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	var val E
	if index == 0 {
		val = *l.head.GetData()
		l.head = l.head.GetNext()
		if l.head != nil {
			l.head.SetPrev(nil)
		}
	} else {
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.GetNext()
		}
		val = *current.GetNext().GetData()
		current.SetNext(current.GetNext().GetNext())
		if current.GetNext() != nil {
			current.GetNext().SetPrev(current)
		}
	}
	l.size = l.size - 1
	return &val, nil
}
func (l *LinkedList[E]) RemoveFirst() (*E, error) {
	if l.size == 0 {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	val := l.head.GetData()
	l.head = l.head.GetNext()
	if l.head != nil {
		l.head.SetPrev(nil)
	}
	l.size = l.size - 1
	return val, nil
}
func (l *LinkedList[E]) RemoveLast() (*E, error) {
	if l.size == 0 {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	if l.size == 1 {
		val := l.head.GetData()
		l.head = nil
		l.size = l.size - 1
		return val, nil
	}
	current := l.head
	for i := 0; i < l.size-2; i++ {
		current = current.GetNext()
	}
	val := current.GetNext().GetData()
	current.SetNext(nil)
	l.size = l.size - 1
	return val, nil
}
func (l *LinkedList[E]) Remove(element E) bool {
	index := l.IndexOf(element)
	if index == -1 {
		return false
	}
	_, _ = l.RemoveAtIndex(index)
	return true
}

func (l *LinkedList[E]) RemoveAll(collection collections.Collection[E]) bool {
	elements := collection.ToArray()
	for _, val := range elements {
		_ = l.Remove(val)
	}
	return true
}

func (l *LinkedList[E]) Set(index int, element E) (*E, error) {
	if index >= l.size || index < 0 {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	current := l.head
	for i := 0; i < index; i++ {
		current = current.GetNext()
	}
	val := current.GetData()
	current.SetData(element)
	return val, nil
}
func (l *LinkedList[E]) Size() int {
	return l.size
}
func (l *LinkedList[E]) ToArray() []E {
	values := make([]E, l.size)
	current := l.head
	for i := 0; i < l.size; i++ {
		values[i] = *current.GetData()
		current = current.GetNext()
	}
	return values
}
func (l *LinkedList[E]) SubList(fromIndex int, toIndex int) (collections.List[E], error) {
	if fromIndex < 0 || toIndex > l.size {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	if fromIndex > toIndex {
		return nil, errors.New(string(errorcodes.IllegalArgumentError))
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
	return NewLinkedListWithInitialCollection[E](values), nil
}
func (l *LinkedList[E]) RemoveHead() (*E, error) {
	return l.RemoveFirst()
}
func (l *LinkedList[E]) Element() (*E, error) {
	return l.GetFirst()
}
func (l *LinkedList[E]) Offer(val E) bool {
	return l.Add(val)
}
func (l *LinkedList[E]) Peek() (*E, error) {
	return l.GetFirst()
}
func (l *LinkedList[E]) Poll() (*E, error) {
	return l.RemoveFirst()
}
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
	return l.RemoveFirst()
}
func (l *LinkedList[E]) PollLast() (*E, error) {
	return l.RemoveLast()
}
func (l *LinkedList[E]) Pop() (*E, error) {
	return l.RemoveFirst()
}
func (l *LinkedList[E]) Push(val E) {
	l.AddFirst(val)
}
func (l *LinkedList[E]) RemoveFirstOccurrence(val E) bool {
	return l.Remove(val)
}
func (l *LinkedList[E]) RemoveLastOccurrence(val E) bool {
	index := l.LastIndexOf(val)
	if index == -1 {
		return false
	}
	_, _ = l.RemoveAtIndex(index)
	return true
}
func (l *LinkedList[E]) Reversed() {
	current := l.head
	var prev ListNode[E] = nil
	for current != nil {
		next := current.GetNext()
		current.SetNext(prev)
		current.SetPrev(next)
		prev = current
		current = next
	}
	l.head = prev
}

func (l *LinkedList[E]) Sort(comparator collections.Comparator[E]) {
	values := l.ToArray()
	sort.Slice(values, func(i, j int) bool {
		return comparator.Compare(values[i], values[j]) < 0
	})
	l.Clear()
	for _, val := range values {
		_ = l.Add(val)
	}
}

type ListIteratorImpl[E comparable] struct {
	current ListNode[E]
}

func NewListIteratorImpl[E comparable](head ListNode[E]) collections.Iterator[E] {
	return &ListIteratorImpl[E]{
		current: head,
	}
}
func (iter *ListIteratorImpl[E]) HasNext() bool {
	return iter.current != nil
}
func (iter *ListIteratorImpl[E]) Next() (*E, error) {
	if iter.current == nil {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	val := iter.current.GetData()
	iter.current = iter.current.GetNext()
	return val, nil
}

type DescendingIteratorImpl[E comparable] struct {
	current ListNode[E]
}

func NewDescendingIteratorImpl[E comparable](head ListNode[E]) collections.Iterator[E] {

	currentNode := head

	for currentNode.GetNext() != nil {
		currentNode = currentNode.GetNext()
	}
	return &DescendingIteratorImpl[E]{
		current: currentNode,
	}
}
func (iter *DescendingIteratorImpl[E]) HasNext() bool {
	return iter.current != nil
}
func (iter *DescendingIteratorImpl[E]) Next() (*E, error) {
	if iter.current == nil {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	val := iter.current.GetData()
	iter.current = iter.current.GetPrev()
	return val, nil
}

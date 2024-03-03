package lists

import (
	"errors"

	"github.com/chiranjeevipavurala/gocollections/errorcodes"
	"github.com/chiranjeevipavurala/gocollections/sets"
)

type LinkedList[E comparable] struct {
	head ListNode[E]
	size int
}

func NewLinkedList[E comparable]() List[E] {
	return &LinkedList[E]{
		head: nil,
		size: 0,
	}
}
func NewLinkedListWithInitialCollection[E comparable](values []E) List[E] {
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

func (l *LinkedList[E]) AddAll(elements []E) bool {
	for _, val := range elements {
		l.Add(val)
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
func (l *LinkedList[E]) ContainsAll(collection []E) (bool, error) {

	if collection == nil {
		return false, errors.New(string(errorcodes.NullPointerError))
	}
	if len(collection) == 0 {
		return false, nil
	}

	for _, val := range collection {
		if !l.Contains(val) {
			return false, nil
		}
	}
	return true, nil
}
func (l *LinkedList[E]) Equals(collection []E) bool {
	if l.size != len(collection) {
		return false
	}
	current := l.head
	for _, val := range collection {
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
func (l *LinkedList[E]) Iterator() sets.Iterator[E] {
	return NewListIteratorImpl[E](l.head)
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
		l.head.SetPrev(nil)
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
func (l *LinkedList[E]) SubList(fromIndex int, toIndex int) (List[E], error) {
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

type ListIteratorImpl[E comparable] struct {
	current ListNode[E]
}

func NewListIteratorImpl[E comparable](head ListNode[E]) sets.Iterator[E] {
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

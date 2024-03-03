package lists

import (
	"errors"

	"github.com/chiranjeevipavurala/gocollections/errorcodes"
	"github.com/chiranjeevipavurala/gocollections/sets"
)

type ArrayList[E comparable] struct {
	values []E
}

func NewArrayList[E comparable]() List[E] {
	values := make([]E, 0)
	return &ArrayList[E]{
		values: values,
	}
}

func NewArrayListWithInitialCapacity[E comparable](capacity int) List[E] {
	values := make([]E, capacity)
	return &ArrayList[E]{
		values: values,
	}
}
func NewArrayListWithInitialCollection[E comparable](values []E) List[E] {
	return &ArrayList[E]{
		values: values,
	}
}

func (a *ArrayList[E]) Add(element E) bool {
	a.values = append(a.values, element)
	return true
}
func (a *ArrayList[E]) AddAtIndex(index int, element E) error {

	if index > len(a.values) || index < 0 {
		return errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	values := make([]E, len(a.values)+1)
	for i := 0; i < index; i++ {
		values[i] = a.values[i]
	}
	values[index] = element
	for i := index; i < len(a.values); i++ {
		values[i+1] = a.values[i]
	}
	a.values = values

	return nil
}
func (a *ArrayList[E]) AddFirst(element E) {

	_ = a.AddAtIndex(0, element)
}
func (a *ArrayList[E]) AddLast(element E) {

	_ = a.AddAtIndex(len(a.values), element)
}

func (a *ArrayList[E]) AddAll(elements []E) bool {
	a.values = append(a.values, elements...)
	return true
}
func (a *ArrayList[E]) AddAllAtIndex(index int, elements []E) (bool, error) {
	if index > len(a.values) || index < 0 {
		return false, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	if elements == nil {
		return false, errors.New(string(errorcodes.NullPointerError))
	}
	values := make([]E, len(a.values)+len(elements))
	for i := 0; i < index; i++ {
		values[i] = a.values[i]
	}
	for i := index; i < index+len(elements); i++ {
		values[i] = elements[i-index]
	}
	for i := index + len(elements); i < len(values); i++ {
		values[i] = a.values[i-len(elements)]
	}
	a.values = values

	return true, nil
}
func (a *ArrayList[E]) Clear() {
	a.values = make([]E, 0)
}

func (a *ArrayList[E]) Contains(element E) bool {
	for _, val := range a.values {
		if val == element {
			return true
		}
	}
	return false
}
func (a *ArrayList[E]) ContainsAll(elements []E) (bool, error) {

	if elements == nil {
		return false, errors.New(string(errorcodes.NullPointerError))
	}
	if len(elements) == 0 {
		return false, nil
	}
	for _, val := range elements {
		if !a.Contains(val) {
			return false, nil
		}
	}
	return true, nil
}

func (a *ArrayList[E]) Equals(elements []E) bool {
	if len(a.values) != len(elements) {
		return false
	}
	for i := 0; i < len(a.values); i++ {
		if a.values[i] != elements[i] {
			return false
		}
	}
	return true
}
func (a *ArrayList[E]) IndexOf(element E) int {
	for i, val := range a.values {
		if val == element {
			return i
		}
	}
	return -1
}
func (a *ArrayList[E]) LastIndexOf(element E) int {
	for i := len(a.values) - 1; i >= 0; i-- {
		if a.values[i] == element {
			return i
		}
	}
	return -1
}

func (a *ArrayList[E]) Get(index int) (*E, error) {
	if index >= len(a.values) || index < 0 {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	return &a.values[index], nil
}
func (a *ArrayList[E]) IsEmpty() bool {
	return len(a.values) == 0
}
func (a *ArrayList[E]) Iterator() sets.Iterator[E] {
	iterator := sets.NewIterator[E](a.values)
	return iterator
}
func (a *ArrayList[E]) Size() int {
	return len(a.values)
}
func (a *ArrayList[E]) Remove(element E) bool {
	index := a.IndexOf(element)
	if index == -1 {
		return false
	}
	_, _ = a.RemoveAtIndex(index)
	return true
}
func (a *ArrayList[E]) RemoveAtIndex(index int) (*E, error) {

	if index >= len(a.values) || index < 0 {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	values := make([]E, len(a.values)-1)
	res := a.values[index]
	for i := 0; i < index; i++ {
		values[i] = a.values[i]
	}
	for i := index; i < len(a.values)-1; i++ {
		values[i] = a.values[i+1]
	}
	a.values = values
	return &res, nil
}
func (a *ArrayList[E]) RemoveFirst() (*E, error) {
	if len(a.values) == 0 {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	return a.RemoveAtIndex(0)

}
func (a *ArrayList[E]) RemoveLast() (*E, error) {
	if len(a.values) == 0 {
		return nil, errors.New(string(errorcodes.NoSuchElementError))
	}
	return a.RemoveAtIndex(len(a.values) - 1)
}

func (a *ArrayList[E]) Set(index int, element E) (*E, error) {

	if index >= len(a.values) || index < 0 {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	a.values[index] = element
	return &element, nil
}
func (a *ArrayList[E]) ToArray() []E {
	return a.values
}
func (a *ArrayList[E]) SubList(fromIndex int, toIndex int) (List[E], error) {
	if fromIndex < 0 || toIndex > len(a.values) {
		return nil, errors.New(string(errorcodes.IndexOutOfBoundsError))
	}
	if fromIndex > toIndex {
		return nil, errors.New(string(errorcodes.IllegalArgumentError))
	}
	values := make([]E, toIndex-fromIndex)
	for i := fromIndex; i < toIndex; i++ {
		values[i-fromIndex] = a.values[i]
	}
	return &ArrayList[E]{
		values: values,
	}, nil
}

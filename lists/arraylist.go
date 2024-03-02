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
		return false, errors.New(string(errorcodes.NullPointerException))
	}

	return true, nil
}
func (a *ArrayList[E]) Clear() {
	a.values = make([]E, 0)
}
func (a *ArrayList[E]) Get(index int) E {
	return a.values[index]
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
		return nil, errors.New(string(errorcodes.NoSuchElementExceptionError))
	}
	return a.RemoveAtIndex(0)

}
func (a *ArrayList[E]) RemoveLast() (*E, error) {
	if len(a.values) == 0 {
		return nil, errors.New(string(errorcodes.NoSuchElementExceptionError))
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

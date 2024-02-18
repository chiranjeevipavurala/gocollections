package sets

type Iterator[E comparable] interface {
	HasNext() bool
	Next() E
}

type IteratorImpl[E comparable] struct {
	currentIndex int
	Values       []E
}

func NewIterator[E comparable](collection []E) Iterator[E] {
	iterator := IteratorImpl[E]{
		currentIndex: 0,
		Values:       collection,
	}
	return &iterator
}
func (iter *IteratorImpl[E]) HasNext() bool {
	return iter.currentIndex != len(iter.Values)
}
func (iter *IteratorImpl[E]) Next() E {
	val := iter.Values[iter.currentIndex]
	iter.currentIndex = iter.currentIndex + 1
	return val
}

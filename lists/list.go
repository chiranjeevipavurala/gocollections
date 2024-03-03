package lists

import "github.com/chiranjeevipavurala/gocollections/sets"

type List[E comparable] interface {
	Add(val E) bool

	AddAtIndex(index int, val E) error
	AddFirst(val E)
	AddLast(val E)
	AddAll([]E) bool

	AddAllAtIndex(index int, collection []E) (bool, error)
	Clear()
	Contains(val E) bool
	ContainsAll(collection []E) (bool, error)
	Equals(collection []E) bool
	Get(index int) (*E, error)
	//HashCode() int
	IndexOf(val E) int
	IsEmpty() bool
	Iterator() sets.Iterator[E]
	LastIndexOf(val E) int
	//listIterator()
	//listIterator(int index)
	RemoveAtIndex(index int) (*E, error)
	RemoveFirst() (*E, error)
	RemoveLast() (*E, error)
	Remove(val E) bool
	//	removeAll(Collection<?> c)
	//replaceAll(UnaryOperator<E> operator)
	//	retainAll(Collection<?> c)
	Set(index int, element E) (*E, error)

	Size() int
	ToArray() []E
	SubList(fromIndex int, toIndex int) (List[E], error)
	/*
		//	sort(Comparator<? super E> c)
		//	spliterator()


	*/
}

type ListNode[E comparable] interface {
	GetNext() ListNode[E]
	SetNext(ListNode[E])
	GetData() *E
	SetData(E)
	GetPrev() ListNode[E]
	SetPrev(ListNode[E])
}
type ListNodeImpl[E comparable] struct {
	data E
	next ListNode[E]
	prev ListNode[E]
}

func NewListNodeImpl[E comparable](val E) ListNode[E] {
	return &ListNodeImpl[E]{
		data: val,
	}
}
func (t *ListNodeImpl[E]) GetData() *E {
	return &t.data
}

func (t *ListNodeImpl[E]) SetData(val E) {
	t.data = val
}

func (t *ListNodeImpl[E]) SetNext(val ListNode[E]) {
	t.next = val
}
func (t *ListNodeImpl[E]) GetNext() ListNode[E] {
	return t.next
}
func (t *ListNodeImpl[E]) SetPrev(val ListNode[E]) {
	t.prev = val
}
func (t *ListNodeImpl[E]) GetPrev() ListNode[E] {
	return t.prev
}

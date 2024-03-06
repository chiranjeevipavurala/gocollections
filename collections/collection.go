package collections

type Collection[E comparable] interface {
	Add(element E) bool
	AddAll(collection Collection[E]) bool
	Clear()
	Contains(element E) bool
	ContainsAll(collection Collection[E]) (bool, error)
	Equals(collection Collection[E]) bool
	//HashCode() int
	IsEmpty() bool
	Iterator() Iterator[E]
	//ParallelStream() Stream[E]
	Remove(element E) bool
	RemoveAll(collection Collection[E]) bool
	//RemoveIf(Predicate<? super E> filter)
	//RetainAll(collection Collection[E]) bool
	Size() int
	//Spliterator() Spliterator[E]
	//Stream() Stream[E]
	ToArray() []E
	//ToString() string
}
type SequencedCollection[E comparable] interface {
	Collection[E]
	AddFirst(val E)
	AddLast(val E)
	GetFirst() (*E, error)
	GetLast() (*E, error)
	RemoveFirst() (*E, error)
	RemoveLast() (*E, error)
	Reversed()
}
type List[E comparable] interface {
	SequencedCollection[E]

	AddAtIndex(index int, val E) error
	AddAllAtIndex(index int, collection []E) (bool, error)
	Get(index int) (*E, error)
	IndexOf(val E) int
	LastIndexOf(val E) int
	//listIterator()
	//listIterator(int index)
	RemoveAtIndex(index int) (*E, error)
	//replaceAll(UnaryOperator<E> operator)
	Set(index int, element E) (*E, error)
	SubList(fromIndex int, toIndex int) (List[E], error)
	Sort(comparator Comparator[E])
}

type Queue[E comparable] interface {
	Collection[E]
	Element() (*E, error)
	Offer(val E) bool
	Peek() (*E, error)
	Poll() (*E, error)
	RemoveHead() (*E, error)
}

type Deque[E comparable] interface {
	Queue[E]
	SequencedCollection[E]
	DescendingIterator() Iterator[E]
	OfferFirst(val E) bool
	OfferLast(val E) bool
	PeekFirst() (*E, error)
	PeekLast() (*E, error)
	PollFirst() (*E, error)
	PollLast() (*E, error)
	Pop() (*E, error)
	Push(val E)
	RemoveFirstOccurrence(val E) bool
	RemoveLastOccurrence(val E) bool
}

type Iterator[E comparable] interface {
	HasNext() bool
	Next() (*E, error)
}
type Comparator[E comparable] interface {
	Compare(a, b E) int
}

type Set[T comparable] interface {
	Collection[T]
	//CopyOf(collection Collection[T]) Set[T]
}

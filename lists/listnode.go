package lists

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

package trees

type TreeNode[E comparable] interface {
	GetLeft() TreeNode[E]
	GetRight() TreeNode[E]
	GetData() E
	SetData(E)
	SetLeft(val E)
	SetRight(val E)
}
type TreeNodeImpl[E comparable] struct {
	data  E
	left  TreeNode[E]
	right TreeNode[E]
}

func NewTreeNodeImpl[E comparable](val E) TreeNode[E] {
	return &TreeNodeImpl[E]{
		data: val,
	}
}
func (t *TreeNodeImpl[E]) GetData() E {
	return t.data
}
func (t *TreeNodeImpl[E]) GetLeft() TreeNode[E] {
	return t.left
}
func (t *TreeNodeImpl[E]) GetRight() TreeNode[E] {
	return t.right
}
func (t *TreeNodeImpl[E]) SetLeft(val E) {
	t.left = NewTreeNodeImpl[E](val)
}
func (t *TreeNodeImpl[E]) SetRight(val E) {
	t.right = NewTreeNodeImpl[E](val)
}
func (t *TreeNodeImpl[E]) SetData(val E) {
	t.data = val
}

package maps

import (
	"errors"
	"sync"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/chiranjeevipavurala/gocollections/sets"
)

// Color represents the color of a node in the Red-Black tree
type Color bool

const (
	Red   Color = true
	Black Color = false
)

// Node represents a node in the Red-Black tree
type Node[K comparable, V comparable] struct {
	key    K
	value  V
	color  Color
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
}

// TreeMap implements a Red-Black tree based map
type TreeMap[K comparable, V comparable] struct {
	root       *Node[K, V]
	size       int
	comparator collections.Comparator[K]
	mu         sync.RWMutex
}

// SortedMap interface for sorted operations
type SortedMap[K comparable, V comparable] interface {
	collections.Map[K, V]
	FirstKey() (*K, error)
	LastKey() (*K, error)
	LowerKey(key K) (*K, error)
	HigherKey(key K) (*K, error)
}

// NewTreeMap creates a new TreeMap with the given comparator
func NewTreeMap[K comparable, V comparable](comparator collections.Comparator[K]) SortedMap[K, V] {
	if comparator == nil {
		return nil
	}
	return &TreeMap[K, V]{
		comparator: comparator,
	}
}

// Clear removes all elements from the map
func (t *TreeMap[K, V]) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = nil
	t.size = 0
}

// ContainsKey returns true if the map contains the given key
func (t *TreeMap[K, V]) ContainsKey(key K) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.getNode(key) != nil
}

// ContainsValue returns true if the map contains the specified value.
func (t *TreeMap[K, V]) ContainsValue(value V) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.containsValue(t.root, value)
}

// containsValue recursively checks if the value exists in the tree.
func (t *TreeMap[K, V]) containsValue(node *Node[K, V], value V) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	return t.containsValue(node.left, value) || t.containsValue(node.right, value)
}

// verifyRedBlackProperties checks if the tree maintains Red-Black properties:
// 1. Every node is either red or black
// 2. Root is black
// 3. If a node is red, both its children are black
// 4. Every path from root to leaves contains the same number of black nodes
func (t *TreeMap[K, V]) verifyRedBlackProperties() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if t.root == nil {
		return true
	}

	// Property 2: Root must be black
	if t.root.color == Red {
		return false
	}

	// Check all paths have same black height
	blackCount := -1
	return t.verifyBlackHeight(t.root, 0, &blackCount)
}

// verifyBlackHeight recursively checks black height and red-black properties
func (t *TreeMap[K, V]) verifyBlackHeight(node *Node[K, V], count int, blackCount *int) bool {
	if node == nil {
		if *blackCount == -1 {
			*blackCount = count
			return true
		}
		return count == *blackCount
	}

	// Property 3: If node is red, both children must be black
	if node.color == Red {
		if (node.left != nil && node.left.color == Red) || (node.right != nil && node.right.color == Red) {
			return false
		}
	}

	// Count black nodes
	if node.color == Black {
		count++
	}

	return t.verifyBlackHeight(node.left, count, blackCount) && t.verifyBlackHeight(node.right, count, blackCount)
}

// EntrySet returns a set of all entries in the map
func (t *TreeMap[K, V]) EntrySet() collections.Set[collections.MapEntry[K, V]] {
	t.mu.RLock()
	defer t.mu.RUnlock()
	entries := sets.NewHashSet[collections.MapEntry[K, V]]()
	t.collectEntries(t.root, entries)
	return entries
}

// Get returns the value associated with the given key
func (t *TreeMap[K, V]) Get(key K) *V {
	t.mu.RLock()
	defer t.mu.RUnlock()
	node := t.getNode(key)
	if node == nil {
		return nil
	}
	return &node.value
}

// IsEmpty returns true if the map is empty
func (t *TreeMap[K, V]) IsEmpty() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.size == 0
}

// KeySet returns a set of all keys in the map
func (t *TreeMap[K, V]) KeySet() collections.Set[K] {
	t.mu.RLock()
	defer t.mu.RUnlock()
	keys := sets.NewHashSet[K]()
	t.collectKeys(t.root, keys)
	return keys
}

// PutAll copies all of the mappings from the specified map to this map
func (t *TreeMap[K, V]) PutAll(m collections.Map[K, V]) {
	if m == nil {
		return
	}

	// Get all entries first to minimize lock time
	entries := m.EntrySet().ToArray()

	t.mu.Lock()
	defer t.mu.Unlock()

	for _, entry := range entries {
		key := entry.GetKey()
		value := entry.GetValue()

		// Check if key exists first
		node := t.getNode(key)
		if node != nil {
			node.value = value
			continue
		}

		// Create new node
		newNode := &Node[K, V]{
			key:   key,
			value: value,
			color: Red,
		}

		// Insert the node
		if t.root == nil {
			t.root = newNode
			t.root.color = Black
			t.size++
			continue
		}

		// Find insertion point
		current := t.root
		var parent *Node[K, V]
		for current != nil {
			parent = current
			cmp := t.comparator.Compare(key, current.key)
			if cmp < 0 {
				current = current.left
			} else {
				current = current.right
			}
		}

		// Insert the node
		newNode.parent = parent
		if t.comparator.Compare(key, parent.key) < 0 {
			parent.left = newNode
		} else {
			parent.right = newNode
		}

		// Fix the tree
		t.fixInsert(newNode)
		t.size++
	}
}

// Put associates the specified value with the specified key in this map
func (t *TreeMap[K, V]) Put(key K, value V) V {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Check if key exists first
	node := t.getNode(key)
	if node != nil {
		oldValue := node.value
		node.value = value
		return oldValue
	}

	// Create new node
	newNode := &Node[K, V]{
		key:   key,
		value: value,
		color: Red,
	}

	// Insert the node
	if t.root == nil {
		t.root = newNode
		t.root.color = Black
		t.size++
		var zero V
		return zero
	}

	// Find insertion point
	current := t.root
	var parent *Node[K, V]
	for current != nil {
		parent = current
		cmp := t.comparator.Compare(key, current.key)
		if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	// Insert the node
	newNode.parent = parent
	if t.comparator.Compare(key, parent.key) < 0 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	// Fix the tree
	t.fixInsert(newNode)
	t.size++

	var zero V
	return zero
}

func (t *TreeMap[K, V]) fixInsert(node *Node[K, V]) {
	if node == nil {
		return
	}

	for node != t.root && node.parent != nil && node.parent.color == Red {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if uncle != nil && uncle.color == Red {
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					t.rotateLeft(node)
				}
				if node.parent != nil {
					node.parent.color = Black
					if node.parent.parent != nil {
						node.parent.parent.color = Red
						t.rotateRight(node.parent.parent)
					}
				}
			}
		} else {
			uncle := node.parent.parent.left
			if uncle != nil && uncle.color == Red {
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					t.rotateRight(node)
				}
				if node.parent != nil {
					node.parent.color = Black
					if node.parent.parent != nil {
						node.parent.parent.color = Red
						t.rotateLeft(node.parent.parent)
					}
				}
			}
		}
	}
	t.root.color = Black
}

// ReplaceKeyWithValue replaces the entry for the specified key only if currently mapped to the given old value
func (t *TreeMap[K, V]) ReplaceKeyWithValue(key K, oldValue V, newValue V) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.getNode(key)
	if node == nil || node.value != oldValue {
		return false
	}

	node.value = newValue
	return true
}

// Replace replaces the entry for the specified key only if it is currently mapped to some value
func (t *TreeMap[K, V]) Replace(key K, value V) V {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.getNode(key)
	if node == nil {
		var zero V
		return zero
	}

	oldValue := node.value
	node.value = value
	return oldValue
}

// RemoveKeyWithValue removes the entry for the specified key only if it is currently mapped to the specified value
func (t *TreeMap[K, V]) RemoveKeyWithValue(key K, value V) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.getNode(key)
	if node == nil || node.value != value {
		return false
	}

	t.root = t.deleteNode(node)
	if t.root != nil {
		t.root.color = Black
	}
	t.size--
	return true
}

// Remove removes the mapping for a key from this map if it is present
func (t *TreeMap[K, V]) Remove(key K) V {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.getNode(key)
	if node == nil {
		var zero V
		return zero
	}

	value := node.value
	t.root = t.deleteNode(node)
	if t.root != nil {
		t.root.color = Black
	}
	t.size--
	return value
}

// Size returns the number of elements in the map
func (t *TreeMap[K, V]) Size() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.size
}

// Values returns a collection of all values in the map
func (t *TreeMap[K, V]) Values() collections.Collection[V] {
	t.mu.RLock()
	defer t.mu.RUnlock()
	values := sets.NewHashSet[V]()
	t.collectValues(t.root, values)
	return values
}

// Equals returns true if this map equals the given map
func (t *TreeMap[K, V]) Equals(other any) bool {
	if other == nil {
		return false
	}
	otherMap, ok := other.(collections.Map[K, V])
	if !ok {
		return false
	}
	if t.Size() != otherMap.Size() {
		return false
	}
	for _, entry := range t.EntrySet().ToArray() {
		otherValue := otherMap.Get(entry.GetKey())
		if otherValue == nil || *otherValue != entry.GetValue() {
			return false
		}
	}
	return true
}

// FirstKey returns the first (lowest) key in the map
func (t *TreeMap[K, V]) FirstKey() (*K, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.root == nil {
		return nil, errors.New("map is empty")
	}
	node := t.getFirstNode(t.root)
	return &node.key, nil
}

// LastKey returns the last (highest) key in the map
func (t *TreeMap[K, V]) LastKey() (*K, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.root == nil {
		return nil, errors.New("map is empty")
	}
	node := t.getLastNode(t.root)
	return &node.key, nil
}

// LowerKey returns the greatest key less than the given key
func (t *TreeMap[K, V]) LowerKey(key K) (*K, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.root == nil {
		return nil, errors.New("map is empty")
	}
	node := t.getLowerNode(key)
	if node == nil {
		return nil, errors.New("no lower key exists")
	}
	return &node.key, nil
}

// HigherKey returns the least key greater than the given key
func (t *TreeMap[K, V]) HigherKey(key K) (*K, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.root == nil {
		return nil, errors.New("map is empty")
	}
	node := t.getHigherNode(key)
	if node == nil {
		return nil, errors.New("no higher key exists")
	}
	return &node.key, nil
}

// HasKey returns true if this map contains a mapping for the specified key
func (t *TreeMap[K, V]) HasKey(key K) bool {
	return t.ContainsKey(key)
}

// HasValue returns true if this map maps one or more keys to the specified value
func (t *TreeMap[K, V]) HasValue(value V) bool {
	return t.ContainsValue(value)
}

// Helper methods for Red-Black tree operations
func (t *TreeMap[K, V]) getNode(key K) *Node[K, V] {
	node := t.root
	for node != nil {
		cmp := t.comparator.Compare(key, node.key)
		if cmp == 0 {
			return node
		}
		if cmp < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func (t *TreeMap[K, V]) deleteNode(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}

	// If node has two children, find successor
	if node.left != nil && node.right != nil {
		successor := t.getFirstNode(node.right)
		if successor == nil {
			return nil
		}
		node.key = successor.key
		node.value = successor.value
		node = successor
	}

	// Get the child node
	var child *Node[K, V]
	if node.left != nil {
		child = node.left
	} else {
		child = node.right
	}

	// Handle color and fix-up
	if node.color == Black {
		if child != nil {
			child.color = Black
		} else {
			t.fixDelete(node)
		}
	}

	// Replace the node with its child
	if node.parent == nil {
		t.root = child
	} else if node == node.parent.left {
		node.parent.left = child
	} else {
		node.parent.right = child
	}

	if child != nil {
		child.parent = node.parent
	}

	return t.root
}

func (t *TreeMap[K, V]) fixDelete(node *Node[K, V]) {
	for node != t.root && node.color == Black {
		if node == node.parent.left {
			sibling := node.parent.right
			if sibling == nil {
				break
			}
			if sibling.color == Red {
				sibling.color = Black
				node.parent.color = Red
				t.rotateLeft(node.parent)
				sibling = node.parent.right
			}
			if (sibling.left == nil || sibling.left.color == Black) &&
				(sibling.right == nil || sibling.right.color == Black) {
				sibling.color = Red
				node = node.parent
			} else {
				if sibling.right == nil || sibling.right.color == Black {
					if sibling.left != nil {
						sibling.left.color = Black
					}
					sibling.color = Red
					t.rotateRight(sibling)
					sibling = node.parent.right
				}
				sibling.color = node.parent.color
				node.parent.color = Black
				if sibling.right != nil {
					sibling.right.color = Black
				}
				t.rotateLeft(node.parent)
				node = t.root
			}
		} else {
			sibling := node.parent.left
			if sibling == nil {
				break
			}
			if sibling.color == Red {
				sibling.color = Black
				node.parent.color = Red
				t.rotateRight(node.parent)
				sibling = node.parent.left
			}
			if (sibling.right == nil || sibling.right.color == Black) &&
				(sibling.left == nil || sibling.left.color == Black) {
				sibling.color = Red
				node = node.parent
			} else {
				if sibling.left == nil || sibling.left.color == Black {
					if sibling.right != nil {
						sibling.right.color = Black
					}
					sibling.color = Red
					t.rotateLeft(sibling)
					sibling = node.parent.left
				}
				sibling.color = node.parent.color
				node.parent.color = Black
				if sibling.left != nil {
					sibling.left.color = Black
				}
				t.rotateRight(node.parent)
				node = t.root
			}
		}
	}
	node.color = Black
}

func (t *TreeMap[K, V]) rotateLeft(node *Node[K, V]) {
	right := node.right
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.parent = node.parent
	if node.parent == nil {
		t.root = right
	} else if node == node.parent.left {
		node.parent.left = right
	} else {
		node.parent.right = right
	}
	right.left = node
	node.parent = right
}

func (t *TreeMap[K, V]) rotateRight(node *Node[K, V]) {
	left := node.left
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.parent = node.parent
	if node.parent == nil {
		t.root = left
	} else if node == node.parent.right {
		node.parent.right = left
	} else {
		node.parent.left = left
	}
	left.right = node
	node.parent = left
}

func (t *TreeMap[K, V]) getFirstNode(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		node = t.root
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *TreeMap[K, V]) getLastNode(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		node = t.root
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func (t *TreeMap[K, V]) getLowerNode(key K) *Node[K, V] {
	node := t.root
	var result *Node[K, V]
	for node != nil {
		cmp := t.comparator.Compare(key, node.key)
		if cmp > 0 {
			result = node
			node = node.right
		} else {
			node = node.left
		}
	}
	return result
}

func (t *TreeMap[K, V]) getHigherNode(key K) *Node[K, V] {
	node := t.root
	var result *Node[K, V]
	for node != nil {
		cmp := t.comparator.Compare(key, node.key)
		if cmp < 0 {
			result = node
			node = node.left
		} else {
			node = node.right
		}
	}
	return result
}

func (t *TreeMap[K, V]) collectEntries(node *Node[K, V], entries collections.Set[collections.MapEntry[K, V]]) {
	if node == nil {
		return
	}
	t.collectEntries(node.left, entries)
	entries.Add(collections.NewHashMapEntry(node.key, node.value))
	t.collectEntries(node.right, entries)
}

func (t *TreeMap[K, V]) collectKeys(node *Node[K, V], keys collections.Set[K]) {
	if node == nil {
		return
	}
	t.collectKeys(node.left, keys)
	keys.Add(node.key)
	t.collectKeys(node.right, keys)
}

func (t *TreeMap[K, V]) collectValues(node *Node[K, V], values collections.Set[V]) {
	if node == nil {
		return
	}
	t.collectValues(node.left, values)
	values.Add(node.value)
	t.collectValues(node.right, values)
}

// PutIfAbsent associates the specified value with the specified key in this map if the key is not already associated with a value
func (t *TreeMap[K, V]) PutIfAbsent(key K, value V) V {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.getNode(key)
	if node != nil {
		return node.value
	}

	// Create new node
	newNode := &Node[K, V]{
		key:   key,
		value: value,
		color: Red,
	}

	// Insert the node
	if t.root == nil {
		t.root = newNode
		t.root.color = Black
		t.size++
		var zero V
		return zero
	}

	// Find insertion point
	current := t.root
	var parent *Node[K, V]
	for current != nil {
		parent = current
		cmp := t.comparator.Compare(key, current.key)
		if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	// Insert the node
	newNode.parent = parent
	if t.comparator.Compare(key, parent.key) < 0 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	// Fix the tree
	t.fixInsert(newNode)
	t.size++

	var zero V
	return zero
}

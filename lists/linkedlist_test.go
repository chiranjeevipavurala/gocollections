package lists

import (
	"reflect"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", *val)
	}
}

func TestLinkedList_Clear(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	list.Clear()

	if list.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size())
	}
}

func TestLinkedList_Get(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", *val)
	}

	_, err = list.Get(3)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	_, err = list.Get(-1)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	list := LinkedList[int]{}
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}

	list.Add(1)
	if list.IsEmpty() {
		t.Errorf("Expected list not to be empty")
	}
}

func TestLinkedList_Size(t *testing.T) {
	list := LinkedList[int]{}
	if list.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size())
	}

	list.Add(1)
	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}
}

func TestLinkedList_ToArray(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	arr := list.ToArray()
	if len(arr) != 3 {
		t.Errorf("Expected array length to be 3, got %d", len(arr))
	}

	if arr[0] != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", arr[0])
	}

	if arr[1] != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", arr[1])
	}

	if arr[2] != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", arr[2])
	}
}

func TestLinkedList_Set(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	_, err := list.Set(1, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	val, err := list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 1 to be 4, got %d", *val)
	}

	_, err = list.Set(3, 5)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}
}

func TestLinkedList_Contains(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if !list.Contains(2) {
		t.Errorf("Expected list to contain element 2")
	}

	if list.Contains(4) {
		t.Errorf("Expected list not to contain element 4")
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.IndexOf(2) != 1 {
		t.Errorf("Expected index of 2 to be 1, got %d", list.IndexOf(2))
	}

	if list.IndexOf(4) != -1 {
		t.Errorf("Expected index of 4 to be -1, got %d", list.IndexOf(4))
	}
}

func TestLinkedList_LastIndexOf(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(2)

	if list.LastIndexOf(2) != 3 {
		t.Errorf("Expected last index of 2 to be 3, got %d", list.LastIndexOf(2))
	}

	if list.LastIndexOf(4) != -1 {
		t.Errorf("Expected last index of 4 to be -1, got %d", list.LastIndexOf(4))
	}
}

func TestLinkedList_RemoveAtIndex(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	removedElement, err := list.RemoveAtIndex(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *removedElement != 4 {
		t.Errorf("Expected removed element to be 4, got %d", *removedElement)
	}

	removedElement, err = list.RemoveAtIndex(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *removedElement != 2 {
		t.Errorf("Expected removed element to be 2, got %d", *removedElement)
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", *val)
	}

	_, err = list.RemoveAtIndex(2)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	_, err = list.RemoveAtIndex(-1)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	_, err = list.RemoveAtIndex(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

}

func TestLinkedList_RemoveFirst(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	removedElement, err := list.RemoveFirst()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *removedElement != 1 {
		t.Errorf("Expected removed element to be 1, got %d", *removedElement)
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %d", val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", val)
	}

	_, err = list.RemoveFirst()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.RemoveFirst()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.RemoveFirst()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	removedElement, err := list.RemoveLast()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *removedElement != 3 {
		t.Errorf("Expected removed element to be 3, got %d", *removedElement)
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	_, err = list.RemoveLast()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.RemoveLast()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.RemoveLast()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}

func TestLinkedList_AddAtIndex(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.AddAtIndex(1, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.Size() != 4 {
		t.Errorf("Expected size to be 4, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 1 to be 4, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 2 to be 2, got %d", *val)
	}

	val, err = list.Get(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 3 to be 3, got %d", *val)
	}

	err = list.AddAtIndex(5, 5)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	err = list.AddAtIndex(-1, 5)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	err = list.AddAtIndex(0, 5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = list.AddAtIndex(4, 6)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
func TestLinkedList_ContainsAll(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test with a collection that contains all elements
	collection := NewLinkedListWithInitialCollection[int]([]int{1, 2, 3})
	result, err := list.ContainsAll(collection)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected ContainsAll to return true, got false")
	}

	// Test with a collection that does not contain all elements
	collection = NewLinkedListWithInitialCollection[int]([]int{1, 2, 4})
	result, err = list.ContainsAll(collection)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected ContainsAll to return false, got true")
	}

	// Test with an empty collection
	collection = NewLinkedListWithInitialCollection[int]([]int{})
	result, err = list.ContainsAll(collection)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected ContainsAll to return false, got true")
	}

	// Test with a nil collection
	result, err = list.ContainsAll(nil)
	if err == nil {
		t.Errorf("Expected NullPointerError, got no error")
	}
	if result {
		t.Errorf("Expected ContainsAll to return false, got true")
	}
}
func TestLinkedList_Equals(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test with a collection that is equal to the list
	collection := NewLinkedListWithInitialCollection[int]([]int{1, 2, 3})
	result := list.Equals(collection)
	if !result {
		t.Errorf("Expected Equals to return true, got false")
	}

	// Test with a collection that is not equal to the list
	collection = NewLinkedListWithInitialCollection[int]([]int{1, 2, 4})
	result = list.Equals(collection)
	if result {
		t.Errorf("Expected Equals to return false, got true")
	}

	// Test with a collection that has a different size than the list
	collection = NewLinkedListWithInitialCollection[int]([]int{1, 2})
	result = list.Equals(collection)
	if result {
		t.Errorf("Expected Equals to return false, got true")
	}

	// Test with an empty collection
	collection = NewLinkedListWithInitialCollection[int]([]int{})
	result = list.Equals(collection)
	if result {
		t.Errorf("Expected Equals to return false, got true")
	}

	// Test with a nil collection
	result = list.Equals(nil)
	if result {
		t.Errorf("Expected Equals to return false, got true")
	}
}
func TestLinkedList_Remove(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test removing an existing element
	result := list.Remove(2)
	if !result {
		t.Errorf("Expected Remove to return true")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}
	if list.Contains(2) {
		t.Errorf("Expected list not to contain element 2")
	}

	// Test removing a non-existing element
	result = list.Remove(4)
	if result {
		t.Errorf("Expected Remove to return false")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}
}
func TestLinkedList_SubList(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	// Test with valid indices
	subList, err := list.SubList(1, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if subList.Size() != 3 {
		t.Errorf("Expected subList size to be 3, got %d", subList.Size())
	}

	val, err := subList.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %d", *val)
	}

	val, err = subList.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", *val)
	}

	val, err = subList.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 2 to be 4, got %d", *val)
	}

	// Test with invalid indices
	_, err = list.SubList(-1, 3)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	_, err = list.SubList(2, 6)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	_, err = list.SubList(4, 2)
	if err == nil {
		t.Errorf("Expected IllegalArgumentError, got no error")
	}
}

func TestLinkedList_Iterator(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	iterator := list.Iterator()

	// Test HasNext()
	if !iterator.HasNext() {
		t.Errorf("Expected HasNext() to return true, got false")
	}

	// Test Next()
	val, err := iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected value to be 1, got %d", *val)
	}

	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected value to be 2, got %d", *val)
	}

	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected value to be 3, got %d", *val)
	}

	// Test Next() when there are no more elements
	if iterator.HasNext() {
		t.Errorf("Expected HasNext() to return false, got true")
	}

	_, err = iterator.Next()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}
func TestLinkedList_AddFirst(t *testing.T) {
	list := LinkedList[int]{}
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddFirst(3)

	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 0 to be 3, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 2 to be 1, got %d", *val)
	}
}
func TestLinkedList_AddAll(t *testing.T) {
	list := LinkedList[int]{}
	elements := NewLinkedListWithInitialCollection[int]([]int{1, 2, 3})

	result := list.AddAll(elements)
	if !result {
		t.Errorf("Expected AddAll to return true, got false")
	}

	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", *val)
	}
}
func TestLinkedList_AddLast(t *testing.T) {
	list := LinkedList[int]{}
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)

	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", *val)
	}
}
func TestLinkedList_AddAllAtIndex(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	elements := []int{4, 5, 6}

	result, err := list.AddAllAtIndex(1, elements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected AddAllAtIndex to return true, got false")
	}

	if list.Size() != 6 {
		t.Errorf("Expected size to be 6, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 1 to be 4, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 5 {
		t.Errorf("Expected element at index 2 to be 5, got %d", *val)
	}

	val, err = list.Get(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 6 {
		t.Errorf("Expected element at index 3 to be 6, got %d", *val)
	}

	val, err = list.Get(4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 4 to be 2, got %d", *val)
	}

	val, err = list.Get(5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 5 to be 3, got %d", *val)
	}

	result, err = list.AddAllAtIndex(7, elements)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}
	if result {
		t.Errorf("Expected AddAllAtIndex to return false, got true")
	}

	result, err = list.AddAllAtIndex(-1, elements)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}
	if result {
		t.Errorf("Expected AddAllAtIndex to return false, got true")
	}

	result, err = list.AddAllAtIndex(0, []int{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected AddAllAtIndex to return false, got true")
	}
}
func TestLinkedList_RemoveHead(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	removedElement, err := list.RemoveHead()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *removedElement != 1 {
		t.Errorf("Expected removed element to be 1, got %d", *removedElement)
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", *val)
	}

	_, err = list.RemoveHead()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.RemoveHead()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.RemoveHead()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}
func TestLinkedList_Element(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	element, err := list.Element()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 1 {
		t.Errorf("Expected element to be 1, got %d", *element)
	}
	list.Clear()
	_, err = list.Element()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}

func TestLinkedList_Peek(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	element, err := list.Peek()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 1 {
		t.Errorf("Expected element to be 1, got %d", *element)
	}
	list.Clear()
	_, err = list.Peek()
	if err == nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
func TestLinkedList_Poll(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	element, err := list.Poll()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 1 {
		t.Errorf("Expected element to be 1, got %d", *element)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}
}
func TestLinkedList_Offer(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	result := list.Offer(3)
	if !result {
		t.Errorf("Expected Offer to return true, got false")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}
}
func TestLinkedList_OfferFirst(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	result := list.OfferFirst(3)
	if !result {
		t.Errorf("Expected OfferFirst to return true, got false")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}
}

func TestLinkedList_OfferLast(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	result := list.OfferLast(3)
	if !result {
		t.Errorf("Expected OfferLast to return true, got false")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}
}
func TestLinkedList_PeekFirst(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	element, err := list.PeekFirst()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 1 {
		t.Errorf("Expected element to be 1, got %d", *element)
	}
	list.Clear()
	_, err = list.PeekFirst()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}
func TestLinkedList_PeekLast(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	element, err := list.PeekLast()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 2 {
		t.Errorf("Expected element to be 2, got %d", *element)
	}
	list.Clear()
	_, err = list.PeekLast()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}

func TestLinkedList_PollFirst(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	element, err := list.PollFirst()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 1 {
		t.Errorf("Expected element to be 1, got %d", *element)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}
}

func TestLinkedList_PollLast(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	element, err := list.PollLast()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 2 {
		t.Errorf("Expected element to be 2, got %d", *element)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}
}

func TestLinkedList_Reversed(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	list.Reversed()

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 0 to be 3, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 2 to be 1, got %d", *val)
	}
}
func TestLinkedList_GetLast(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	lastElement, err := list.GetLast()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *lastElement != 3 {
		t.Errorf("Expected last element to be 3, got %d", *lastElement)
	}

	list.Clear()

	_, err = list.GetLast()
	if err == nil {
		t.Errorf("Expected NoSuchElementError, got no error")
	}
}
func TestLinkedList_DescendingIterator(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	iterator := list.DescendingIterator()

	// Test hasNext()
	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next element")
	}

	// Test next()
	val, err := iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected next element to be 3, got %d", *val)
	}

	// Test hasNext() after calling next()
	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next element")
	}

	// Test next() again
	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected next element to be 2, got %d", *val)
	}

	// Test hasNext() after calling next() again
	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next element")
	}

	// Test next() for the last element
	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected next element to be 1, got %d", *val)
	}

	// Test hasNext() after calling next() for the last element
	if iterator.HasNext() {
		t.Errorf("Expected iterator to have no next element")
	}

	// Test next() when there are no more elements
	_, err = iterator.Next()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}
func TestLinkedList_Sort(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(3)
	list.Add(1)
	list.Add(2)

	comparator := IntComparator{}
	list.Sort(comparator)

	expected := []int{1, 2, 3}
	arr := list.ToArray()
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected sorted array to be %v, got %v", expected, arr)
	}
}

type IntComparator struct{}

func (c IntComparator) Compare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
func TestLinkedList_Pop(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	poppedElement, err := list.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *poppedElement != 1 {
		t.Errorf("Expected popped element to be 1, got %d", *poppedElement)
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", *val)
	}

	_, err = list.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = list.Pop()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}
func TestLinkedList_Push(t *testing.T) {
	list := LinkedList[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)

	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 0 to be 3, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 2 to be 1, got %d", *val)
	}
}
func TestLinkedList_RemoveFirstOccurrence(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test removing an existing element
	result := list.RemoveFirstOccurrence(2)
	if !result {
		t.Errorf("Expected RemoveFirstOccurrence to return true, got false")
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", *val)
	}

	// Test removing a non-existing element
	result = list.RemoveFirstOccurrence(4)
	if result {
		t.Errorf("Expected RemoveFirstOccurrence to return false, got true")
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err = list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", *val)
	}
}
func TestLinkedList_RemoveLastOccurrence(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(2)

	removed := list.RemoveLastOccurrence(2)
	if !removed {
		t.Errorf("Expected RemoveLastOccurrence to return true")
	}

	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", *val)
	}

	removed = list.RemoveLastOccurrence(4)
	if removed {
		t.Errorf("Expected RemoveLastOccurrence to return false")
	}

	removed = list.RemoveLastOccurrence(3)
	if !removed {
		t.Errorf("Expected RemoveLastOccurrence to return true")
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	val, err = list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	removed = list.RemoveLastOccurrence(1)
	if !removed {
		t.Errorf("Expected RemoveLastOccurrence to return true")
	}

	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}

	val, err = list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %d", *val)
	}

	removed = list.RemoveLastOccurrence(2)
	if !removed {
		t.Errorf("Expected RemoveLastOccurrence to return true")
	}

	if list.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size())
	}

	removed = list.RemoveLastOccurrence(2)
	if removed {
		t.Errorf("Expected RemoveLastOccurrence to return false")
	}
}

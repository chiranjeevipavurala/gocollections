package lists

import (
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	list := ArrayList[int]{}
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

func TestArrayList_AddAll(t *testing.T) {
	list := NewArrayList[int]()
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

	list.AddAll([]int{4, 5, 6})

	if list.Size() != 6 {
		t.Errorf("Expected size to be 6, got %d", list.Size())
	}

	val, err = list.Get(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 3 to be 4, got %d", *val)
	}

	val, err = list.Get(4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 5 {
		t.Errorf("Expected element at index 4 to be 5, got %d", *val)

	}

	val, err = list.Get(5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 6 {
		t.Errorf("Expected element at index 5 to be 6, got %d", *val)
	}

}

func TestArrayList_Clear(t *testing.T) {
	list := NewArrayList[int]()
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

func TestArrayList_Get(t *testing.T) {
	list := NewArrayList[int]()
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
}

func TestArrayList_IsEmpty(t *testing.T) {
	list := NewArrayList[int]()
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}

	list.Add(1)
	if list.IsEmpty() {
		t.Errorf("Expected list not to be empty")
	}
}

func TestArrayList_Size(t *testing.T) {
	list := NewArrayList[int]()
	if list.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size())
	}

	list.Add(1)
	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}
}

func TestArrayList_ToArray(t *testing.T) {
	list := NewArrayList[int]()
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

func TestArrayList_Set(t *testing.T) {
	list := NewArrayList[int]()
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

func TestArrayList_Contains(t *testing.T) {
	list := NewArrayList[int]()
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

func TestArrayList_Equals(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if !list.Equals([]int{1, 2, 3}) {
		t.Errorf("Expected list to equal [1, 2, 3]")
	}

	if list.Equals([]int{1, 2}) {
		t.Errorf("Expected list not to equal [1, 2]")
	}
	if list.Equals([]int{1, 3, 2}) {
		t.Errorf("Expected list not to equal [1, 3, 2]")
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	list := NewArrayList[int]()
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

func TestArrayList_LastIndexOf(t *testing.T) {
	list := NewArrayList[int]()
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

func TestArrayList_RemoveAtIndex(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	removedElement, err := list.RemoveAtIndex(1)
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
}

func TestArrayList_RemoveFirst(t *testing.T) {
	list := NewArrayList[int]()
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

func TestArrayList_RemoveLast(t *testing.T) {
	list := NewArrayList[int]()
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
func TestArrayList_AddAtIndex(t *testing.T) {
	list := NewArrayList[int]()
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
}
func TestArrayList_AddFirst(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	list.AddFirst(0)

	if list.Size() != 4 {
		t.Errorf("Expected size to be 4, got %d", list.Size())
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 0 {
		t.Errorf("Expected element at index 0 to be 0, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 1 to be 1, got %d", *val)
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

}
func TestArrayList_AddLast(t *testing.T) {
	list := NewArrayListWithInitialCollection[int]([]int{1, 2, 3})

	list.AddLast(4)

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

	val, err = list.Get(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 3 to be 4, got %d", *val)
	}
}

func TestArrayList_AddAllAtIndex(t *testing.T) {
	list := NewArrayListWithInitialCapacity[int](3)
	list.Set(0, 1)
	list.Set(1, 2)
	list.Set(2, 3)

	_, err := list.AddAllAtIndex(1, []int{4, 5})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.Size() != 5 {
		t.Errorf("Expected size to be 5, got %d", list.Size())
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
	if *val != 2 {
		t.Errorf("Expected element at index 3 to be 2, got %d", *val)
	}

	val, err = list.Get(4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 4 to be 3, got %d", *val)
	}

	_, err = list.AddAllAtIndex(6, []int{6, 7})
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}
	_, err = list.AddAllAtIndex(1, nil)
	if err == nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
func TestArrayList_Iterator(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	iterator := list.Iterator()
	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next")
	}

	val, err := iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *val != 1 {
		t.Errorf("Expected element to be 1, got %d", *val)
	}

	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next")
	}

	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *val != 2 {
		t.Errorf("Expected element to be 2, got %d", *val)
	}

	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next")
	}

	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if *val != 3 {
		t.Errorf("Expected element to be 3, got %d", *val)
	}

	if iterator.HasNext() {
		t.Errorf("Expected iterator not to have next")
	}

	_, err = iterator.Next()
	if err == nil {
		t.Errorf("Expected NoSuchElementExceptionError, got no error")
	}
}
func TestArrayList_Remove(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if !list.Remove(2) {
		t.Errorf("Expected element to be removed")
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

	if list.Remove(4) {
		t.Errorf("Expected element not to be removed")
	}
}
func TestArrayList_SubList(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	subList, err := list.SubList(1, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if subList.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", subList.Size())
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

	_, err = list.SubList(3, 1)
	if err == nil {
		t.Errorf("Expected IllegalArgumentException, got no error")
	}
	_, err = list.SubList(-1, 1)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsException, got no error")
	}
	_, err = list.SubList(0, 10)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsException, got no error")
	}
}
func TestArrayList_ContainsAll(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test when elements is nil
	result, err := list.ContainsAll(nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if result {
		t.Errorf("Expected result to be false, got true")
	}

	// Test when elements is empty
	result, err = list.ContainsAll([]int{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected result to be false, got true")
	}

	// Test when elements contains all elements in the list
	result, err = list.ContainsAll([]int{1, 2, 3})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected result to be true, got false")
	}

	// Test when elements does not contain all elements in the list
	result, err = list.ContainsAll([]int{1, 2, 4})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected result to be false, got true")
	}
}

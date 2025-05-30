package lists

import (
	"reflect"
	"sync"
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

	// Test setting element at index 0
	oldVal, err := list.Set(0, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *oldVal != 1 {
		t.Errorf("Expected old value at index 0 to be 1, got %d", *oldVal)
	}

	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 0 to be 4, got %d", *val)
	}

	// Test setting element at middle index
	oldVal, err = list.Set(1, 5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *oldVal != 2 {
		t.Errorf("Expected old value at index 1 to be 2, got %d", *oldVal)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 5 {
		t.Errorf("Expected element at index 1 to be 5, got %d", *val)
	}

	// Test setting element at last index
	oldVal, err = list.Set(2, 6)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *oldVal != 3 {
		t.Errorf("Expected old value at index 2 to be 3, got %d", *oldVal)
	}

	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 6 {
		t.Errorf("Expected element at index 2 to be 6, got %d", *val)
	}

	// Test with invalid indices
	_, err = list.Set(3, 5)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}

	_, err = list.Set(-1, 5)
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

func TestLinkedList_RemoveAll(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	elements := NewLinkedListWithInitialCollection[int]([]int{1, 2, 4})

	result := list.RemoveAll(elements)
	if !result {
		t.Errorf("Expected RemoveAll to return true, got false")
	}

	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}

	if list.Contains(1) {
		t.Errorf("Expected list not to contain element 1")
	}

	if list.Contains(2) {
		t.Errorf("Expected list not to contain element 2")
	}

	if !list.Contains(3) {
		t.Errorf("Expected list to contain element 3")
	}

	// Test with empty collection
	emptyElements := NewLinkedListWithInitialCollection[int]([]int{})
	result = list.RemoveAll(emptyElements)
	if result {
		t.Errorf("Expected RemoveAll to return false for empty collection since no elements were removed, got true")
	}
	if list.Size() != 1 {
		t.Errorf("Expected size to remain 1 after removing empty collection, got %d", list.Size())
	}

	// Test with nil collection
	result = list.RemoveAll(nil)
	if result {
		t.Errorf("Expected RemoveAll to return false for nil collection, got true")
	}
	if list.Size() != 1 {
		t.Errorf("Expected size to remain 1 after removing nil collection, got %d", list.Size())
	}

	// Test removing all elements (should set tail to nil)
	allElements := NewLinkedListWithInitialCollection[int]([]int{3})
	result = list.RemoveAll(allElements)
	if !result {
		t.Errorf("Expected RemoveAll to return true when removing last element")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size())
	}
	if list.Contains(3) {
		t.Errorf("Expected list not to contain element 3")
	}
	// Verify list is empty by trying to get first element
	_, err := list.GetFirst()
	if err == nil {
		t.Errorf("Expected NoSuchElementError when getting first element from empty list")
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
	tests := []struct {
		name          string
		initialValues []int
		index         int
		element       int
		wantErr       bool
		wantValues    []int
	}{
		{
			name:          "Add at beginning of empty list",
			initialValues: []int{},
			index:         0,
			element:       1,
			wantErr:       false,
			wantValues:    []int{1},
		},
		{
			name:          "Add at beginning of non-empty list",
			initialValues: []int{2, 3, 4},
			index:         0,
			element:       1,
			wantErr:       false,
			wantValues:    []int{1, 2, 3, 4},
		},
		{
			name:          "Add at end of list",
			initialValues: []int{1, 2, 3},
			index:         3,
			element:       4,
			wantErr:       false,
			wantValues:    []int{1, 2, 3, 4},
		},
		{
			name:          "Add in middle of list",
			initialValues: []int{1, 3, 4},
			index:         1,
			element:       2,
			wantErr:       false,
			wantValues:    []int{1, 2, 3, 4},
		},
		{
			name:          "Add in middle of large list",
			initialValues: []int{1, 2, 4, 5, 6, 7, 8},
			index:         2,
			element:       3,
			wantErr:       false,
			wantValues:    []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:          "Add in middle of list with duplicate elements",
			initialValues: []int{1, 1, 3, 3, 4, 4},
			index:         2,
			element:       2,
			wantErr:       false,
			wantValues:    []int{1, 1, 2, 3, 3, 4, 4},
		},
		{
			name:          "Add in middle of list with negative numbers",
			initialValues: []int{-3, -1, 1, 3},
			index:         1,
			element:       -2,
			wantErr:       false,
			wantValues:    []int{-3, -2, -1, 1, 3},
		},
		{
			name:          "Add at negative index",
			initialValues: []int{1, 2, 3},
			index:         -1,
			element:       0,
			wantErr:       true,
			wantValues:    []int{1, 2, 3},
		},
		{
			name:          "Add at index beyond size",
			initialValues: []int{1, 2, 3},
			index:         4,
			element:       5,
			wantErr:       true,
			wantValues:    []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewLinkedListWithInitialCollection(tt.initialValues)
			err := list.AddAtIndex(tt.index, tt.element)

			if (err != nil) != tt.wantErr {
				t.Errorf("AddAtIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				values := list.ToArray()
				if !reflect.DeepEqual(values, tt.wantValues) {
					t.Errorf("AddAtIndex() values = %v, want %v", values, tt.wantValues)
				}
			}
		})
	}
}

// TestLinkedList_AddAtIndex_Concurrent tests concurrent AddAtIndex operations
func TestLinkedList_AddAtIndex_Concurrent(t *testing.T) {
	list := NewLinkedList[int]()
	const goroutines = 10
	const elementsPerGoroutine = 100
	var wg sync.WaitGroup

	// Test concurrent additions at different indices
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < elementsPerGoroutine; j++ {
				element := goroutineID*elementsPerGoroutine + j
				err := list.AddAtIndex(0, element) // Always add at beginning
				if err != nil {
					t.Errorf("Concurrent AddAtIndex failed: %v", err)
				}
			}
		}(i)
	}

	wg.Wait()

	// Verify final size
	if list.Size() != goroutines*elementsPerGoroutine {
		t.Errorf("Expected size %d, got %d", goroutines*elementsPerGoroutine, list.Size())
	}

	// Verify all elements are present
	values := list.ToArray()
	valueMap := make(map[int]bool)
	for _, v := range values {
		valueMap[v] = true
	}

	for i := 0; i < goroutines; i++ {
		for j := 0; j < elementsPerGoroutine; j++ {
			element := i*elementsPerGoroutine + j
			if !valueMap[element] {
				t.Errorf("Element %d not found in final list", element)
			}
		}
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
	if !result {
		t.Errorf("Expected ContainsAll to return true for empty collection, got false")
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
	// Test removing from middle of list
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	result := list.Remove(3)
	if !result {
		t.Errorf("Expected Remove to return true for middle element")
	}
	if list.Size() != 4 {
		t.Errorf("Expected size to be 4, got %d", list.Size())
	}
	if list.Contains(3) {
		t.Errorf("Expected list not to contain element 3")
	}
	val, err := list.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 2 to be 4, got %d", *val)
	}

	// Test removing from head
	result = list.Remove(1)
	if !result {
		t.Errorf("Expected Remove to return true for head element")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}
	if list.Contains(1) {
		t.Errorf("Expected list not to contain element 1")
	}
	val, err = list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 0 to be 2, got %d", *val)
	}

	// Test removing from tail
	result = list.Remove(5)
	if !result {
		t.Errorf("Expected Remove to return true for tail element")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}
	if list.Contains(5) {
		t.Errorf("Expected list not to contain element 5")
	}
	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 4 {
		t.Errorf("Expected element at index 1 to be 4, got %d", *val)
	}

	// Test removing non-existent element
	result = list.Remove(6)
	if result {
		t.Errorf("Expected Remove to return false for non-existent element")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", list.Size())
	}

	// Test removing from empty list
	emptyList := LinkedList[int]{}
	result = emptyList.Remove(1)
	if result {
		t.Errorf("Expected Remove to return false for empty list")
	}
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0, got %d", emptyList.Size())
	}

	// Test removing last element (should set tail to nil)
	singleElementList := LinkedList[int]{}
	singleElementList.Add(1)
	result = singleElementList.Remove(1)
	if !result {
		t.Errorf("Expected Remove to return true for last element")
	}
	if singleElementList.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", singleElementList.Size())
	}
	if singleElementList.Contains(1) {
		t.Errorf("Expected list not to contain element 1")
	}
	// Verify list is empty by trying to get first element
	_, err = singleElementList.GetFirst()
	if err == nil {
		t.Errorf("Expected NoSuchElementError when getting first element from empty list")
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

	// Test with empty collection
	emptyElements := NewLinkedListWithInitialCollection[int]([]int{})
	result = list.AddAll(emptyElements)
	if result {
		t.Errorf("Expected AddAll to return false for empty collection since no elements were added, got true")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to remain 3 after adding empty collection, got %d", list.Size())
	}

	// Test with nil collection
	result = list.AddAll(nil)
	if result {
		t.Errorf("Expected AddAll to return false for nil collection, got true")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to remain 3 after adding nil collection, got %d", list.Size())
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

	elements := NewLinkedListWithInitialCollection([]int{4, 5, 6})

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

	// Test with empty collection
	result, err = list.AddAllAtIndex(0, NewLinkedListWithInitialCollection[int]([]int{}))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected AddAllAtIndex to return false for empty collection since no elements were added, got true")
	}
	if list.Size() != 6 {
		t.Errorf("Expected size to remain 6 after adding empty collection, got %d", list.Size())
	}

	// Test with nil collection
	result, err = list.AddAllAtIndex(0, nil)
	if err == nil {
		t.Errorf("Expected NullPointerError, got no error")
	}
	if result {
		t.Errorf("Expected AddAllAtIndex to return false for nil collection, got true")
	}
	if list.Size() != 6 {
		t.Errorf("Expected size to remain 6 after adding nil collection, got %d", list.Size())
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
	// Test removing last occurrence from head
	list := NewLinkedList[int]()
	list.AddAllBatch([]int{1, 2, 3, 1, 4, 1})
	if !list.RemoveLastOccurrence(1) {
		t.Error("RemoveLastOccurrence failed to remove last occurrence of 1")
	}
	if list.Size() != 5 {
		t.Errorf("Expected size 5 after removing last occurrence of 1, got %d", list.Size())
	}
	expected := []int{1, 2, 3, 1, 4}
	actual := list.ToArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v after removing last occurrence of 1, got %v", expected, actual)
	}

	// Test removing last occurrence from tail
	list = NewLinkedList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 3})
	if !list.RemoveLastOccurrence(3) {
		t.Error("RemoveLastOccurrence failed to remove last occurrence of 3")
	}
	if list.Size() != 4 {
		t.Errorf("Expected size 4 after removing last occurrence of 3, got %d", list.Size())
	}
	expected = []int{1, 2, 3, 4}
	actual = list.ToArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v after removing last occurrence of 3, got %v", expected, actual)
	}

	// Test removing last occurrence from middle
	list = NewLinkedList[int]()
	list.AddAllBatch([]int{1, 2, 3, 2, 4, 2})
	if !list.RemoveLastOccurrence(2) {
		t.Error("RemoveLastOccurrence failed to remove last occurrence of 2")
	}
	if list.Size() != 5 {
		t.Errorf("Expected size 5 after removing last occurrence of 2, got %d", list.Size())
	}
	expected = []int{1, 2, 3, 2, 4}
	actual = list.ToArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v after removing last occurrence of 2, got %v", expected, actual)
	}

	// Test removing non-existent element
	list = NewLinkedList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4})
	if list.RemoveLastOccurrence(5) {
		t.Error("RemoveLastOccurrence should return false for non-existent element")
	}
	if list.Size() != 4 {
		t.Errorf("Expected size 4 after attempting to remove non-existent element, got %d", list.Size())
	}

	// Test removing from empty list
	list = NewLinkedList[int]()
	if list.RemoveLastOccurrence(1) {
		t.Error("RemoveLastOccurrence should return false for empty list")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0 for empty list, got %d", list.Size())
	}

	// Test removing last element (to test l.tail = nil case)
	list = NewLinkedList[int]()
	list.Add(1)
	if !list.RemoveLastOccurrence(1) {
		t.Error("RemoveLastOccurrence failed to remove last element")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0 after removing last element, got %d", list.Size())
	}
	if list.head != nil || list.tail != nil {
		t.Error("Expected head and tail to be nil after removing last element")
	}

	// Test removing last occurrence in the middle (not head or tail)
	list = NewLinkedList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 2, 5})
	if !list.RemoveLastOccurrence(2) {
		t.Error("RemoveLastOccurrence failed to remove last occurrence of 2 in the middle")
	}
	if list.Size() != 5 {
		t.Errorf("Expected size 5 after removing last occurrence of 2 in the middle, got %d", list.Size())
	}
	expected = []int{1, 2, 3, 4, 5}
	actual = list.ToArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v after removing last occurrence of 2 in the middle, got %v", expected, actual)
	}

	// Test removing last occurrence at head (hits l.head.SetPrev(nil))
	list = NewLinkedList[int]()
	list.AddAllBatch([]int{2, 1, 3, 4, 5})
	if !list.RemoveLastOccurrence(2) {
		t.Error("RemoveLastOccurrence failed to remove last occurrence of 2 at head")
	}
	if list.Size() != 4 {
		t.Errorf("Expected size 4 after removing last occurrence of 2 at head, got %d", list.Size())
	}
	expected = []int{1, 3, 4, 5}
	actual = list.ToArray()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v after removing last occurrence of 2 at head, got %v", expected, actual)
	}
	// Check that new head's prev is nil
	if list.head.GetPrev() != nil {
		t.Error("Expected new head's prev to be nil after removing head")
	}
}

func TestLinkedList_RemoveHead(t *testing.T) {
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
		t.Errorf("Expected element at index 0 to be 2, got %d", *val)
	}

	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 1 to be 3, got %d", *val)
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
	// Test with non-empty list
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)

	element, err := list.Peek()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *element != 1 {
		t.Errorf("Expected element to be 1, got %d", *element)
	}

	// Test with empty list
	emptyList := LinkedList[int]{}
	element, err = emptyList.Peek()
	if err != nil {
		t.Errorf("Expected no error for empty list, got %v", err)
	}
	if element != nil {
		t.Errorf("Expected nil element for empty list, got %v", element)
	}
}
func TestLinkedList_Poll(t *testing.T) {
	// Test with non-empty list
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

	// Test with empty list
	emptyList := LinkedList[int]{}
	element, err = emptyList.Poll()
	if err != nil {
		t.Errorf("Expected no error for empty list, got %v", err)
	}
	if element != nil {
		t.Errorf("Expected nil element for empty list, got %v", element)
	}
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0, got %d", emptyList.Size())
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
	// Test with non-empty list
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

	// Test with empty list
	emptyList := LinkedList[int]{}
	element, err = emptyList.PeekFirst()
	if err == nil {
		t.Errorf("Expected NoSuchElementError for empty list, got no error")
	}
	if element != nil {
		t.Errorf("Expected nil element for empty list, got %v", element)
	}
}
func TestLinkedList_PeekLast(t *testing.T) {
	// Test with non-empty list
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

	// Test with empty list
	emptyList := LinkedList[int]{}
	element, err = emptyList.PeekLast()
	if err == nil {
		t.Errorf("Expected NoSuchElementError for empty list, got no error")
	}
	if element != nil {
		t.Errorf("Expected nil element for empty list, got %v", element)
	}
}

func TestLinkedList_PollFirst(t *testing.T) {
	// Test with non-empty list
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

	// Test with empty list
	emptyList := LinkedList[int]{}
	element, err = emptyList.PollFirst()
	if err != nil {
		t.Errorf("Expected no error for empty list, got %v", err)
	}
	if element != nil {
		t.Errorf("Expected nil element for empty list, got %v", element)
	}
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0, got %d", emptyList.Size())
	}
}

func TestLinkedList_PollLast(t *testing.T) {
	// Test with non-empty list
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

	// Test with empty list
	emptyList := LinkedList[int]{}
	element, err = emptyList.PollLast()
	if err != nil {
		t.Errorf("Expected no error for empty list, got %v", err)
	}
	if element != nil {
		t.Errorf("Expected nil element for empty list, got %v", element)
	}
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0, got %d", emptyList.Size())
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

	comparator := &IntComparator{}
	list.Sort(comparator)

	expected := []int{1, 2, 3}
	arr := list.ToArray()
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected sorted array to be %v, got %v", expected, arr)
	}
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

func TestLinkedList_CopyOf(t *testing.T) {
	// Test with non-empty collection
	original := NewLinkedListWithInitialCollection[int]([]int{1, 2, 3})
	copy := original.CopyOf(original)

	if copy.Size() != 3 {
		t.Errorf("Expected copy size to be 3, got %d", copy.Size())
	}

	val, err := copy.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}

	val, err = copy.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}

	val, err = copy.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", *val)
	}

	// Test with empty collection
	emptyList := NewLinkedList[int]()
	copy = emptyList.CopyOf(emptyList)
	if copy.Size() != 0 {
		t.Errorf("Expected copy size to be 0 for empty collection, got %d", copy.Size())
	}

	// Test with nil collection
	copy = emptyList.CopyOf(nil)
	if copy.Size() != 0 {
		t.Errorf("Expected copy size to be 0 for nil collection, got %d", copy.Size())
	}
}

func TestLinkedList_AddAllAtIndexBatch(t *testing.T) {
	// Test adding at index 0 when list is empty
	list := LinkedList[int]{}
	elements := []int{1, 2, 3}
	result, err := list.AddAllAtIndexBatch(0, elements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected AddAllAtIndexBatch to return true, got false")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}

	// Test adding at index 0 when list has elements
	elements = []int{4, 5}
	result, err = list.AddAllAtIndexBatch(0, elements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected AddAllAtIndexBatch to return true, got false")
	}
	if list.Size() != 5 {
		t.Errorf("Expected size to be 5, got %d", list.Size())
	}

	// Test adding at the end of the list
	elements = []int{6, 7}
	result, err = list.AddAllAtIndexBatch(5, elements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected AddAllAtIndexBatch to return true, got false")
	}
	if list.Size() != 7 {
		t.Errorf("Expected size to be 7, got %d", list.Size())
	}

	// Test adding in the middle of the list
	elements = []int{8, 9}
	result, err = list.AddAllAtIndexBatch(3, elements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected AddAllAtIndexBatch to return true, got false")
	}
	if list.Size() != 9 {
		t.Errorf("Expected size to be 9, got %d", list.Size())
	}

	// Test adding with invalid index
	result, err = list.AddAllAtIndexBatch(10, elements)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}
	if result {
		t.Errorf("Expected AddAllAtIndexBatch to return false, got true")
	}

	// Test adding with negative index
	result, err = list.AddAllAtIndexBatch(-1, elements)
	if err == nil {
		t.Errorf("Expected IndexOutOfBoundsError, got no error")
	}
	if result {
		t.Errorf("Expected AddAllAtIndexBatch to return false, got true")
	}

	// Test adding empty elements
	result, err = list.AddAllAtIndexBatch(0, []int{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected AddAllAtIndexBatch to return false for empty elements, got true")
	}
	if list.Size() != 9 {
		t.Errorf("Expected size to remain 9, got %d", list.Size())
	}
}

func TestLinkedList_RemoveAllBatch(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(2)
	list.Add(4)
	list.Add(2)

	// Test removing multiple elements
	elements := []int{2, 3}
	result := list.RemoveAllBatch(elements)
	if !result {
		t.Errorf("Expected RemoveAllBatch to return true, got false")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}
	if list.Contains(2) || list.Contains(3) {
		t.Errorf("Expected all occurrences of 2 and 3 to be removed")
	}

	// Test removing non-existent elements
	elements = []int{5, 6}
	result = list.RemoveAllBatch(elements)
	if result {
		t.Errorf("Expected RemoveAllBatch to return false for non-existent elements, got true")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", list.Size())
	}

	// Test with empty elements
	result = list.RemoveAllBatch([]int{})
	if result {
		t.Errorf("Expected RemoveAllBatch to return false for empty elements, got true")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", list.Size())
	}

	// Test with nil elements
	result = list.RemoveAllBatch(nil)
	if result {
		t.Errorf("Expected RemoveAllBatch to return false for nil elements, got true")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to remain 2 after removing nil elements, got %d", list.Size())
	}

	// Test removing from empty list
	emptyList := LinkedList[int]{}
	result = emptyList.RemoveAllBatch(elements)
	if result {
		t.Errorf("Expected RemoveAllBatch to return false for empty list, got true")
	}
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0, got %d", emptyList.Size())
	}

	// Test removing elements that appear multiple times
	list = LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(1)
	list.Add(3)
	list.Add(1)
	list.Add(4)
	list.Add(1)

	elements = []int{1, 3}
	result = list.RemoveAllBatch(elements)
	if !result {
		t.Errorf("Expected RemoveAllBatch to return true, got false")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}
	if list.Contains(1) || list.Contains(3) {
		t.Errorf("Expected all occurrences of 1 and 3 to be removed")
	}

	// Test removing elements at start and end
	list = LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(1)

	elements = []int{1}
	result = list.RemoveAllBatch(elements)
	if !result {
		t.Errorf("Expected RemoveAllBatch to return true, got false")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", list.Size())
	}
	if list.Contains(1) {
		t.Errorf("Expected all occurrences of 1 to be removed")
	}
	if list.IndexOf(2) != 0 {
		t.Errorf("Expected first element to be 2")
	}
	if list.LastIndexOf(4) != 2 {
		t.Errorf("Expected last element to be 4")
	}
}

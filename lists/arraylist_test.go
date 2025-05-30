package lists

import (
	"reflect"
	"sync"
	"testing"
	"time"

	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
	"github.com/stretchr/testify/assert"
)

// TestArrayList_Constructor tests different constructor methods
func TestArrayList_Constructor(t *testing.T) {
	// Test default constructor
	list := NewArrayList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test constructor with initial capacity
	list = NewArrayListWithInitialCapacity[int](20)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, 20, list.Capacity())

	// Test constructor with negative capacity
	list = NewArrayListWithInitialCapacity[int](-1)
	assert.NotNil(t, list)
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test constructor with initial collection
	initialValues := []int{1, 2, 3}
	list = NewArrayListWithInitialCollection[int](initialValues)
	assert.NotNil(t, list)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 3, list.Capacity())
}

// TestArrayList_CapacityManagement tests capacity-related operations
func TestArrayList_CapacityManagement(t *testing.T) {
	list := NewArrayList[int]()

	// Test initial capacity
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test capacity growth
	for i := 0; i < 20; i++ {
		list.Add(i)
	}
	assert.Greater(t, list.Capacity(), DefaultCapacity)

	// Test TrimToSize
	list.TrimToSize()
	assert.Equal(t, list.Size(), list.Capacity())

	// Test EnsureCapacity
	list.EnsureCapacity(50)
	assert.GreaterOrEqual(t, list.Capacity(), 50)
}

// TestArrayList_ConcurrentOperations tests concurrent operations
func TestArrayList_ConcurrentOperations(t *testing.T) {
	list := NewArrayList[int]()
	done := make(chan bool)

	// Test concurrent additions
	for i := 0; i < 1000; i++ {
		go func(val int) {
			list.Add(val)
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 1000; i++ {
		<-done
	}

	assert.Equal(t, 1000, list.Size())

	// Test concurrent reads and writes
	for i := 0; i < 1000; i++ {
		go func() {
			list.Add(1)
			list.Size()
			done <- true
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}
}

// TestArrayList_ErrorHandling tests error conditions
func TestArrayList_ErrorHandling(t *testing.T) {
	list := NewArrayList[int]()

	// Test index out of bounds
	_, err := list.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	_, err = list.Get(0)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test invalid index in AddAtIndex
	err = list.AddAtIndex(-1, 1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test nil collection in AddAll
	_, err = list.AddAllAtIndex(0, nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
}

// TestArrayList_FunctionalOperations tests functional operations
func TestArrayList_FunctionalOperations(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}

	// Test FindFirst
	val, err := list.FindFirst(func(x int) bool { return x > 5 })
	assert.NoError(t, err)
	assert.Equal(t, 6, *val)

	// Test FindAll
	vals := list.FindAll(func(x int) bool { return x%2 == 0 })
	assert.Equal(t, 5, len(vals))

	// Test Filter
	filtered := list.Filter(func(x int) bool { return x > 5 })
	assert.Equal(t, 5, filtered.Size())

	// Test ForEach
	sum := 0
	list.ForEach(func(x int) { sum += x })
	assert.Equal(t, 55, sum)
}

// TestArrayList_BatchOperations tests batch operations
func TestArrayList_BatchOperations(t *testing.T) {
	list := NewArrayList[int]()
	otherList := NewArrayList[int]()

	// Test AddAll
	for i := 1; i <= 5; i++ {
		otherList.Add(i)
	}
	result := list.AddAll(otherList)
	assert.True(t, result)
	assert.Equal(t, 5, list.Size())

	// Test RemoveAll
	toRemove := NewArrayList[int]()
	toRemove.Add(2)
	toRemove.Add(4)
	result = list.RemoveAll(toRemove)
	assert.True(t, result)
	assert.Equal(t, 3, list.Size())

	// Test RetainAll
	toRetain := NewArrayList[int]()
	toRetain.Add(1)
	toRetain.Add(3)
	result, err := list.RetainAll(toRetain)
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 2, list.Size())
}

// TestArrayList_Iterator tests iterator operations
func TestArrayList_Iterator(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	// Test normal iteration
	iterator := list.Iterator()
	count := 0
	for iterator.HasNext() {
		val, err := iterator.Next()
		assert.NoError(t, err)
		assert.Equal(t, count+1, *val)
		count++
	}
	assert.Equal(t, 5, count)

	// Test iterator with empty list
	emptyList := NewArrayList[int]()
	emptyIterator := emptyList.Iterator()
	assert.False(t, emptyIterator.HasNext())
	_, err := emptyIterator.Next()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())
}

// TestArrayList_SortAndShuffle tests sorting and shuffling
func TestArrayList_SortAndShuffle(t *testing.T) {
	list := NewArrayList[int]()
	for i := 5; i > 0; i-- {
		list.Add(i)
	}

	// Test Sort
	comparator := &IntComparator{}
	list.Sort(comparator)
	for i := 0; i < 5; i++ {
		val, _ := list.Get(i)
		assert.Equal(t, i+1, *val)
	}

	// Test Shuffle
	original := make([]int, list.Size())
	copy(original, list.ToArray())
	list.Shuffle()
	// Note: There's a very small chance this could fail if shuffle returns original order
	assert.NotEqual(t, original, list.ToArray())
}

// TestArrayList_SubList tests sublist operations
func TestArrayList_SubList(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}

	// Test valid sublist
	sublist, err := list.SubList(2, 5)
	assert.NoError(t, err)
	assert.Equal(t, 3, sublist.Size())
	val, _ := sublist.Get(0)
	assert.Equal(t, 3, *val)

	// Test invalid indices
	_, err = list.SubList(-1, 5)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
	_, err = list.SubList(2, 11)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
	_, err = list.SubList(5, 2)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IllegalArgumentError), err.Error())
}

// TestArrayList_Performance tests performance with large collections
func TestArrayList_Performance(t *testing.T) {
	list := NewArrayList[int]()
	start := time.Now()

	// Test adding large number of elements
	for i := 0; i < 100000; i++ {
		list.Add(i)
	}
	addTime := time.Since(start)
	assert.Less(t, addTime, time.Second)

	// Test searching in large collection
	start = time.Now()
	for i := 0; i < 1000; i++ {
		list.Contains(i)
	}
	searchTime := time.Since(start)
	assert.Less(t, searchTime, time.Second)
}

// TestArrayList_AddFirstAndLast tests AddFirst and AddLast operations
func TestArrayList_AddFirstAndLast(t *testing.T) {
	list := NewArrayList[int]()

	// Test AddFirst
	list.AddFirst(1)
	list.AddFirst(2)
	assert.Equal(t, 2, list.Size())
	val, _ := list.Get(0)
	assert.Equal(t, 2, *val)

	// Test AddLast
	list.AddLast(3)
	list.AddLast(4)
	assert.Equal(t, 4, list.Size())
	val, _ = list.Get(3)
	assert.Equal(t, 4, *val)
}

// TestArrayList_AddAllFirstAndLast tests AddAllFirst and AddAllLast operations
func TestArrayList_AddAllFirstAndLast(t *testing.T) {
	list := NewArrayList[int]()
	otherList := NewArrayList[int]()
	for i := 1; i <= 3; i++ {
		otherList.Add(i)
	}

	// Test AddAllFirst
	result, err := list.AddAllFirst(otherList)
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 3, list.Size())
	val, _ := list.Get(0)
	assert.Equal(t, 1, *val)

	// Test AddAllLast
	result, err = list.AddAllLast(otherList)
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 6, list.Size())
	val, _ = list.Get(5)
	assert.Equal(t, 3, *val)

	// Test with nil collection
	result, err = list.AddAllFirst(nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
	assert.False(t, result)
}

// TestArrayList_AddAllFirst tests AddAllFirst operation
func TestArrayList_AddAllFirst(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	elements := NewArrayListWithInitialCollection[int]([]int{1, 2, 3})
	result, err := emptyList.AddAllFirst(elements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected AddAllFirst to return true, got false")
	}
	if emptyList.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", emptyList.Size())
	}
	val, err := emptyList.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected element at index 0 to be 1, got %d", *val)
	}
	val, err = emptyList.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %d", *val)
	}
	val, err = emptyList.Get(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected element at index 2 to be 3, got %d", *val)
	}

	// Test with nil collection
	result, err = emptyList.AddAllFirst(nil)
	if err == nil {
		t.Errorf("Expected NullPointerError, got no error")
	}
	if result {
		t.Errorf("Expected AddAllFirst to return false for nil collection, got true")
	}
	if emptyList.Size() != 3 {
		t.Errorf("Expected size to remain 3 after adding nil collection, got %d", emptyList.Size())
	}

	// Test with empty collection
	emptyElements := NewArrayListWithInitialCollection[int]([]int{})
	result, err = emptyList.AddAllFirst(emptyElements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected AddAllFirst to return false for empty collection, got true")
	}
	if emptyList.Size() != 3 {
		t.Errorf("Expected size to remain 3 after adding empty collection, got %d", emptyList.Size())
	}
}

// TestArrayList_RemoveIf tests RemoveIf operation
func TestArrayList_RemoveIf(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}

	// Test removing even numbers
	result := list.RemoveIf(func(x int) bool { return x%2 == 0 })
	assert.True(t, result)
	assert.Equal(t, 5, list.Size())
	for i := 0; i < list.Size(); i++ {
		val, _ := list.Get(i)
		assert.True(t, *val%2 != 0)
	}

	// Test removing non-existent elements
	result = list.RemoveIf(func(x int) bool { return x > 10 })
	assert.False(t, result)
}

// TestArrayList_ReplaceAll tests ReplaceAll operation
func TestArrayList_ReplaceAll(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	list.ReplaceAll(func(x int) int { return x * 2 })
	assert.Equal(t, 5, list.Size())
	for i := 0; i < list.Size(); i++ {
		val, _ := list.Get(i)
		assert.Equal(t, (i+1)*2, *val)
	}
}

// TestArrayList_RemoveRange tests RemoveRange operation
func TestArrayList_RemoveRange(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}

	// Test valid range removal
	err := list.RemoveRange(2, 5)
	assert.NoError(t, err)
	assert.Equal(t, 7, list.Size())
	val, _ := list.Get(2)
	assert.Equal(t, 6, *val)

	// Test invalid range
	err = list.RemoveRange(-1, 5)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
	err = list.RemoveRange(5, 2)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
	err = list.RemoveRange(2, 20)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
}

// TestArrayList_FastBatchOperations tests fast batch operations
func TestArrayList_FastBatchOperations(t *testing.T) {
	list := NewArrayList[int]()
	elements := []int{1, 2, 3, 4, 5}

	// Test AddAllBatch
	result := list.AddAllBatch(elements)
	assert.True(t, result)
	assert.Equal(t, 5, list.Size())

	// Test FastContains
	assert.True(t, list.FastContains(3))
	assert.False(t, list.FastContains(6))

	// Test FastRemoveAll
	toRemove := []int{2, 4}
	result = list.FastRemoveAll(toRemove)
	assert.True(t, result)
	assert.Equal(t, 3, list.Size())
	assert.False(t, list.Contains(2))
	assert.False(t, list.Contains(4))

	// Test FastRetainAll
	toRetain := []int{1, 3}
	result, err := list.FastRetainAll(toRetain)
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 2, list.Size())
	assert.True(t, list.Contains(1))
	assert.True(t, list.Contains(3))
}

// TestArrayList_FastOperations tests fast operations with sorted lists
func TestArrayList_FastOperations(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}

	// Test FastIndexOf with sorted list
	comparator := &IntComparator{}
	index := list.FastIndexOf(5, true, comparator)
	assert.Equal(t, 4, index)

	// Test FastLastIndexOf
	index = list.FastLastIndexOf(5)
	assert.Equal(t, 4, index)

	// Test FastSubList
	sublist, err := list.FastSubList(2, 5)
	assert.NoError(t, err)
	assert.Equal(t, 3, sublist.Size())
	val, _ := sublist.Get(0)
	assert.Equal(t, 3, *val)
}

// TestArrayList_Clone tests Clone operation
func TestArrayList_Clone(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	clone := list.Clone()
	assert.Equal(t, list.Size(), clone.Size())
	for i := 0; i < list.Size(); i++ {
		val1, _ := list.Get(i)
		val2, _ := clone.Get(i)
		assert.Equal(t, *val1, *val2)
	}

	// Test that clone is independent
	clone.Add(6)
	assert.Equal(t, 5, list.Size())
	assert.Equal(t, 6, clone.Size())
}

// TestArrayList_CopyOf tests CopyOf operation
func TestArrayList_CopyOf(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	// Test with valid collection
	copy := list.CopyOf(list)
	assert.Equal(t, list.Size(), copy.Size())
	for i := 0; i < list.Size(); i++ {
		val1, _ := list.Get(i)
		val2, _ := copy.Get(i)
		assert.Equal(t, *val1, *val2)
	}

	// Test with nil collection
	copy = list.CopyOf(nil)
	assert.Nil(t, copy)

	// Test with empty collection
	emptyList := NewArrayList[int]()
	copy = list.CopyOf(emptyList)
	assert.Equal(t, 0, copy.Size())
}

// TestArrayList_ConcurrentModification tests concurrent modification scenarios
func TestArrayList_ConcurrentModification(t *testing.T) {
	list := NewArrayList[int]()
	done := make(chan bool)

	// Test concurrent modifications during iteration
	for i := 0; i < 1000; i++ {
		go func() {
			list.Add(1)
			done <- true
		}()
	}

	// Concurrent iteration
	go func() {
		iterator := list.Iterator()
		for iterator.HasNext() {
			_, _ = iterator.Next()
		}
		done <- true
	}()

	// Wait for all goroutines
	for i := 0; i < 1001; i++ {
		<-done
	}
}

// TestArrayList_EdgeCases tests various edge cases
func TestArrayList_EdgeCases(t *testing.T) {
	list := NewArrayList[int]()

	// Test with MaxArraySize
	list.EnsureCapacity(MaxArraySize - 1)
	assert.Equal(t, MaxArraySize-1, list.Capacity())

	// Test with negative capacity
	list = NewArrayListWithInitialCapacity[int](-1)
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test with zero capacity
	list = NewArrayListWithInitialCapacity[int](0)
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test with nil collection
	list = NewArrayListWithInitialCollection[int](nil)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, DefaultCapacity, list.Capacity())
}

// TestArrayList_IteratorConcurrentModification tests iterator behavior with concurrent modifications
func TestArrayList_IteratorConcurrentModification(t *testing.T) {
	list := NewArrayList[int]()
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	iterator := list.Iterator()
	val, err := iterator.Next()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)

	// Modify list while iterating
	list.Add(6)
	val, err = iterator.Next()
	assert.NoError(t, err)
	assert.Equal(t, 2, *val)
}

// TestArrayList_Shuffle tests shuffle operation with different sizes
func TestArrayList_Shuffle(t *testing.T) {
	// Test with empty list
	list := NewArrayList[int]()
	list.Shuffle()
	assert.Equal(t, 0, list.Size())

	// Test with single element
	list.Add(1)
	list.Shuffle()
	assert.Equal(t, 1, list.Size())
	val, _ := list.Get(0)
	assert.Equal(t, 1, *val)

	// Test with multiple elements
	list = NewArrayList[int]()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}
	original := make([]int, list.Size())
	copy(original, list.ToArray())
	list.Shuffle()
	// Note: There's a very small chance this could fail if shuffle returns original order
	assert.NotEqual(t, original, list.ToArray())
}

// TestArrayList_Remove tests Remove operation
func TestArrayList_Remove(t *testing.T) {
	// Test removing from empty list
	emptyList := NewArrayList[int]()
	if emptyList.Remove(1) {
		t.Errorf("Expected Remove to return false for empty list")
	}
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0 after removing from empty list")
	}

	// Test removing non-existent element
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.Remove(4) {
		t.Errorf("Expected Remove to return false for non-existent element")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to remain 3 after removing non-existent element")
	}

	// Test removing first element
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if !list.Remove(1) {
		t.Errorf("Expected Remove to return true for existing element")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2 after removing first element")
	}
	if list.Contains(1) {
		t.Errorf("Expected list not to contain removed element")
	}
	val, err := list.Get(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected first element to be 2 after removal, got %d", *val)
	}

	// Test removing middle element
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if !list.Remove(2) {
		t.Errorf("Expected Remove to return true for existing element")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2 after removing middle element")
	}
	if list.Contains(2) {
		t.Errorf("Expected list not to contain removed element")
	}
	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected second element to be 3 after removal, got %d", *val)
	}

	// Test removing last element
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if !list.Remove(3) {
		t.Errorf("Expected Remove to return true for existing element")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2 after removing last element")
	}
	if list.Contains(3) {
		t.Errorf("Expected list not to contain removed element")
	}
	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected second element to be 2 after removal, got %d", *val)
	}

	// Test removing duplicate elements (should remove first occurrence)
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(2)
	list.Add(3)
	if !list.Remove(2) {
		t.Errorf("Expected Remove to return true for existing element")
	}
	if list.Size() != 3 {
		t.Errorf("Expected size to be 3 after removing duplicate element")
	}
	val, err = list.Get(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected second element to still be 2 (second occurrence), got %d", *val)
	}

	// Test removing all elements
	list = NewArrayList[int]()
	list.Add(1)
	if !list.Remove(1) {
		t.Errorf("Expected Remove to return true for existing element")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size to be 0 after removing all elements")
	}
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty after removing all elements")
	}
}

// TestArrayList_RemoveFirst tests RemoveFirst operation
func TestArrayList_RemoveFirst(t *testing.T) {
	// Test removing from empty list
	emptyList := NewArrayList[int]()
	_, err := emptyList.RemoveFirst()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())
	assert.Equal(t, 0, emptyList.Size())

	// Test removing from single element list
	list := NewArrayList[int]()
	list.Add(1)
	val, err := list.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)
	assert.Equal(t, 0, list.Size())
	assert.True(t, list.IsEmpty())

	// Test removing from multi-element list
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	val, err = list.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)
	assert.Equal(t, 2, list.Size())
	firstVal, _ := list.Get(0)
	assert.Equal(t, 2, *firstVal)
	secondVal, _ := list.Get(1)
	assert.Equal(t, 3, *secondVal)

	// Test removing all elements
	val, err = list.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 2, *val)
	val, err = list.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)
	assert.Equal(t, 0, list.Size())
	assert.True(t, list.IsEmpty())

	// Test removing after clearing
	list.Clear()
	_, err = list.RemoveFirst()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())
}

// TestArrayList_RemoveLast tests RemoveLast operation
func TestArrayList_RemoveLast(t *testing.T) {
	// Test removing from empty list
	emptyList := NewArrayList[int]()
	_, err := emptyList.RemoveLast()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())
	assert.Equal(t, 0, emptyList.Size())

	// Test removing from single element list
	list := NewArrayList[int]()
	list.Add(1)
	val, err := list.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)
	assert.Equal(t, 0, list.Size())
	assert.True(t, list.IsEmpty())

	// Test removing from multi-element list
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	val, err = list.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)
	assert.Equal(t, 2, list.Size())
	firstVal, _ := list.Get(0)
	assert.Equal(t, 1, *firstVal)
	secondVal, _ := list.Get(1)
	assert.Equal(t, 2, *secondVal)

	// Test removing all elements
	val, err = list.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 2, *val)
	val, err = list.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)
	assert.Equal(t, 0, list.Size())
	assert.True(t, list.IsEmpty())

	// Test removing after clearing
	list.Clear()
	_, err = list.RemoveLast()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())
}

// TestArrayList_Set tests Set operation
func TestArrayList_Set(t *testing.T) {
	// Test setting in empty list
	emptyList := NewArrayList[int]()
	_, err := emptyList.Set(0, 1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test setting with negative index
	list := NewArrayList[int]()
	list.Add(1)
	_, err = list.Set(-1, 2)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test setting with index out of bounds
	_, err = list.Set(1, 2)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test setting first element
	oldVal, err := list.Set(0, 2)
	assert.NoError(t, err)
	assert.Equal(t, 1, *oldVal)
	val, _ := list.Get(0)
	assert.Equal(t, 2, *val)

	// Test setting middle element
	list.Add(3)
	list.Add(4)
	oldVal, err = list.Set(1, 5)
	assert.NoError(t, err)
	assert.Equal(t, 3, *oldVal)
	val, _ = list.Get(1)
	assert.Equal(t, 5, *val)

	// Test setting last element
	oldVal, err = list.Set(2, 6)
	assert.NoError(t, err)
	assert.Equal(t, 4, *oldVal)
	val, _ = list.Get(2)
	assert.Equal(t, 6, *val)

	// Test setting with same value
	oldVal, err = list.Set(0, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, *oldVal)
	val, _ = list.Get(0)
	assert.Equal(t, 2, *val)

	// Test setting after clearing
	list.Clear()
	_, err = list.Set(0, 1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
}

// TestArrayList_ContainsAll tests ContainsAll operation
func TestArrayList_ContainsAll(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test with a collection that contains all elements
	collection := NewArrayList[int]()
	collection.Add(1)
	collection.Add(2)
	collection.Add(3)
	result, err := list.ContainsAll(collection)
	assert.NoError(t, err)
	assert.True(t, result)

	// Test with a collection that does not contain all elements
	collection = NewArrayList[int]()
	collection.Add(1)
	collection.Add(2)
	collection.Add(4)
	result, err = list.ContainsAll(collection)
	assert.NoError(t, err)
	assert.False(t, result)

	// Test with an empty collection
	collection = NewArrayList[int]()
	result, err = list.ContainsAll(collection)
	assert.NoError(t, err)
	assert.True(t, result)

	// Test with nil collection
	result, err = list.ContainsAll(nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
	assert.False(t, result)
}

// TestArrayList_Equals tests Equals operation
func TestArrayList_Equals(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test equal collection
	equalList := NewArrayList[int]()
	equalList.Add(1)
	equalList.Add(2)
	equalList.Add(3)
	assert.True(t, list.Equals(equalList))

	// Test non-equal collection
	nonEqualList := NewArrayList[int]()
	nonEqualList.Add(1)
	nonEqualList.Add(2)
	nonEqualList.Add(4)
	assert.False(t, list.Equals(nonEqualList))

	// Test different size collection
	smallerList := NewArrayList[int]()
	smallerList.Add(1)
	smallerList.Add(2)
	assert.False(t, list.Equals(smallerList))

	// Test empty collection
	emptyList := NewArrayList[int]()
	assert.False(t, list.Equals(emptyList))

	// Test nil collection
	assert.False(t, list.Equals(nil))

	// Test empty list equals empty list
	emptyList2 := NewArrayList[int]()
	assert.True(t, emptyList.Equals(emptyList2))
}

// TestArrayList_GetFirst tests GetFirst operation
func TestArrayList_GetFirst(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	_, err := emptyList.GetFirst()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())

	// Test with single element
	list := NewArrayList[int]()
	list.Add(1)
	val, err := list.GetFirst()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)

	// Test with multiple elements
	list.Add(2)
	list.Add(3)
	val, err = list.GetFirst()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)

	// Test after clearing
	list.Clear()
	_, err = list.GetFirst()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())
}

// TestArrayList_GetLast tests GetLast operation
func TestArrayList_GetLast(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	_, err := emptyList.GetLast()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())

	// Test with single element
	list := NewArrayList[int]()
	list.Add(1)
	val, err := list.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)

	// Test with multiple elements
	list.Add(2)
	list.Add(3)
	val, err = list.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)

	// Test after clearing
	list.Clear()
	_, err = list.GetLast()
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NoSuchElementError), err.Error())
}

// TestArrayList_IndexOf tests IndexOf operation
func TestArrayList_IndexOf(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(2)
	list.Add(4)

	// Test finding first occurrence
	index := list.IndexOf(2)
	assert.Equal(t, 1, index)

	// Test finding non-existent element
	index = list.IndexOf(5)
	assert.Equal(t, -1, index)

	// Test with empty list
	emptyList := NewArrayList[int]()
	index = emptyList.IndexOf(1)
	assert.Equal(t, -1, index)

	// Test after clearing
	list.Clear()
	index = list.IndexOf(1)
	assert.Equal(t, -1, index)
}

// TestArrayList_Reversed tests Reversed operation
func TestArrayList_Reversed(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	emptyList.Reversed()
	assert.Equal(t, 0, emptyList.Size())

	// Test with single element
	list := NewArrayList[int]()
	list.Add(1)
	list.Reversed()
	val, _ := list.Get(0)
	assert.Equal(t, 1, *val)

	// Test with multiple elements
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Reversed()
	val, _ = list.Get(0)
	assert.Equal(t, 3, *val)
	val, _ = list.Get(1)
	assert.Equal(t, 2, *val)
	val, _ = list.Get(2)
	assert.Equal(t, 1, *val)

	// Test with even number of elements
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Reversed()
	val, _ = list.Get(0)
	assert.Equal(t, 4, *val)
	val, _ = list.Get(1)
	assert.Equal(t, 3, *val)
	val, _ = list.Get(2)
	assert.Equal(t, 2, *val)
	val, _ = list.Get(3)
	assert.Equal(t, 1, *val)
}

// TestArrayList_RetainAll tests RetainAll operation
func TestArrayList_RetainAll(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Case 1: collection is nil
	result, err := list.RetainAll(nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // list should remain unchanged

	// Case 2: collection is empty
	emptyList := NewArrayList[int]()
	result, err = list.RetainAll(emptyList)
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 0, list.Size()) // list should be cleared
}

// TestArrayList_AddAll_EmptyOrNil tests AddAll with empty and nil collections
func TestArrayList_AddAll_EmptyOrNil(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Empty collection
	emptyList := NewArrayList[int]()
	result := list.AddAll(emptyList)
	assert.False(t, result)
	assert.Equal(t, 3, list.Size())

	// Nil collection
	result = list.AddAll(nil)
	assert.False(t, result)
	assert.Equal(t, 3, list.Size())
}

// TestArrayList_AddAllAtIndex_ErrorCases tests error scenarios for AddAllAtIndex
func TestArrayList_AddAllAtIndex_ErrorCases(t *testing.T) {
	// Test with negative index
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	collection := NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	_, err := list.AddAllAtIndex(-1, collection)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with index greater than size
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	collection = NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	_, err = list.AddAllAtIndex(4, collection)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with nil collection
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	_, err = list.AddAllAtIndex(1, nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())

	// Test with empty list
	emptyList := NewArrayList[int]()
	collection = NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	_, err = emptyList.AddAllAtIndex(1, collection)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with empty list and index 0
	emptyList = NewArrayList[int]()
	collection = NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	_, err = emptyList.AddAllAtIndex(0, collection)
	assert.NoError(t, err) // This should succeed as index 0 is valid for empty list

	// Test with empty collection
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	emptyCollection := NewArrayList[int]()
	var result bool
	result, err = list.AddAllAtIndex(1, emptyCollection)
	assert.NoError(t, err)
	assert.False(t, result) // Should return false as no elements were added

	// Test after clearing the list
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Clear()
	collection = NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	_, err = list.AddAllAtIndex(1, collection)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with single element list and index 1
	singleList := NewArrayList[int]()
	singleList.Add(1)
	collection = NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	_, err = singleList.AddAllAtIndex(1, collection)
	assert.NoError(t, err) // This should succeed as index 1 is valid for single element list

	// Test with single element list and index 2
	singleList = NewArrayList[int]()
	singleList.Add(1)
	collection = NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	_, err = singleList.AddAllAtIndex(2, collection)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
}

// TestArrayList_CalculateNewCapacity tests the calculateNewCapacity function
func TestArrayList_CalculateNewCapacity(t *testing.T) {
	// Test with zero or negative current capacity
	assert.Equal(t, MinCapacity, calculateNewCapacity(0, 5))
	assert.Equal(t, MinCapacity, calculateNewCapacity(-1, 5))

	// Test with small current capacity
	assert.Equal(t, 15, calculateNewCapacity(10, 12)) // 10 * 1.5 = 15

	// Test with required size larger than growth factor
	assert.Equal(t, 20, calculateNewCapacity(10, 20)) // Should use required size

	// Test with MaxArraySize
	assert.Equal(t, MaxArraySize, calculateNewCapacity(MaxArraySize-1, MaxArraySize))
	assert.Equal(t, MaxArraySize, calculateNewCapacity(MaxArraySize, MaxArraySize+1))

	// Test with growth factor calculation
	assert.Equal(t, 15, calculateNewCapacity(10, 11)) // 10 * 1.5 = 15
	assert.Equal(t, 22, calculateNewCapacity(15, 16)) // 15 * 1.5 = 22.5, rounded to 22

	// Test with small required size
	assert.Equal(t, 15, calculateNewCapacity(10, 5)) // Should still grow by factor

	// Test with very large numbers
	assert.Equal(t, MaxArraySize, calculateNewCapacity(MaxArraySize/2, MaxArraySize))
	assert.Equal(t, MaxArraySize, calculateNewCapacity(MaxArraySize-1, MaxArraySize))

	// Test with minimum capacity
	assert.Equal(t, MinCapacity, calculateNewCapacity(MinCapacity-1, MinCapacity-1))
	assert.Equal(t, int(float64(MinCapacity)*GrowthFactor), calculateNewCapacity(MinCapacity, MinCapacity))
}

// TestArrayList_NewArrayListWithInitialCapacity tests the NewArrayListWithInitialCapacity function
func TestArrayList_NewArrayListWithInitialCapacity(t *testing.T) {
	// Test with positive capacity
	list := NewArrayListWithInitialCapacity[int](20)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, 20, list.Capacity())

	// Test with zero capacity (should use DefaultCapacity)
	list = NewArrayListWithInitialCapacity[int](0)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test with negative capacity (should use DefaultCapacity)
	list = NewArrayListWithInitialCapacity[int](-1)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test with capacity equal to MaxArraySize
	list = NewArrayListWithInitialCapacity[int](MaxArraySize)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, MaxArraySize, list.Capacity())

	// Test with capacity greater than MaxArraySize (should be capped at MaxArraySize)
	list = NewArrayListWithInitialCapacity[int](MaxArraySize + 1)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, MaxArraySize, list.Capacity())

	// Test with capacity equal to MinCapacity
	list = NewArrayListWithInitialCapacity[int](MinCapacity)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, MinCapacity, list.Capacity())

	// Test with capacity less than MinCapacity (should use DefaultCapacity)
	list = NewArrayListWithInitialCapacity[int](MinCapacity - 1)
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, DefaultCapacity, list.Capacity())

	// Test with different types
	stringList := NewArrayListWithInitialCapacity[string](10)
	assert.NotNil(t, stringList)
	assert.Equal(t, 0, stringList.Size())
	assert.Equal(t, 10, stringList.Capacity())

	// Test that the list is empty and can be used
	list = NewArrayListWithInitialCapacity[int](10)
	list.Add(1)
	list.Add(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 10, list.Capacity())
	val, err := list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)
	val, err = list.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, *val)
}

// TestArrayList_RemoveAtIndex_EmptyOrNil tests RemoveAtIndex with empty list
func TestArrayList_RemoveAtIndex_EmptyOrNil(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	_, err := emptyList.RemoveAtIndex(0)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with negative index
	_, err = emptyList.RemoveAtIndex(-1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with index equal to size
	_, err = emptyList.RemoveAtIndex(emptyList.Size())
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with index greater than size
	_, err = emptyList.RemoveAtIndex(emptyList.Size() + 1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test after clearing the list
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Clear()
	_, err = list.RemoveAtIndex(0)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
}

// TestArrayList_RemoveAll_EmptyOrNil tests RemoveAll with empty or nil collections
func TestArrayList_RemoveAll_EmptyOrNil(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	emptyCollection := NewArrayList[int]()
	result := emptyList.RemoveAll(emptyCollection)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with nil collection
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	result = list.RemoveAll(nil)
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // List should remain unchanged

	// Test with empty collection
	result = list.RemoveAll(emptyCollection)
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // List should remain unchanged

	// Test with empty list and non-empty collection
	emptyList = NewArrayList[int]()
	collection := NewArrayList[int]()
	collection.Add(1)
	collection.Add(2)
	result = emptyList.RemoveAll(collection)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test after clearing the list
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Clear()
	result = list.RemoveAll(collection)
	assert.False(t, result)
	assert.Equal(t, 0, list.Size())

	// Test with collection containing elements not in the list
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	collection = NewArrayList[int]()
	collection.Add(4)
	collection.Add(5)
	result = list.RemoveAll(collection)
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // List should remain unchanged
}

// TestArrayList_AddAllBatch_EmptyOrNil tests AddAllBatch with empty list and nil/empty elements
func TestArrayList_AddAllBatch_EmptyOrNil(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	elements := []int{1, 2, 3}
	result := emptyList.AddAllBatch(elements)
	assert.True(t, result)
	assert.Equal(t, 3, emptyList.Size())
	val, err := emptyList.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)
	val, err = emptyList.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, *val)
	val, err = emptyList.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)

	// Test with nil elements
	result = emptyList.AddAllBatch(nil)
	assert.False(t, result)
	assert.Equal(t, 3, emptyList.Size()) // Size should remain unchanged

	// Test with empty elements
	result = emptyList.AddAllBatch([]int{})
	assert.False(t, result)
	assert.Equal(t, 3, emptyList.Size()) // Size should remain unchanged

	// Test with empty list and empty elements
	emptyList = NewArrayList[int]()
	result = emptyList.AddAllBatch([]int{})
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with empty list and nil elements
	emptyList = NewArrayList[int]()
	result = emptyList.AddAllBatch(nil)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())
}

// TestArrayList_AddAllLast_EmptyOrNil tests AddAllLast with empty list and nil/empty collections
func TestArrayList_AddAllLast_EmptyOrNil(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	elements := NewArrayListWithInitialCollection[int]([]int{1, 2, 3})
	result, err := emptyList.AddAllLast(elements)
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 3, emptyList.Size())
	val, err := emptyList.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)
	val, err = emptyList.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, *val)
	val, err = emptyList.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)

	// Test with nil collection
	result, err = emptyList.AddAllLast(nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
	assert.False(t, result)
	assert.Equal(t, 3, emptyList.Size()) // Size should remain unchanged

	// Test with empty collection
	emptyElements := NewArrayListWithInitialCollection[int]([]int{})
	result, err = emptyList.AddAllLast(emptyElements)
	assert.NoError(t, err)
	assert.False(t, result)
	assert.Equal(t, 3, emptyList.Size()) // Size should remain unchanged

	// Test with empty list and empty collection
	emptyList = NewArrayList[int]()
	result, err = emptyList.AddAllLast(emptyElements)
	assert.NoError(t, err)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with empty list and nil collection
	emptyList = NewArrayList[int]()
	result, err = emptyList.AddAllLast(nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())
}

// TestArrayList_FastRemoveAll_EmptyOrNil tests FastRemoveAll with empty list and nil/empty elements
func TestArrayList_FastRemoveAll_EmptyOrNil(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	elements := []int{1, 2, 3}
	result := emptyList.FastRemoveAll(elements)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with nil elements
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	result = list.FastRemoveAll(nil)
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // Size should remain unchanged

	// Test with empty elements
	result = list.FastRemoveAll([]int{})
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // Size should remain unchanged

	// Test with empty list and empty elements
	emptyList = NewArrayList[int]()
	result = emptyList.FastRemoveAll([]int{})
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with empty list and nil elements
	emptyList = NewArrayList[int]()
	result = emptyList.FastRemoveAll(nil)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with elements not in the list
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	result = list.FastRemoveAll([]int{4, 5, 6})
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // Size should remain unchanged
}

// TestArrayList_FastRetainAll_EmptyOrNil tests FastRetainAll with empty list and nil/empty elements
func TestArrayList_FastRetainAll_EmptyOrNil(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	elements := []int{1, 2, 3}
	result, err := emptyList.FastRetainAll(elements)
	assert.NoError(t, err)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with nil elements
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	result, err = list.FastRetainAll(nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
	assert.False(t, result)
	assert.Equal(t, 3, list.Size()) // Size should remain unchanged

	// Test with empty elements
	result, err = list.FastRetainAll([]int{})
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 0, list.Size()) // All elements should be removed

	// Test with empty list and empty elements
	emptyList = NewArrayList[int]()
	result, err = emptyList.FastRetainAll([]int{})
	assert.NoError(t, err)
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with empty list and nil elements
	emptyList = NewArrayList[int]()
	result, err = emptyList.FastRetainAll(nil)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.NullPointerError), err.Error())
	assert.False(t, result)
	assert.Equal(t, 0, emptyList.Size())

	// Test with elements not in the list
	list = NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	result, err = list.FastRetainAll([]int{4, 5})
	assert.NoError(t, err)
	assert.True(t, result)
	assert.Equal(t, 0, list.Size()) // All elements should be removed
}

// TestArrayList_FastRetainAll tests FastRetainAll operation
func TestArrayList_FastRetainAll(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 5})

	// Test retaining multiple elements
	elements := []int{2, 4}
	result, err := list.FastRetainAll(elements)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected FastRetainAll to return true, got false")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}
	if !list.Contains(2) || !list.Contains(4) {
		t.Errorf("Expected list to contain 2 and 4")
	}

	// Test with nil elements
	result, err = list.FastRetainAll(nil)
	if err == nil {
		t.Errorf("Expected NullPointerError, got no error")
	}
	if result {
		t.Errorf("Expected FastRetainAll to return false for nil elements, got true")
	}
	if list.Size() != 2 {
		t.Errorf("Expected size to remain 2 after retaining nil elements, got %d", list.Size())
	}

	// Test with empty elements
	result, err = list.FastRetainAll([]int{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected FastRetainAll to return true for empty elements, got false")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size to be 0 after retaining empty elements, got %d", list.Size())
	}

	// Test with non-existent elements
	list = NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3})
	result, err = list.FastRetainAll([]int{4, 5})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected FastRetainAll to return true for non-existent elements, got false")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size to be 0 after retaining non-existent elements, got %d", list.Size())
	}

	// Test with empty source list
	emptyList := NewArrayList[int]()
	result, err = emptyList.FastRetainAll([]int{1, 2, 3})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	assert.False(t, result)
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0 after retaining elements in empty list, got %d", emptyList.Size())
	}

	// Test with empty source list and empty elements
	emptyList = NewArrayList[int]()
	result, err = emptyList.FastRetainAll([]int{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	assert.False(t, result)
	if emptyList.Size() != 0 {
		t.Errorf("Expected size to remain 0 after retaining empty elements in empty list, got %d", emptyList.Size())
	}
}

// TestArrayList_FastSubList_ErrorCases tests error handling in FastSubList
func TestArrayList_FastSubList_ErrorCases(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 5})

	// Test with negative fromIndex
	_, err := list.FastSubList(-1, 3)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with toIndex greater than size
	_, err = list.FastSubList(0, 6)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())

	// Test with fromIndex greater than toIndex
	_, err = list.FastSubList(3, 2)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IllegalArgumentError), err.Error())

	// Test with empty list
	emptyList := NewArrayList[int]()
	_, err = emptyList.FastSubList(0, 0)
	assert.NoError(t, err) // Should succeed as 0,0 is valid for empty list

	// Test with fromIndex equal to size
	_, err = list.FastSubList(5, 5)
	assert.NoError(t, err) // Should succeed as 5,5 is valid for list of size 5

	// Test with both indices equal
	_, err = list.FastSubList(2, 2)
	assert.NoError(t, err) // Should succeed as 2,2 is valid

	// Test with fromIndex equal to toIndex and both at size
	_, err = list.FastSubList(5, 5)
	assert.NoError(t, err) // Should succeed as 5,5 is valid for list of size 5

	// Test with both indices at 0
	_, err = list.FastSubList(0, 0)
	assert.NoError(t, err) // Should succeed as 0,0 is valid

	// Test with fromIndex at 0 and toIndex at size
	_, err = list.FastSubList(0, 5)
	assert.NoError(t, err) // Should succeed as 0,5 is valid for list of size 5

	// Test with negative toIndex
	_, err = list.FastSubList(0, -1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IllegalArgumentError), err.Error())

	// Test with both indices negative
	_, err = list.FastSubList(-2, -1)
	assert.Error(t, err)
	assert.Equal(t, string(errcodes.IndexOutOfBoundsError), err.Error())
}

// TestArrayList_FastLastIndexOf tests FastLastIndexOf operation
func TestArrayList_FastLastIndexOf(t *testing.T) {
	list := NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3, 2, 4, 2, 5})

	// Test finding last occurrence of existing element
	index := list.FastLastIndexOf(2)
	assert.Equal(t, 5, index) // Should find the last occurrence of 2

	// Test finding non-existent element
	index = list.FastLastIndexOf(6)
	assert.Equal(t, -1, index) // Should return -1 for non-existent element

	// Test with empty list
	emptyList := NewArrayList[int]()
	index = emptyList.FastLastIndexOf(1)
	assert.Equal(t, -1, index) // Should return -1 for empty list

	// Test with single element list
	singleList := NewArrayList[int]()
	singleList.Add(1)
	index = singleList.FastLastIndexOf(2)
	assert.Equal(t, -1, index) // Should return -1 for non-existent element

	// Test with list containing only the target element
	singleList = NewArrayList[int]()
	singleList.Add(1)
	index = singleList.FastLastIndexOf(1)
	assert.Equal(t, 0, index) // Should find the element at index 0

	// Test with list containing multiple occurrences at the end
	list = NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 5, 5, 5})
	index = list.FastLastIndexOf(5)
	assert.Equal(t, 6, index) // Should find the last occurrence of 5

	// Test with list containing the element only at the beginning
	list = NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 5})
	index = list.FastLastIndexOf(1)
	assert.Equal(t, 0, index) // Should find the element at index 0

	// Test with list containing the element only at the end
	list = NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 5})
	index = list.FastLastIndexOf(5)
	assert.Equal(t, 4, index) // Should find the element at index 4

	// Test with list containing the element in the middle
	list = NewArrayList[int]()
	list.AddAllBatch([]int{1, 2, 3, 4, 5})
	index = list.FastLastIndexOf(3)
	assert.Equal(t, 2, index) // Should find the element at index 2
}

// TestArrayList_FastContains tests FastContains operation
func TestArrayList_FastContains(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	assert.False(t, emptyList.FastContains(1))

	// Test with single element
	singleList := NewArrayList[int]()
	singleList.Add(1)
	assert.True(t, singleList.FastContains(1))
	assert.False(t, singleList.FastContains(2))

	// Test with small list (using linear search)
	smallList := NewArrayList[int]()
	smallList.AddAllBatch([]int{1, 2, 3, 4, 5})
	assert.True(t, smallList.FastContains(3))
	assert.False(t, smallList.FastContains(6))

	// Test with large list (using map)
	largeList := NewArrayList[int]()
	for i := 0; i < 1001; i++ {
		largeList.Add(i)
	}
	assert.True(t, largeList.FastContains(500))
	assert.False(t, largeList.FastContains(1001))

	// Test with duplicate elements
	duplicateList := NewArrayList[int]()
	duplicateList.AddAllBatch([]int{1, 2, 2, 3, 2, 4})
	assert.True(t, duplicateList.FastContains(2))

	// Test with elements at boundaries
	boundaryList := NewArrayList[int]()
	boundaryList.AddAllBatch([]int{1, 2, 3, 4, 5})
	assert.True(t, boundaryList.FastContains(1)) // First element
	assert.True(t, boundaryList.FastContains(5)) // Last element

	// Test with list containing nil (for pointer types)
	stringList := NewArrayList[*string]()
	var nilStr *string
	str1 := "test"
	stringList.Add(nilStr)
	stringList.Add(&str1)
	assert.True(t, stringList.FastContains(nilStr))
	assert.True(t, stringList.FastContains(&str1))

	// Test with list after clearing
	clearedList := NewArrayList[int]()
	clearedList.AddAllBatch([]int{1, 2, 3})
	clearedList.Clear()
	assert.False(t, clearedList.FastContains(1))

	// Test with list after removing elements
	removedList := NewArrayList[int]()
	removedList.AddAllBatch([]int{1, 2, 3, 4, 5})
	removedList.Remove(3)
	assert.False(t, removedList.FastContains(3))
	assert.True(t, removedList.FastContains(4))

	// Test with list at threshold size (1000 elements)
	thresholdList := NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		thresholdList.Add(i)
	}
	assert.True(t, thresholdList.FastContains(999))
	assert.False(t, thresholdList.FastContains(1000))

	// Test with list just above threshold size (1001 elements)
	aboveThresholdList := NewArrayList[int]()
	for i := 0; i < 1001; i++ {
		aboveThresholdList.Add(i)
	}
	assert.True(t, aboveThresholdList.FastContains(1000))
	assert.False(t, aboveThresholdList.FastContains(1001))
}

// TestArrayList_FastIndexOf tests FastIndexOf operation
func TestArrayList_FastIndexOf(t *testing.T) {
	// Test with empty list
	emptyList := NewArrayList[int]()
	index := emptyList.FastIndexOf(1, false, nil)
	assert.Equal(t, -1, index)

	// Test with single element list
	singleList := NewArrayList[int]()
	singleList.Add(1)
	index = singleList.FastIndexOf(1, false, nil)
	assert.Equal(t, 0, index)
	index = singleList.FastIndexOf(2, false, nil)
	assert.Equal(t, -1, index)

	// Test with unsorted list and nil comparator
	unsortedList := NewArrayList[int]()
	unsortedList.AddAllBatch([]int{3, 1, 4, 2, 5})
	index = unsortedList.FastIndexOf(4, false, nil)
	assert.Equal(t, 2, index)
	index = unsortedList.FastIndexOf(6, false, nil)
	assert.Equal(t, -1, index)

	// Test with sorted list and nil comparator
	sortedList := NewArrayList[int]()
	sortedList.AddAllBatch([]int{1, 2, 3, 4, 5})
	index = sortedList.FastIndexOf(3, true, nil)
	assert.Equal(t, 2, index) // Should use linear search since comparator is nil
	index = sortedList.FastIndexOf(6, true, nil)
	assert.Equal(t, -1, index)

	// Test with sorted list and valid comparator
	comparator := &IntComparator{}
	index = sortedList.FastIndexOf(3, true, comparator)
	assert.Equal(t, 2, index) // Should use binary search
	index = sortedList.FastIndexOf(6, true, comparator)
	assert.Equal(t, -1, index)

	// Test with duplicate elements
	duplicateList := NewArrayList[int]()
	duplicateList.AddAllBatch([]int{1, 2, 2, 3, 2, 4})
	index = duplicateList.FastIndexOf(2, false, nil)
	assert.Equal(t, 1, index) // Should find first occurrence

	// Test with elements at boundaries
	boundaryList := NewArrayList[int]()
	boundaryList.AddAllBatch([]int{1, 2, 3, 4, 5})
	index = boundaryList.FastIndexOf(1, false, nil)
	assert.Equal(t, 0, index) // First element
	index = boundaryList.FastIndexOf(5, false, nil)
	assert.Equal(t, 4, index) // Last element

	// Test with list containing nil (for pointer types)
	stringList := NewArrayList[*string]()
	var nilStr *string
	str1 := "test"
	stringList.Add(nilStr)
	stringList.Add(&str1)
	index = stringList.FastIndexOf(nilStr, false, nil)
	assert.Equal(t, 0, index)
	index = stringList.FastIndexOf(&str1, false, nil)
	assert.Equal(t, 1, index)

	// Test with list after clearing
	clearedList := NewArrayList[int]()
	clearedList.AddAllBatch([]int{1, 2, 3})
	clearedList.Clear()
	index = clearedList.FastIndexOf(1, false, nil)
	assert.Equal(t, -1, index)

	// Test with list after removing elements
	removedList := NewArrayList[int]()
	removedList.AddAllBatch([]int{1, 2, 3, 4, 5})
	removedList.Remove(3)
	index = removedList.FastIndexOf(3, false, nil)
	assert.Equal(t, -1, index)
	index = removedList.FastIndexOf(4, false, nil)
	assert.Equal(t, 2, index)

	// Test with incorrectly marked sorted list
	incorrectlySortedList := NewArrayList[int]()
	incorrectlySortedList.AddAllBatch([]int{3, 1, 4, 2, 5})
	index = incorrectlySortedList.FastIndexOf(4, true, comparator)
	assert.Equal(t, 2, index) // Should still find the element, but might not be optimal

	// Test with large sorted list
	largeSortedList := NewArrayList[int]()
	for i := 0; i < 1000; i++ {
		largeSortedList.Add(i)
	}
	index = largeSortedList.FastIndexOf(500, true, comparator)
	assert.Equal(t, 500, index)
	index = largeSortedList.FastIndexOf(1000, true, comparator)
	assert.Equal(t, -1, index)

	// Test with large unsorted list
	largeUnsortedList := NewArrayList[int]()
	for i := 999; i >= 0; i-- {
		largeUnsortedList.Add(i)
	}
	index = largeUnsortedList.FastIndexOf(500, false, nil)
	assert.Equal(t, 499, index)
	index = largeUnsortedList.FastIndexOf(1000, false, nil)
	assert.Equal(t, -1, index)
}

// TestArrayList_AddAtIndex tests AddAtIndex operation
func TestArrayList_AddAtIndex(t *testing.T) {
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
			list := NewArrayListWithInitialCollection(tt.initialValues)
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

// TestArrayList_AddAtIndex_Concurrent tests concurrent AddAtIndex operations
func TestArrayList_AddAtIndex_Concurrent(t *testing.T) {
	list := NewArrayList[int]()
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

// TestArrayList_AddAtIndex_Capacity tests capacity management during AddAtIndex
func TestArrayList_AddAtIndex_Capacity(t *testing.T) {
	list := NewArrayList[int]()
	initialCapacity := list.Capacity()

	// Add elements until we exceed initial capacity
	for i := 0; i < initialCapacity+1; i++ {
		err := list.AddAtIndex(i, i)
		if err != nil {
			t.Errorf("AddAtIndex failed at index %d: %v", i, err)
		}
	}

	// Verify capacity was increased
	if list.Capacity() <= initialCapacity {
		t.Errorf("Capacity was not increased. Initial: %d, Current: %d", initialCapacity, list.Capacity())
	}

	// Verify all elements are present
	values := list.ToArray()
	if len(values) != initialCapacity+1 {
		t.Errorf("Expected %d elements, got %d", initialCapacity+1, len(values))
	}

	for i := 0; i < len(values); i++ {
		if values[i] != i {
			t.Errorf("Expected value %d at index %d, got %d", i, i, values[i])
		}
	}
}
func TestStringComparator_Compare(t *testing.T) {
	comparator := &StringComparator{}

	// Test cases for string comparison
	testCases := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{
			name:     "Equal strings",
			a:        "hello",
			b:        "hello",
			expected: 0,
		},
		{
			name:     "First string less than second",
			a:        "apple",
			b:        "banana",
			expected: -1,
		},
		{
			name:     "First string greater than second",
			a:        "zebra",
			b:        "apple",
			expected: 1,
		},
		{
			name:     "Empty strings",
			a:        "",
			b:        "",
			expected: 0,
		},
		{
			name:     "First string empty",
			a:        "",
			b:        "hello",
			expected: -1,
		},
		{
			name:     "Second string empty",
			a:        "hello",
			b:        "",
			expected: 1,
		},
		{
			name:     "Case sensitive comparison",
			a:        "Hello",
			b:        "hello",
			expected: -1,
		},
		{
			name:     "Special characters",
			a:        "!@#$",
			b:        "1234",
			expected: -1,
		},
		{
			name:     "Unicode characters",
			a:        "こんにちは",
			b:        "世界",
			expected: -1,
		},
		{
			name:     "Numbers as strings",
			a:        "123",
			b:        "456",
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := comparator.Compare(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Compare(%q, %q) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

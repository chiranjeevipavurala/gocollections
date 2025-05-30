package queues

import (
	"reflect"
	"testing"

	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
	"github.com/chiranjeevipavurala/gocollections/lists"
)

type MaxIntComparator struct{}

func (c *MaxIntComparator) Compare(a, b int) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}

func TestNewPriorityQueue(t *testing.T) {
	// Test with nil comparator
	pq := NewPriorityQueue[int](nil)
	if pq != nil {
		t.Error("NewPriorityQueue should return nil when comparator is nil")
	}

	// Test with valid comparator
	pq = NewPriorityQueue[int](&IntComparator[int]{})
	if pq == nil {
		t.Error("NewPriorityQueue returned nil")
	}
	if !pq.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if pq.Size() != 0 {
		t.Error("New queue should have size 0")
	}
}

func TestNewPriorityQueueWithCapacity(t *testing.T) {
	// Test with nil comparator
	pq := NewPriorityQueueWithCapacity[int](20, nil)
	if pq != nil {
		t.Error("NewPriorityQueueWithCapacity should return nil when comparator is nil")
	}

	// Test with initialCapacity < 1
	pq = NewPriorityQueueWithCapacity[int](0, &IntComparator[int]{})
	if pq == nil {
		t.Error("NewPriorityQueueWithCapacity should not return nil for invalid capacity")
	}
	if cap(pq.(*PriorityQueue[int]).elements) != DefaultCapacity {
		t.Errorf("Expected capacity %d for invalid capacity, got %d", DefaultCapacity, cap(pq.(*PriorityQueue[int]).elements))
	}

	pq = NewPriorityQueueWithCapacity[int](-1, &IntComparator[int]{})
	if pq == nil {
		t.Error("NewPriorityQueueWithCapacity should not return nil for negative capacity")
	}
	if cap(pq.(*PriorityQueue[int]).elements) != DefaultCapacity {
		t.Errorf("Expected capacity %d for negative capacity, got %d", DefaultCapacity, cap(pq.(*PriorityQueue[int]).elements))
	}

	// Test with valid comparator
	pq = NewPriorityQueueWithCapacity[int](20, &IntComparator[int]{})
	if pq == nil {
		t.Error("NewPriorityQueueWithCapacity returned nil")
	}
	if !pq.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if cap(pq.(*PriorityQueue[int]).elements) != 20 {
		t.Errorf("Expected capacity 20, got %d", cap(pq.(*PriorityQueue[int]).elements))
	}
}

func TestPriorityQueue_Add(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test adding elements
	if !pq.Add(5) {
		t.Error("Add should return true")
	}
	if !pq.Add(3) {
		t.Error("Add should return true")
	}
	if !pq.Add(7) {
		t.Error("Add should return true")
	}

	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	// Test that elements are ordered correctly
	val, err := pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}
}

func TestPriorityQueue_Offer(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	if !pq.Offer(5) {
		t.Error("Offer should return true")
	}
	if !pq.Offer(3) {
		t.Error("Offer should return true")
	}

	if pq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", pq.Size())
	}
}

func TestPriorityQueue_Poll(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test empty queue
	val, err := pq.Poll()
	if err == nil {
		t.Error("Poll on empty queue should return error")
	}
	if err.Error() != string(errcodes.NoSuchElementError) {
		t.Errorf("Expected QueueIsEmptyError, got %v", err)
	}
	if val != nil {
		t.Error("Poll on empty queue should return nil value")
	}

	// Test with elements
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	val, err = pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}

	val, err = pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 5 {
		t.Errorf("Expected 5, got %d", *val)
	}

	val, err = pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 7 {
		t.Errorf("Expected 7, got %d", *val)
	}

	// Test empty queue again after removing all elements
	val, err = pq.Poll()
	if err == nil {
		t.Error("Poll on empty queue should return error")
	}
	if err.Error() != string(errcodes.NoSuchElementError) {
		t.Errorf("Expected QueueIsEmptyError, got %v", err)
	}
	if val != nil {
		t.Error("Poll on empty queue should return nil value")
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test empty queue
	val, err := pq.Peek()
	if err != nil {
		t.Errorf("Peek returned error: %v", err)
	}
	if val != nil {
		t.Error("Peek on empty queue should return nil")
	}

	// Test with elements
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	val, err = pq.Peek()
	if err != nil {
		t.Errorf("Peek returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}

	// Verify element wasn't removed
	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}
}

func TestPriorityQueue_Element(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test empty queue
	_, err := pq.Element()
	if err == nil {
		t.Error("Element on empty queue should return error")
	}

	// Test with elements
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	val, err := pq.Element()
	if err != nil {
		t.Errorf("Element returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}
}

func TestPriorityQueue_Element_Empty(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test Element() on empty queue
	element, err := pq.Element()
	if err == nil {
		t.Errorf("Expected NoSuchElementError, got no error")
	}
	if element != nil {
		t.Errorf("Expected nil element, got %v", element)
	}
}

func TestPriorityQueue_RemoveHead(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test empty queue
	_, err := pq.RemoveHead()
	if err == nil {
		t.Error("RemoveHead on empty queue should return error")
	} else if err.Error() != string(errcodes.NoSuchElementError) {
		t.Errorf("Expected NoSuchElementError, got %v", err)
	}

	// Test with elements
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	val, err := pq.RemoveHead()
	if err != nil {
		t.Errorf("RemoveHead returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}

	if pq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", pq.Size())
	}
}

func TestPriorityQueue_Clear(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])

	// Add elements
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	// Verify initial state
	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}
	if len(pq.elements) != 3 {
		t.Errorf("Expected elements length 3, got %d", len(pq.elements))
	}

	// Clear the queue
	pq.Clear()

	// Verify queue is empty
	if !pq.IsEmpty() {
		t.Error("Queue should be empty after Clear")
	}
	if pq.Size() != 0 {
		t.Errorf("Expected size 0, got %d", pq.Size())
	}
	if len(pq.elements) != 0 {
		t.Errorf("Expected elements length 0, got %d", len(pq.elements))
	}

	// Add elements again to verify queue still works
	pq.Add(1)
	pq.Add(2)
	if pq.Size() != 2 {
		t.Errorf("Expected size 2 after adding new elements, got %d", pq.Size())
	}
	if len(pq.elements) != 2 {
		t.Errorf("Expected elements length 2 after adding new elements, got %d", len(pq.elements))
	}
}

func TestPriorityQueue_Contains(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	if !pq.Contains(5) {
		t.Error("Queue should contain 5")
	}
	if !pq.Contains(3) {
		t.Error("Queue should contain 3")
	}
	if pq.Contains(4) {
		t.Error("Queue should not contain 4")
	}
}

func TestPriorityQueue_Remove(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	if !pq.Remove(3) {
		t.Error("Remove should return true for existing element")
	}
	if pq.Remove(4) {
		t.Error("Remove should return false for non-existing element")
	}

	if pq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", pq.Size())
	}

	// Verify heap property is maintained
	val, err := pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 5 {
		t.Errorf("Expected 5, got %d", *val)
	}
}

func TestPriorityQueue_AddAll(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test with nil collection
	if pq.AddAll(nil) {
		t.Error("AddAll with nil collection should return false")
	}

	// Test with empty collection
	emptyCollection := lists.NewArrayList[int]()
	if pq.AddAll(emptyCollection) {
		t.Error("AddAll with empty collection should return false")
	}

	// Create a collection with elements
	collection := lists.NewArrayList[int]()
	collection.Add(5)
	collection.Add(3)
	collection.Add(7)

	if !pq.AddAll(collection) {
		t.Error("AddAll should return true")
	}

	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	// Verify elements are ordered correctly
	val, err := pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}
}

func TestPriorityQueue_RemoveAll(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test with nil collection
	if pq.RemoveAll(nil) {
		t.Error("RemoveAll with nil collection should return false")
	}

	// Test with empty collection
	emptyCollection := lists.NewArrayList[int]()
	if pq.RemoveAll(emptyCollection) {
		t.Error("RemoveAll with empty collection should return false")
	}

	pq.Add(5)
	pq.Add(3)
	pq.Add(7)
	pq.Add(9)

	// Create a collection to remove
	collection := lists.NewArrayList[int]()
	collection.Add(3)
	collection.Add(7)

	if !pq.RemoveAll(collection) {
		t.Error("RemoveAll should return true")
	}

	if pq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", pq.Size())
	}

	// Verify remaining elements
	if !pq.Contains(5) || !pq.Contains(9) {
		t.Error("Queue should still contain 5 and 9")
	}
}

func TestPriorityQueue_ContainsAll(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	// Create a collection
	collection := lists.NewArrayList[int]()
	collection.Add(3)
	collection.Add(7)

	contains, err := pq.ContainsAll(collection)
	if err != nil {
		t.Errorf("ContainsAll returned error: %v", err)
	}
	if !contains {
		t.Error("Queue should contain all elements from collection")
	}

	// Test with non-contained element
	collection.Add(4)
	contains, err = pq.ContainsAll(collection)
	if err != nil {
		t.Errorf("ContainsAll returned error: %v", err)
	}
	if contains {
		t.Error("Queue should not contain all elements from collection")
	}
}

func TestPriorityQueue_ContainsAll_NullCollection(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test ContainsAll with nil collection
	contains, err := pq.ContainsAll(nil)
	if err == nil {
		t.Error("ContainsAll with nil collection should return error")
	}
	if err.Error() != string(errcodes.NullPointerError) {
		t.Errorf("Expected NullPointerError, got %v", err)
	}
	if contains {
		t.Error("ContainsAll with nil collection should return false")
	}

	// Add some elements to the queue
	pq.Add(1)
	pq.Add(2)
	pq.Add(3)

	// Test again with nil collection after adding elements
	contains, err = pq.ContainsAll(nil)
	if err == nil {
		t.Error("ContainsAll with nil collection should return error")
	}
	if err.Error() != string(errcodes.NullPointerError) {
		t.Errorf("Expected NullPointerError, got %v", err)
	}
	if contains {
		t.Error("ContainsAll with nil collection should return false")
	}
}

func TestPriorityQueue_Equals(t *testing.T) {
	pq1 := NewPriorityQueue[int](&IntComparator[int]{})
	pq2 := NewPriorityQueue[int](&IntComparator[int]{})

	pq1.Add(5)
	pq1.Add(3)
	pq1.Add(7)

	pq2.Add(5)
	pq2.Add(3)
	pq2.Add(7)

	if !pq1.Equals(pq2) {
		t.Error("Queues should be equal")
	}

	pq2.Add(4)
	if pq1.Equals(pq2) {
		t.Error("Queues should not be equal")
	}
}

func TestPriorityQueue_Equals_NilCollection(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test Equals with nil collection
	result := pq.Equals(nil)
	if result {
		t.Error("Equals with nil collection should return false")
	}

	// Add some elements to the queue
	pq.Add(1)
	pq.Add(2)
	pq.Add(3)

	// Test again with nil collection after adding elements
	result = pq.Equals(nil)
	if result {
		t.Error("Equals with nil collection should return false")
	}
}

func TestPriorityQueue_Equals_DifferentElements(t *testing.T) {
	pq1 := NewPriorityQueue[int](&IntComparator[int]{})
	pq2 := NewPriorityQueue[int](&IntComparator[int]{})

	// Test case 1: Different sizes
	pq1.Add(1)
	pq1.Add(2)
	pq1.Add(3)

	pq2.Add(1)
	pq2.Add(2)

	if pq1.Equals(pq2) {
		t.Error("Queues with different sizes should not be equal")
	}

	// Test case 2: Same size but different elements
	pq1.Clear()
	pq2.Clear()

	pq1.Add(1)
	pq1.Add(2)
	pq1.Add(3)

	pq2.Add(1)
	pq2.Add(2)
	pq2.Add(4)

	if pq1.Equals(pq2) {
		t.Error("Queues with different elements should not be equal")
	}

	// Test case 3: Same elements in different order
	pq1.Clear()
	pq2.Clear()

	pq1.Add(1)
	pq1.Add(2)
	pq1.Add(3)

	pq2.Add(3)
	pq2.Add(1)
	pq2.Add(2)

	if !pq1.Equals(pq2) {
		t.Error("Queues with same elements in different order should be equal")
	}

	// Test case 4: Different types of collections
	pq1.Clear()
	list := lists.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if pq1.Equals(list) {
		t.Error("Queue should not be equal to a different type of collection")
	}
}

func TestPriorityQueue_Iterator(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	iterator := pq.Iterator()

	// Test HasNext
	if !iterator.HasNext() {
		t.Error("Iterator should have next element")
	}

	// Test Next
	val, err := iterator.Next()
	if err != nil {
		t.Errorf("Next returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}

	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Next returned error: %v", err)
	}
	if *val != 5 {
		t.Errorf("Expected 5, got %d", *val)
	}

	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Next returned error: %v", err)
	}
	if *val != 7 {
		t.Errorf("Expected 7, got %d", *val)
	}

	// Test end of iteration
	if iterator.HasNext() {
		t.Error("Iterator should not have more elements")
	}

	_, err = iterator.Next()
	if err == nil {
		t.Error("Next should return error at end of iteration")
	}
}

func TestPriorityQueue_ToArray(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	arr := pq.ToArray()
	if len(arr) != 3 {
		t.Errorf("Expected array length 3, got %d", len(arr))
	}

	// Note: ToArray doesn't guarantee order
	contains := make(map[int]bool)
	for _, v := range arr {
		contains[v] = true
	}

	if !contains[3] || !contains[5] || !contains[7] {
		t.Error("Array should contain all elements")
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{0, 0, 0},
		{-1, 1, 1},
		{-2, -1, -1},
		{100, 99, 100},
		{-100, -99, -99},
	}

	for _, tc := range testCases {
		result := max(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("max(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestPriorityQueue_RemoveIf(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])

	// Test case: null predicate
	removed := pq.RemoveIf(nil)
	if removed {
		t.Error("RemoveIf should return false when predicate is nil")
	}
	if pq.Size() != 0 {
		t.Errorf("Expected size 0 after null predicate, got %d", pq.Size())
	}

	// Test case 1: Empty queue
	removed = pq.RemoveIf(func(x int) bool {
		return x%2 == 0
	})
	if removed {
		t.Error("RemoveIf should return false for empty queue")
	}

	// Test case 2: Remove even numbers
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)
	pq.Add(9)
	pq.Add(2)
	pq.Add(4)
	pq.Add(6)
	pq.Add(8)

	removed = pq.RemoveIf(func(x int) bool {
		return x%2 == 0
	})
	if !removed {
		t.Error("RemoveIf should return true as even numbers were removed")
	}

	if pq.Size() != 4 {
		t.Errorf("Expected size 4, got %d", pq.Size())
	}

	// Verify remaining elements (odd numbers)
	contains := make(map[int]bool)
	for _, v := range pq.ToArray() {
		contains[v] = true
	}

	if !contains[3] || !contains[5] || !contains[7] || !contains[9] {
		t.Error("Queue should still contain odd numbers")
	}

	// Test case 3: Remove numbers greater than 5
	removed = pq.RemoveIf(func(x int) bool {
		return x > 5
	})
	if !removed {
		t.Error("RemoveIf should return true as numbers > 5 were removed")
	}

	if pq.Size() != 2 {
		t.Errorf("Expected size 2, got %d", pq.Size())
	}

	// Verify remaining elements (numbers <= 5)
	contains = make(map[int]bool)
	for _, v := range pq.ToArray() {
		contains[v] = true
	}

	if !contains[3] || !contains[5] {
		t.Error("Queue should still contain 3 and 5")
	}

	// Test case 4: Remove all elements
	removed = pq.RemoveIf(func(x int) bool {
		return true
	})
	if !removed {
		t.Error("RemoveIf should return true as all elements were removed")
	}

	if !pq.IsEmpty() {
		t.Error("Queue should be empty")
	}

	// Test case 5: Remove from empty queue
	removed = pq.RemoveIf(func(x int) bool {
		return true
	})
	if removed {
		t.Error("RemoveIf should return false for empty queue")
	}

	// Test case 6: Predicate that doesn't match any elements
	pq.Add(1)
	pq.Add(2)
	pq.Add(3)
	removed = pq.RemoveIf(func(x int) bool {
		return x > 10
	})
	if removed {
		t.Error("RemoveIf should return false when no elements match predicate")
	}

	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	// Test case 7: Remove elements that would break heap property
	pq.Clear()
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)
	pq.Add(1)
	pq.Add(9)

	removed = pq.RemoveIf(func(x int) bool {
		return x < 4
	})
	if !removed {
		t.Error("RemoveIf should return true as elements were removed")
	}

	// Verify heap property is maintained
	val, err := pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 5 {
		t.Errorf("Expected 5, got %d", *val)
	}

	// Test case 8: Concurrent modifications
	pq.Clear()
	for i := 0; i < 100; i++ {
		pq.Add(i)
	}

	removed = pq.RemoveIf(func(x int) bool {
		return x%2 == 0
	})
	if !removed {
		t.Error("RemoveIf should return true as elements were removed")
	}

	// Verify all remaining elements are odd
	for _, v := range pq.ToArray() {
		if v%2 == 0 {
			t.Errorf("Found even number %d in queue after removing even numbers", v)
		}
	}
}

func TestPriorityQueue_ForEach(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])

	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	sum := 0
	pq.ForEach(func(x int) {
		sum += x
	})

	if sum != 15 {
		t.Errorf("Expected sum 15, got %d", sum)
	}

	// Test case: null action
	pq.ForEach(nil) // Should not panic
	if pq.Size() != 3 {
		t.Errorf("Expected size 3 after null action, got %d", pq.Size())
	}
}

func TestPriorityQueue_GetComparator(t *testing.T) {
	comparator := &IntComparator[int]{}
	pq := NewPriorityQueue[int](comparator).(*PriorityQueue[int])

	if pq.GetComparator() != comparator {
		t.Error("GetComparator should return the same comparator")
	}
}

func TestPriorityQueue_NewPriorityQueueFromCollection(t *testing.T) {
	// Test with nil comparator
	pq := NewPriorityQueueFromCollection[int](nil, nil)
	if pq != nil {
		t.Error("NewPriorityQueueFromCollection should return nil when comparator is nil")
	}

	// Test with nil collection
	pq = NewPriorityQueueFromCollection[int](nil, &IntComparator[int]{})
	if pq == nil {
		t.Error("NewPriorityQueueFromCollection should not return nil")
	}
	if !pq.IsEmpty() {
		t.Error("New queue from nil collection should be empty")
	}
	if pq.Size() != 0 {
		t.Errorf("Expected size 0, got %d", pq.Size())
	}

	// Test with non-nil collection
	collection := lists.NewArrayList[int]()
	collection.Add(5)
	collection.Add(3)
	collection.Add(7)

	pq = NewPriorityQueueFromCollection[int](collection, &IntComparator[int]{})

	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	// Verify elements are ordered correctly
	val, err := pq.Poll()
	if err != nil {
		t.Errorf("Poll returned error: %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected 3, got %d", *val)
	}
}

func TestPriorityQueue_ConcurrentOperations(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Test concurrent Add operations
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			pq.Add(i)
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	if pq.Size() != 10 {
		t.Errorf("Expected size 10, got %d", pq.Size())
	}

	// Test concurrent Poll operations
	for i := 0; i < 10; i++ {
		go func() {
			_, err := pq.Poll()
			if err != nil {
				t.Errorf("Poll returned error: %v", err)
			}
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	if !pq.IsEmpty() {
		t.Error("Queue should be empty after all polls")
	}
}

func TestPriorityQueue_ToArrayWithType(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])

	// Add some elements
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	// Test with []int type
	result, err := pq.ToArrayWithType(reflect.TypeOf([]int{}))
	if err != nil {
		t.Errorf("ToArrayWithType returned error: %v", err)
	}

	arr, ok := result.([]int)
	if !ok {
		t.Error("Result should be of type []int")
	}

	if len(arr) != 3 {
		t.Errorf("Expected array length 3, got %d", len(arr))
	}

	// Verify elements (order doesn't matter for priority queue)
	contains := make(map[int]bool)
	for _, v := range arr {
		contains[v] = true
	}

	if !contains[3] || !contains[5] || !contains[7] {
		t.Error("Array should contain all elements")
	}

	// Test with invalid type
	_, err = pq.ToArrayWithType(reflect.TypeOf(0)) // int is not an array/slice type
	if err == nil {
		t.Error("ToArrayWithType should return error for non-array type")
	}

	// Test with empty queue
	pq.Clear()
	result, err = pq.ToArrayWithType(reflect.TypeOf([]int{}))
	if err != nil {
		t.Errorf("ToArrayWithType returned error: %v", err)
	}

	arr, ok = result.([]int)
	if !ok {
		t.Error("Result should be of type []int")
	}

	if len(arr) != 0 {
		t.Errorf("Expected empty array, got length %d", len(arr))
	}
}

func TestPriorityQueue_NewPriorityQueueFromSortedSet(t *testing.T) {
	// Test with nil sorted set
	pq := NewPriorityQueueFromSortedSet[int](nil)
	if pq == nil {
		t.Error("NewPriorityQueueFromSortedSet should not return nil")
	}
	if !pq.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if pq.Size() != 0 {
		t.Errorf("Expected size 0, got %d", pq.Size())
	}

	// Test with non-nil sorted set
	// Create a priority queue directly with the same elements
	pq = NewPriorityQueue[int](&IntComparator[int]{})
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)

	if pq == nil {
		t.Error("NewPriorityQueueFromSortedSet should not return nil")
	}
	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	// Verify elements are in the queue
	contains := make(map[int]bool)
	for _, v := range pq.ToArray() {
		contains[v] = true
	}

	if !contains[3] || !contains[5] || !contains[7] {
		t.Error("Queue should contain all elements from sorted set")
	}
}

func TestPriorityQueue_RightChild(t *testing.T) {
	// Create a priority queue with a simple integer comparator
	comparator := &IntComparator[int]{}
	pq := NewPriorityQueue[int](comparator).(*PriorityQueue[int])

	// Test cases for rightChild method
	testCases := []struct {
		index    int
		expected int
	}{
		{0, 2},  // Root node (0) -> right child at index 2
		{1, 4},  // Node at index 1 -> right child at index 4
		{2, 6},  // Node at index 2 -> right child at index 6
		{3, 8},  // Node at index 3 -> right child at index 8
		{4, 10}, // Node at index 4 -> right child at index 10
	}

	for _, tc := range testCases {
		result := pq.rightChild(tc.index)
		if result != tc.expected {
			t.Errorf("rightChild(%d) = %d; want %d", tc.index, result, tc.expected)
		}
	}
}

func TestPriorityQueue_EnsureCapacity(t *testing.T) {
	// Test case 1: No capacity increase needed
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])
	initialCapacity := cap(pq.elements)
	pq.ensureCapacity(initialCapacity - 1)
	if cap(pq.elements) != initialCapacity {
		t.Errorf("Capacity should not change when not needed, got %d, want %d", cap(pq.elements), initialCapacity)
	}

	// Test case 2: Capacity increase needed
	pq = NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])
	initialCapacity = cap(pq.elements)
	pq.ensureCapacity(initialCapacity + 1)
	if cap(pq.elements) <= initialCapacity {
		t.Errorf("Capacity should increase when needed, got %d, want > %d", cap(pq.elements), initialCapacity)
	}

	// Test case 3: Growth factor calculation
	pq = NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])
	initialCapacity = cap(pq.elements)
	expectedNewCapacity := int(float64(initialCapacity) * GrowthFactor)
	pq.ensureCapacity(expectedNewCapacity)
	if cap(pq.elements) != expectedNewCapacity {
		t.Errorf("Capacity should grow by growth factor, got %d, want %d", cap(pq.elements), expectedNewCapacity)
	}

	// Test case 4: Elements are preserved after capacity increase
	pq = NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])
	// Add some elements
	for i := 0; i < 5; i++ {
		pq.Add(i)
	}
	// Create a copy of elements with the same length
	originalElements := make([]int, len(pq.elements))
	copy(originalElements, pq.elements)
	originalLength := len(pq.elements)

	// Force capacity increase
	pq.ensureCapacity(cap(pq.elements) + 1)

	// Verify elements are preserved
	if len(pq.elements) != originalLength {
		t.Errorf("Length changed after capacity increase, got %d, want %d", len(pq.elements), originalLength)
	}
	for i := 0; i < originalLength; i++ {
		if pq.elements[i] != originalElements[i] {
			t.Errorf("Element at index %d changed after capacity increase, got %d, want %d",
				i, pq.elements[i], originalElements[i])
		}
	}

	// Test case 5: Large capacity increase
	pq = NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])
	initialCapacity = cap(pq.elements)
	largeCapacity := initialCapacity * 10
	pq.ensureCapacity(largeCapacity)
	if cap(pq.elements) < largeCapacity {
		t.Errorf("Capacity should accommodate large increase, got %d, want >= %d",
			cap(pq.elements), largeCapacity)
	}

	// Test case 6: Custom initial capacity
	customCapacity := 20
	pq = NewPriorityQueueWithCapacity[int](customCapacity, &IntComparator[int]{}).(*PriorityQueue[int])
	if cap(pq.elements) != customCapacity {
		t.Errorf("Initial capacity should be %d, got %d", customCapacity, cap(pq.elements))
	}

	// Test capacity increase with custom initial capacity
	pq.ensureCapacity(customCapacity + 1)
	if cap(pq.elements) <= customCapacity {
		t.Errorf("Capacity should increase from custom capacity, got %d, want > %d",
			cap(pq.elements), customCapacity)
	}

	// Test case 7: Multiple capacity increases
	pq = NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])
	initialCapacity = cap(pq.elements)

	// Force multiple capacity increases
	for i := 0; i < 5; i++ {
		oldCapacity := cap(pq.elements)
		pq.ensureCapacity(oldCapacity + 1)
		if cap(pq.elements) <= oldCapacity {
			t.Errorf("Capacity should increase on iteration %d, got %d, want > %d",
				i, cap(pq.elements), oldCapacity)
		}
	}
}

func TestPriorityQueue_RemoveElement(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])

	// Test case 1: Remove element that requires sifting up
	pq.Add(5) // [5]
	pq.Add(3) // [3,5]
	pq.Add(7) // [3,5,7]
	pq.Add(1) // [1,3,7,5]
	pq.Add(9) // [1,3,7,5,9]

	// Remove 5 which will require sifting up 9
	removed := pq.removeElement(5)
	if !removed {
		t.Error("removeElement should return true for existing element")
	}

	// Verify heap property is maintained
	expected := []int{1, 3, 7, 9}
	for i := 0; i < len(expected); i++ {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Poll returned error: %v", err)
		}
		if *val != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, *val)
		}
	}

	// Test case 2: Remove element that requires multiple sift up operations
	pq.Clear()
	pq.Add(10) // [10]
	pq.Add(8)  // [8,10]
	pq.Add(6)  // [6,10,8]
	pq.Add(4)  // [4,6,8,10]
	pq.Add(2)  // [2,4,8,10,6]

	// Remove 4 which will require sifting up 6
	removed = pq.removeElement(4)
	if !removed {
		t.Error("removeElement should return true for existing element")
	}

	// Verify heap property is maintained
	expected = []int{2, 6, 8, 10}
	for i := 0; i < len(expected); i++ {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Poll returned error: %v", err)
		}
		if *val != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, *val)
		}
	}

	// Test case 3: Remove element that doesn't require sifting up
	pq.Clear()
	pq.Add(1) // [1]
	pq.Add(2) // [1,2]
	pq.Add(3) // [1,2,3]
	pq.Add(4) // [1,2,3,4]
	pq.Add(5) // [1,2,3,4,5]

	// Remove 5 which is at the end and doesn't require sifting up
	removed = pq.removeElement(5)
	if !removed {
		t.Error("removeElement should return true for existing element")
	}

	// Verify heap property is maintained
	expected = []int{1, 2, 3, 4}
	for i := 0; i < len(expected); i++ {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Poll returned error: %v", err)
		}
		if *val != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, *val)
		}
	}

	// Test case 4: Remove non-existent element
	removed = pq.removeElement(10)
	if removed {
		t.Error("removeElement should return false for non-existent element")
	}

	// Test case 5: Remove from empty queue
	pq.Clear()
	removed = pq.removeElement(1)
	if removed {
		t.Error("removeElement should return false for empty queue")
	}
}

func TestPriorityQueue_SiftUp(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])

	// Test case 1: Basic sift up
	pq.elements = []int{5, 3, 7, 9, 1} // [5,3,7,9,1]
	pq.siftUp(4)                       // Sift up the last element (1)

	// After sifting up, 1 should be at the root
	if pq.elements[0] != 1 {
		t.Errorf("Expected root to be 1, got %d", pq.elements[0])
	}

	// Test case 2: Multiple sift up operations
	pq.elements = []int{10, 8, 6, 4, 2} // [10,8,6,4,2]
	pq.siftUp(4)                        // Sift up the last element (2)

	// After sifting up, 2 should be at the root
	if pq.elements[0] != 2 {
		t.Errorf("Expected root to be 2, got %d", pq.elements[0])
	}

	// Test case 3: No sift up needed
	pq.elements = []int{1, 3, 5, 7, 9} // [1,3,5,7,9]
	pq.siftUp(4)                       // Try to sift up the last element (9)

	// 9 should remain at the end as it's larger than its parent
	if pq.elements[4] != 9 {
		t.Errorf("Expected last element to remain 9, got %d", pq.elements[4])
	}

	// Test case 4: Sift up from middle of heap
	pq.elements = []int{5, 10, 7, 15, 3} // [5,10,7,15,3]
	pq.siftUp(4)                         // Sift up the last element (3)

	// After sifting up, 3 should be at the root
	if pq.elements[0] != 3 {
		t.Errorf("Expected root to be 3, got %d", pq.elements[0])
	}
}

func TestPriorityQueue_Remove_WithSiftUp(t *testing.T) {
	pq := NewPriorityQueue[int](&IntComparator[int]{}).(*PriorityQueue[int])

	// Test case 1: Remove element that requires sifting up
	pq.Add(5) // [5]
	pq.Add(3) // [3,5]
	pq.Add(7) // [3,5,7]
	pq.Add(1) // [1,3,7,5]
	pq.Add(9) // [1,3,7,5,9]

	// Remove 5 which will require sifting up 9
	if !pq.Remove(5) {
		t.Error("Remove should return true for existing element")
	}

	// Verify heap property is maintained
	expected := []int{1, 3, 7, 9}
	for i := 0; i < len(expected); i++ {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Poll returned error: %v", err)
		}
		if *val != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, *val)
		}
	}

	// Test case 2: Remove element that requires multiple sift up operations
	pq.Clear()
	pq.Add(10) // [10]
	pq.Add(8)  // [8,10]
	pq.Add(6)  // [6,10,8]
	pq.Add(4)  // [4,6,8,10]
	pq.Add(2)  // [2,4,8,10,6]

	// Remove 4 which will require sifting up 6
	if !pq.Remove(4) {
		t.Error("Remove should return true for existing element")
	}

	// Verify heap property is maintained
	expected = []int{2, 6, 8, 10}
	for i := 0; i < len(expected); i++ {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Poll returned error: %v", err)
		}
		if *val != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, *val)
		}
	}

	// Test case 3: Remove root element
	pq.Clear()
	pq.Add(1) // [1]
	pq.Add(2) // [1,2]
	pq.Add(3) // [1,2,3]
	pq.Add(4) // [1,2,3,4]
	pq.Add(5) // [1,2,3,4,5]

	if !pq.Remove(1) {
		t.Error("Remove should return true for existing element")
	}

	// Verify heap property is maintained
	expected = []int{2, 3, 4, 5}
	for i := 0; i < len(expected); i++ {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Poll returned error: %v", err)
		}
		if *val != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, *val)
		}
	}

	// Test case 4: Remove last element
	pq.Clear()
	pq.Add(1) // [1]
	pq.Add(2) // [1,2]
	pq.Add(3) // [1,2,3]

	if !pq.Remove(3) {
		t.Error("Remove should return true for existing element")
	}

	// Verify heap property is maintained
	expected = []int{1, 2}
	for i := 0; i < len(expected); i++ {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Poll returned error: %v", err)
		}
		if *val != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, *val)
		}
	}
}

func TestPriorityQueue_RemoveElementWithSiftUp(t *testing.T) {
	// Create a priority queue with a comparator that orders integers in ascending order
	pq := NewPriorityQueue[int](&IntComparator[int]{})

	// Add elements in a way that will require siftUp after removal
	// We'll add elements in this order: 5, 3, 7, 2, 1
	// This creates a heap like:
	//       1
	//     /   \
	//    2     7
	//   / \
	//  5   3
	pq.Add(5)
	pq.Add(3)
	pq.Add(7)
	pq.Add(2)
	pq.Add(1)

	// Remove element 7, which will require siftUp
	// After removal, the heap should be:
	//       1
	//     /   \
	//    2     3
	//   /
	//  5
	removed := pq.Remove(7)
	if !removed {
		t.Errorf("Expected Remove to return true")
	}

	// Verify the heap property is maintained
	// The elements should be in ascending order when polled
	expected := []int{1, 2, 3, 5}
	for i, expectedVal := range expected {
		val, err := pq.Poll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if *val != expectedVal {
			t.Errorf("Expected element at position %d to be %d, got %d", i, expectedVal, *val)
		}
	}
}

func TestPriorityQueue_RemoveElementWithSiftUp_MaxHeap(t *testing.T) {
	pq := NewPriorityQueue[int](&MaxIntComparator{}).(*PriorityQueue[int])

	// Add elements to form a max heap: [9, 7, 8, 1, 2]
	pq.Add(7)
	pq.Add(8)
	pq.Add(9)
	pq.Add(1)
	pq.Add(2)
	// Heap:      9
	//          /   \
	//         7     8
	//        / \
	//       1   2

	// Remove 1 (index 3), last element (2) moves to index 3
	// 2 > its parent (7)? Yes! siftUp is called.
	removed := pq.removeElement(1)
	if !removed {
		t.Errorf("Expected removeElement to return true")
	}

	// After removal, heap should be: [9, 7, 8, 2]
	expected := []int{9, 7, 8, 2}
	actual := pq.ToArray()
	for _, v := range expected {
		found := false
		for _, a := range actual {
			if a == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected value %d in heap, but not found", v)
		}
	}
}

func TestPriorityQueue_RemoveElementWithSiftUp_MaxHeap_Trigger(t *testing.T) {
	pq := NewPriorityQueue[int](&MaxIntComparator{}).(*PriorityQueue[int])

	// Add elements to form a max heap: [10, 9, 8, 1, 7]
	pq.Add(7)
	pq.Add(8)
	pq.Add(9)
	pq.Add(1)
	pq.Add(10)
	// Heap:      10
	//           /   \
	//         9       8
	//        / \
	//       1   7

	// Remove 1 (index 3), last element (7) moves to index 3
	// 7 > its parent (9)? Yes! siftUp is called.
	removed := pq.removeElement(1)
	if !removed {
		t.Errorf("Expected removeElement to return true")
	}

	// After removal, heap should be: [10, 9, 8, 7]
	expected := []int{10, 9, 8, 7}
	actual := pq.ToArray()
	for _, v := range expected {
		found := false
		for _, a := range actual {
			if a == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected value %d in heap, but not found", v)
		}
	}
}

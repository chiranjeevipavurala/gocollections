package lists

import (
	"testing"

	errcodes "github.com/chiranjeevipavurala/gocollections/errors"
	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()

	result := stack.Push(1)
	if !result {
		t.Errorf("Expected Push to return true")
	}
	result = stack.Push(2)
	if !result {
		t.Errorf("Expected Push to return true")
	}
	result = stack.Push(3)
	if !result {
		t.Errorf("Expected Push to return true")
	}

	if stack.IsEmpty() {
		t.Errorf("Expected stack not to be empty")
	}
}

func TestStack_Pop(t *testing.T) {
	// Test popping from empty stack first
	emptyStack := NewStack[int]()
	val, err := emptyStack.Pop()
	if val != nil {
		t.Errorf("Expected nil value when popping from empty stack, got %v", val)
	}
	if err == nil {
		t.Errorf("Expected error when popping from empty stack, got nil")
	}
	if err.Error() != string(errcodes.EmptyStackError) {
		t.Errorf("Expected error to be EmptyStackError, got %v", err)
	}

	// Test normal pop operations
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected popped value to be 3, got %d", *val)
	}

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected popped value to be 2, got %d", *val)
	}

	val, err = stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected popped value to be 1, got %d", *val)
	}

	// Test popping from now-empty stack
	val, err = stack.Pop()
	if val != nil {
		t.Errorf("Expected nil value when popping from empty stack, got %v", val)
	}
	if err == nil {
		t.Errorf("Expected error when popping from empty stack, got nil")
	}
	if err.Error() != string(errcodes.EmptyStackError) {
		t.Errorf("Expected error to be EmptyStackError, got %v", err)
	}

	// Test popping after clearing stack
	stack.Push(1)
	stack.Clear()
	val, err = stack.Pop()
	if val != nil {
		t.Errorf("Expected nil value when popping from cleared stack, got %v", val)
	}
	if err == nil {
		t.Errorf("Expected error when popping from cleared stack, got nil")
	}
	if err.Error() != string(errcodes.EmptyStackError) {
		t.Errorf("Expected error to be EmptyStackError, got %v", err)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()

	// Test peeking empty stack
	val, err := stack.Peek()
	if val != nil {
		t.Errorf("Expected nil value when peeking empty stack, got %v", val)
	}
	if err == nil {
		t.Errorf("Expected error when peeking empty stack, got nil")
	}
	if err.Error() != string(errcodes.EmptyStackError) {
		t.Errorf("Expected error to be EmptyStackError, got %v", err)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, err = stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected peeked value to be 3, got %d", *val)
	}

	stack.Pop()

	val, err = stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected peeked value to be 2, got %d", *val)
	}

	stack.Pop()

	val, err = stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected peeked value to be 1, got %d", *val)
	}

	stack.Pop()

	// Test peeking after last element is popped
	val, err = stack.Peek()
	if val != nil {
		t.Errorf("Expected nil value when peeking empty stack, got %v", val)
	}
	if err == nil {
		t.Errorf("Expected error when peeking empty stack, got nil")
	}
	if err.Error() != string(errcodes.EmptyStackError) {
		t.Errorf("Expected error to be EmptyStackError, got %v", err)
	}

	// Test peeking after clearing stack
	stack.Push(1)
	stack.Clear()
	val, err = stack.Peek()
	if val != nil {
		t.Errorf("Expected nil value when peeking cleared stack, got %v", val)
	}
	if err == nil {
		t.Errorf("Expected error when peeking cleared stack, got nil")
	}
	if err.Error() != string(errcodes.EmptyStackError) {
		t.Errorf("Expected error to be EmptyStackError, got %v", err)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	stack.Push(1)

	if stack.IsEmpty() {
		t.Errorf("Expected stack not to be empty")
	}
}

func TestStack_Search(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Search(1) != 3 {
		t.Errorf("Expected search result to be 3, got %d", stack.Search(1))
	}

	if stack.Search(2) != 2 {
		t.Errorf("Expected search result to be 2, got %d", stack.Search(2))
	}

	if stack.Search(3) != 1 {
		t.Errorf("Expected search result to be 1, got %d", stack.Search(3))
	}

	if stack.Search(4) != -1 {
		t.Errorf("Expected search result to be -1, got %d", stack.Search(4))
	}
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[int]()

	if stack.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", stack.Size())
	}

	stack.Push(1)
	if stack.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", stack.Size())
	}

	stack.Push(2)
	if stack.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", stack.Size())
	}

	stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", stack.Size())
	}
}

func TestStack_Clear(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Clear()

	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after Clear")
	}
	if stack.Size() != 0 {
		t.Errorf("Expected size to be 0 after Clear, got %d", stack.Size())
	}
}

func TestStack_Contains(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if !stack.Contains(1) {
		t.Errorf("Expected stack to contain 1")
	}
	if !stack.Contains(2) {
		t.Errorf("Expected stack to contain 2")
	}
	if !stack.Contains(3) {
		t.Errorf("Expected stack to contain 3")
	}
	if stack.Contains(4) {
		t.Errorf("Expected stack not to contain 4")
	}
}

func TestStack_ToArray(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	arr := stack.ToArray()
	if len(arr) != 3 {
		t.Errorf("Expected array length to be 3, got %d", len(arr))
	}
	if arr[0] != 1 {
		t.Errorf("Expected first element to be 1, got %d", arr[0])
	}
	if arr[1] != 2 {
		t.Errorf("Expected second element to be 2, got %d", arr[1])
	}
	if arr[2] != 3 {
		t.Errorf("Expected third element to be 3, got %d", arr[2])
	}
}

func TestStack_Clone(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	clone := stack.Clone()

	if clone.Size() != stack.Size() {
		t.Errorf("Expected clone size to be %d, got %d", stack.Size(), clone.Size())
	}

	// Modify original stack
	stack.Pop()

	// Check that clone is not affected
	if clone.Size() != 3 {
		t.Errorf("Expected clone size to remain 3, got %d", clone.Size())
	}
}

func TestStack_AddAll(t *testing.T) {
	stack := NewStack[int]()
	otherList := NewArrayList[int]()

	otherList.Add(1)
	otherList.Add(2)
	otherList.Add(3)

	result := stack.AddAll(otherList)
	if !result {
		t.Errorf("Expected AddAll to return true")
	}
	if stack.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", stack.Size())
	}
}

func TestStack_RemoveAll(t *testing.T) {
	stack := NewStack[int]()
	otherList := NewArrayList[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	otherList.Add(2)
	otherList.Add(4)

	result := stack.RemoveAll(otherList)
	if !result {
		t.Errorf("Expected RemoveAll to return true")
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", stack.Size())
	}
	if !stack.Contains(1) || !stack.Contains(3) {
		t.Errorf("Expected stack to contain 1 and 3")
	}
}

func TestStack_RetainAll(t *testing.T) {
	stack := NewStack[int]()
	otherList := NewArrayList[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	otherList.Add(2)
	otherList.Add(4)

	result := stack.RetainAll(otherList)
	if !result {
		t.Errorf("Expected RetainAll to return true")
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", stack.Size())
	}
	if !stack.Contains(2) || !stack.Contains(4) {
		t.Errorf("Expected stack to contain 2 and 4")
	}
}

func TestStack_Equals(t *testing.T) {
	// Test equal stacks
	stack1 := NewStack[int]()
	stack2 := NewStack[int]()

	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)

	stack2.Push(1)
	stack2.Push(2)
	stack2.Push(3)

	assert.True(t, stack1.Equals(stack2))
	assert.True(t, stack2.Equals(stack1))

	// Test different sizes
	stack3 := NewStack[int]()
	stack3.Push(1)
	stack3.Push(2)
	assert.False(t, stack1.Equals(stack3))

	// Test different elements
	stack4 := NewStack[int]()
	stack4.Push(1)
	stack4.Push(2)
	stack4.Push(4)
	assert.False(t, stack1.Equals(stack4))

	// Test empty stacks
	emptyStack1 := NewStack[int]()
	emptyStack2 := NewStack[int]()
	assert.True(t, emptyStack1.Equals(emptyStack2))

	// Test nil collection
	assert.False(t, stack1.Equals(nil))

	// Test with different collection types
	arrayList := NewArrayList[int]()
	arrayList.Add(1)
	arrayList.Add(2)
	arrayList.Add(3)
	assert.True(t, stack1.Equals(arrayList))

	// Test with different order
	stack5 := NewStack[int]()
	stack5.Push(3)
	stack5.Push(2)
	stack5.Push(1)
	assert.False(t, stack1.Equals(stack5))

	// Test with duplicate elements
	stack6 := NewStack[int]()
	stack6.Push(1)
	stack6.Push(1)
	stack6.Push(1)
	stack7 := NewStack[int]()
	stack7.Push(1)
	stack7.Push(1)
	stack7.Push(1)
	assert.True(t, stack6.Equals(stack7))

	// Test with different types
	stringStack := NewStack[string]()
	stringStack.Push("1")
	stringStack.Push("2")
	stringStack.Push("3")

	// Instead of comparing different types, verify that stringStack contains the expected elements
	assert.Equal(t, 3, stringStack.Size())
	assert.True(t, stringStack.Contains("1"))
	assert.True(t, stringStack.Contains("2"))
	assert.True(t, stringStack.Contains("3"))
}

func TestStack_ForEach(t *testing.T) {
	stack := NewStack[int]()
	sum := 0

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.ForEach(func(element int) {
		sum += element
	})

	if sum != 6 {
		t.Errorf("Expected sum to be 6, got %d", sum)
	}
}

func TestStack_Filter(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	filtered := stack.Filter(func(element int) bool {
		return element%2 == 0
	})

	if filtered.Size() != 2 {
		t.Errorf("Expected filtered size to be 2, got %d", filtered.Size())
	}
	if !filtered.Contains(2) || !filtered.Contains(4) {
		t.Errorf("Expected filtered stack to contain 2 and 4")
	}
}

func TestStack_FindFirst(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	element := stack.FindFirst(func(element int) bool {
		return element > 3
	})

	if element == nil {
		t.Errorf("Expected to find element > 3")
	}
	if *element != 4 {
		t.Errorf("Expected to find 4, got %d", *element)
	}

	element = stack.FindFirst(func(element int) bool {
		return element > 5
	})

	if element != nil {
		t.Errorf("Expected not to find element > 5")
	}
}

func TestStack_FindAll(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	elements := stack.FindAll(func(element int) bool {
		return element%2 == 0
	})

	if len(elements) != 2 {
		t.Errorf("Expected to find 2 elements, got %d", len(elements))
	}
	if elements[0] != 2 || elements[1] != 4 {
		t.Errorf("Expected to find 2 and 4, got %v", elements)
	}
}

func TestStack_RemoveIf(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	result := stack.RemoveIf(func(element int) bool {
		return element%2 == 0
	})

	if !result {
		t.Errorf("Expected RemoveIf to return true")
	}
	if stack.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", stack.Size())
	}
	if !stack.Contains(1) || !stack.Contains(3) || !stack.Contains(5) {
		t.Errorf("Expected stack to contain 1, 3, and 5")
	}
}

func TestStack_ReplaceAll(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.ReplaceAll(func(element int) int {
		return element * 2
	})

	arr := stack.ToArray()
	if arr[0] != 2 {
		t.Errorf("Expected first element to be 2, got %d", arr[0])
	}
	if arr[1] != 4 {
		t.Errorf("Expected second element to be 4, got %d", arr[1])
	}
	if arr[2] != 6 {
		t.Errorf("Expected third element to be 6, got %d", arr[2])
	}
}

func TestStack_Remove(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(2) // Add duplicate

	// Test removing non-existent element
	if stack.Remove(4) {
		t.Errorf("Expected Remove to return false for non-existent element")
	}

	// Test removing existing element
	if !stack.Remove(2) {
		t.Errorf("Expected Remove to return true for existing element")
	}
	if stack.Size() != 3 {
		t.Errorf("Expected size to be 3 after removing one element, got %d", stack.Size())
	}

	// Test removing another instance of the same element
	if !stack.Remove(2) {
		t.Errorf("Expected Remove to return true for second instance of element")
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size to be 2 after removing second instance, got %d", stack.Size())
	}

	// Test removing last element
	if !stack.Remove(1) {
		t.Errorf("Expected Remove to return true for last element")
	}
	if stack.Size() != 1 {
		t.Errorf("Expected size to be 1 after removing last element, got %d", stack.Size())
	}
}

func TestStack_ContainsAll(t *testing.T) {
	stack := NewStack[int]()
	otherList := NewArrayList[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test with empty collection
	result, err := stack.ContainsAll(otherList)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected ContainsAll to return true for empty collection")
	}

	// Test with subset of elements
	otherList.Add(1)
	otherList.Add(2)
	result, err = stack.ContainsAll(otherList)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !result {
		t.Errorf("Expected ContainsAll to return true for subset of elements")
	}

	// Test with non-subset
	otherList.Add(4)
	result, err = stack.ContainsAll(otherList)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		t.Errorf("Expected ContainsAll to return false for non-subset")
	}

	// Test with nil collection
	_, err = stack.ContainsAll(nil)
	if err == nil {
		t.Errorf("Expected error for nil collection")
	}
}

func TestStack_Add(t *testing.T) {
	stack := NewStack[int]()

	// Test adding first element
	if !stack.Add(1) {
		t.Errorf("Expected Add to return true for first element")
	}
	if stack.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", stack.Size())
	}
	if !stack.Contains(1) {
		t.Errorf("Expected stack to contain 1")
	}

	// Test adding multiple elements
	if !stack.Add(2) {
		t.Errorf("Expected Add to return true for second element")
	}
	if !stack.Add(3) {
		t.Errorf("Expected Add to return true for third element")
	}
	if stack.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", stack.Size())
	}

	// Verify elements are in correct order (LIFO)
	arr := stack.ToArray()
	if arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Errorf("Expected elements in order [1,2,3], got %v", arr)
	}
}

func TestStack_Iterator(t *testing.T) {
	stack := NewStack[int]()

	// Test empty stack iterator
	iterator := stack.Iterator()
	if iterator.HasNext() {
		t.Errorf("Expected empty stack iterator to have no next elements")
	}
	_, err := iterator.Next()
	if err == nil {
		t.Errorf("Expected error when calling Next on empty stack iterator")
	}

	// Add elements to stack
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)

	// Test iterator with elements
	iterator = stack.Iterator()

	// Verify first element
	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next element")
	}
	val, err := iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 1 {
		t.Errorf("Expected first element to be 1, got %d", *val)
	}

	// Verify second element
	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next element")
	}
	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 2 {
		t.Errorf("Expected second element to be 2, got %d", *val)
	}

	// Verify third element
	if !iterator.HasNext() {
		t.Errorf("Expected iterator to have next element")
	}
	val, err = iterator.Next()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if *val != 3 {
		t.Errorf("Expected third element to be 3, got %d", *val)
	}

	// Verify no more elements
	if iterator.HasNext() {
		t.Errorf("Expected iterator to have no more elements")
	}
	_, err = iterator.Next()
	if err == nil {
		t.Errorf("Expected error when calling Next after all elements")
	}

	// Test concurrent modification
	stack.Add(4)
	if iterator.HasNext() {
		t.Errorf("Expected iterator to be invalidated after stack modification")
	}
}

func TestAddAllNilCollection(t *testing.T) {
	t.Parallel()
	s := NewStack[int]()

	// Test with nil collection
	modified := s.AddAll(nil)
	if modified {
		t.Error("AddAll(nil) should return false")
	}
	if s.Size() != 0 {
		t.Error("AddAll(nil) should not modify the stack")
	}

	// Test with empty collection
	emptyCollection := NewArrayList[int]()
	modified = s.AddAll(emptyCollection)
	if modified {
		t.Error("AddAll(empty) should return false")
	}
	if s.Size() != 0 {
		t.Error("AddAll(empty) should not modify the stack")
	}
}

func TestAddAllWithCollection(t *testing.T) {
	t.Parallel()
	s := NewStack[int]()
	// Create a collection with some elements
	collection := NewArrayList[int]()
	collection.Add(1)
	collection.Add(2)
	collection.Add(3)

	modified := s.AddAll(collection)
	if !modified {
		t.Error("AddAll should return true when collection is not empty")
	}
	if s.Size() != 3 {
		t.Errorf("Stack size should be 3, got %d", s.Size())
	}
	// Verify elements were added in correct order
	expected := []int{1, 2, 3}
	actual := s.ToArray()
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %d at position %d, got %d", v, i, actual[i])
		}
	}
}

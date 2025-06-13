package sets

import (
	"sync"
	"testing"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/stretchr/testify/assert"
)

func TestNewLinkedHashSet(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	assert.NotNil(t, lhs, "NewLinkedHashSet should not return nil")
	assert.True(t, lhs.IsEmpty(), "New LinkedHashSet should be empty")
}

func TestLinkedHashSet_Add(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	assert.True(t, lhs.Add(1), "Add should return true for new element")
	assert.True(t, lhs.Contains(1), "Set should contain added element")
	assert.False(t, lhs.Add(1), "Add should return false for duplicate element")
}

func TestLinkedHashSet_AddAll(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	other := NewLinkedHashSet[int]()
	other.Add(1)
	other.Add(2)
	assert.True(t, lhs.AddAll(other), "AddAll should return true if set is modified")
	assert.True(t, lhs.Contains(1), "Set should contain elements from other set")
	assert.True(t, lhs.Contains(2), "Set should contain elements from other set")
	assert.False(t, lhs.AddAll(other), "AddAll should return false if no new elements are added")
}

func TestLinkedHashSet_Clear(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	lhs.Add(1)
	lhs.Add(2)
	lhs.Clear()
	assert.True(t, lhs.IsEmpty(), "Set should be empty after Clear")
	assert.False(t, lhs.Contains(1), "Set should not contain cleared elements")
}

func TestLinkedHashSet_Contains(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	lhs.Add(1)
	assert.True(t, lhs.Contains(1), "Set should contain added element")
	assert.False(t, lhs.Contains(2), "Set should not contain non-added element")
}

func TestLinkedHashSet_ContainsAll(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	lhs.Add(1)
	lhs.Add(2)
	other := NewLinkedHashSet[int]()
	other.Add(1)
	other.Add(2)
	containsAll, err := lhs.ContainsAll(other)
	assert.NoError(t, err, "ContainsAll should not return error")
	assert.True(t, containsAll, "Set should contain all elements from other set")
	other.Add(3)
	containsAll, err = lhs.ContainsAll(other)
	assert.NoError(t, err, "ContainsAll should not return error")
	assert.False(t, containsAll, "Set should not contain all elements from other set")
}

func TestLinkedHashSet_Equals(t *testing.T) {
	// Test with nil input
	lhs := NewLinkedHashSet[int]()
	assert.False(t, lhs.Equals(nil), "Equals should return false for nil input")

	// Test with different type
	var otherCollection collections.Collection[int] = NewLinkedHashSet[int]()
	assert.True(t, lhs.Equals(otherCollection), "Equals should return true for same type")

	// Test with empty sets
	other := NewLinkedHashSet[int]()
	assert.True(t, lhs.Equals(other), "Empty sets should be equal")

	// Test with different sizes
	lhs.Add(1)
	assert.False(t, lhs.Equals(other), "Sets with different sizes should not be equal")

	// Test with same content
	other.Add(1)
	assert.True(t, lhs.Equals(other), "Sets with same content should be equal")

	// Test with different content
	other.Clear()
	other.Add(2)
	assert.False(t, lhs.Equals(other), "Sets with different content should not be equal")

	// Test with multiple entries
	lhs.Clear()
	lhs.Add(1)
	lhs.Add(2)
	lhs.Add(3)
	other.Clear()
	other.Add(1)
	other.Add(2)
	other.Add(3)
	assert.True(t, lhs.Equals(other), "Sets with multiple same entries should be equal")

	// Test with different order (should still be equal for LinkedHashSet)
	other.Clear()
	other.Add(3)
	other.Add(1)
	other.Add(2)
	assert.True(t, lhs.Equals(other), "Sets with same elements in different order should be equal for LinkedHashSet")

	// Test with empty string values
	lhs.Clear()
	other.Clear()
	lhs.Add(1)
	other.Add(1)
	assert.True(t, lhs.Equals(other), "Sets with empty string values should be equal")

	// Test with custom type
	type customType struct {
		value int
	}
	lhsCustom := NewLinkedHashSet[customType]()
	otherCustom := NewLinkedHashSet[customType]()

	// Test with same custom type values
	val1 := customType{value: 1}
	val2 := customType{value: 2}
	lhsCustom.Add(val1)
	lhsCustom.Add(val2)
	otherCustom.Add(val1)
	otherCustom.Add(val2)
	assert.True(t, lhsCustom.Equals(otherCustom), "Sets with same custom type values should be equal")

	// Test with different custom type values
	otherCustom.Clear()
	otherCustom.Add(customType{value: 3})
	assert.False(t, lhsCustom.Equals(otherCustom), "Sets with different custom type values should not be equal")
}

func TestLinkedHashSet_IsEmpty(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	assert.True(t, lhs.IsEmpty(), "New set should be empty")
	lhs.Add(1)
	assert.False(t, lhs.IsEmpty(), "Set should not be empty after adding element")
	lhs.Clear()
	assert.True(t, lhs.IsEmpty(), "Set should be empty after Clear")
}

func TestLinkedHashSet_Iterator(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	lhs.Add(1)
	lhs.Add(2)
	lhs.Add(3)
	it := lhs.Iterator()
	assert.True(t, it.HasNext(), "Iterator should have next element")
	val, err := it.Next()
	assert.NoError(t, err, "Iterator should not return error")
	assert.Equal(t, 1, *val, "Iterator should return elements in insertion order")
	val, err = it.Next()
	assert.NoError(t, err, "Iterator should not return error")
	assert.Equal(t, 2, *val, "Iterator should return elements in insertion order")
	val, err = it.Next()
	assert.NoError(t, err, "Iterator should not return error")
	assert.Equal(t, 3, *val, "Iterator should return elements in insertion order")
	assert.False(t, it.HasNext(), "Iterator should not have more elements")
	_, err = it.Next()
	assert.Error(t, err, "Iterator should return error when no more elements")
}

func TestLinkedHashSet_Remove(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	lhs.Add(1)
	lhs.Add(2)
	assert.True(t, lhs.Remove(1), "Remove should return true for existing element")
	assert.False(t, lhs.Contains(1), "Set should not contain removed element")
	assert.False(t, lhs.Remove(1), "Remove should return false for non-existing element")
}

func TestLinkedHashSet_RemoveAll(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	lhs.Add(1)
	lhs.Add(2)
	lhs.Add(3)
	other := NewLinkedHashSet[int]()
	other.Add(1)
	other.Add(2)
	assert.True(t, lhs.RemoveAll(other), "RemoveAll should return true if set is modified")
	assert.False(t, lhs.Contains(1), "Set should not contain removed elements")
	assert.False(t, lhs.Contains(2), "Set should not contain removed elements")
	assert.True(t, lhs.Contains(3), "Set should contain non-removed elements")
}

func TestLinkedHashSet_RemoveAll_EdgeCases(t *testing.T) {
	// Test removing from empty set
	lhs := NewLinkedHashSet[int]()
	other := NewLinkedHashSet[int]()
	other.Add(1)
	other.Add(2)
	assert.False(t, lhs.RemoveAll(other), "RemoveAll should return false when removing from empty set")
	assert.True(t, lhs.IsEmpty(), "Empty set should remain empty after RemoveAll")

	// Test removing non-existent elements
	lhs.Add(3)
	lhs.Add(4)
	assert.False(t, lhs.RemoveAll(other), "RemoveAll should return false when removing non-existent elements")
	assert.Equal(t, 2, lhs.Size(), "Set size should remain unchanged when removing non-existent elements")
	assert.True(t, lhs.Contains(3), "Set should retain its elements when removing non-existent elements")
	assert.True(t, lhs.Contains(4), "Set should retain its elements when removing non-existent elements")

	// Test removing from nil collection
	assert.False(t, lhs.RemoveAll(nil), "RemoveAll should return false for nil collection")
	assert.Equal(t, 2, lhs.Size(), "Set size should remain unchanged when removing from nil collection")

	// Test removing all elements
	other.Clear()
	other.Add(3)
	other.Add(4)
	assert.True(t, lhs.RemoveAll(other), "RemoveAll should return true when removing all elements")
	assert.True(t, lhs.IsEmpty(), "Set should be empty after removing all elements")

	// Test removing some elements
	lhs.Add(1)
	lhs.Add(2)
	lhs.Add(3)
	other.Clear()
	other.Add(2)
	other.Add(4) // non-existent element
	assert.True(t, lhs.RemoveAll(other), "RemoveAll should return true when removing some elements")
	assert.Equal(t, 2, lhs.Size(), "Set should have correct size after partial removal")
	assert.True(t, lhs.Contains(1), "Set should retain non-removed elements")
	assert.True(t, lhs.Contains(3), "Set should retain non-removed elements")
	assert.False(t, lhs.Contains(2), "Set should not contain removed elements")
}

func TestLinkedHashSet_Size(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	assert.Equal(t, 0, lhs.Size(), "New set should have size 0")
	lhs.Add(1)
	assert.Equal(t, 1, lhs.Size(), "Set should have size 1 after adding element")
	lhs.Add(2)
	assert.Equal(t, 2, lhs.Size(), "Set should have size 2 after adding element")
	lhs.Remove(1)
	assert.Equal(t, 1, lhs.Size(), "Set should have size 1 after removing element")
}

func TestLinkedHashSet_ToArray(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	lhs.Add(1)
	lhs.Add(2)
	lhs.Add(3)
	arr := lhs.ToArray()
	assert.Equal(t, 3, len(arr), "ToArray should return correct length")
	assert.Equal(t, 1, arr[0], "ToArray should return elements in insertion order")
	assert.Equal(t, 2, arr[1], "ToArray should return elements in insertion order")
	assert.Equal(t, 3, arr[2], "ToArray should return elements in insertion order")
}

func TestLinkedHashSet_Concurrent(t *testing.T) {
	lhs := NewLinkedHashSet[int]()
	var wg sync.WaitGroup
	iterations := 1000
	goroutines := 10
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := id*iterations + j
				lhs.Add(key)
				lhs.Contains(key)
				lhs.Remove(key)
			}
		}(i)
	}
	wg.Wait()
	assert.True(t, lhs.IsEmpty(), "Set should be empty after concurrent operations")
}

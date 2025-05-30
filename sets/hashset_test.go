package sets

import (
	"testing"

	"github.com/chiranjeevipavurala/gocollections/lists"
)

func TestNewHashSet(t *testing.T) {
	// Test default constructor
	set := NewHashSet[int]()
	if set == nil {
		t.Error("NewHashSet returned nil")
	}
	if !set.IsEmpty() {
		t.Error("New HashSet should be empty")
	}

	// Test constructor with capacity
	set = NewHashSetWithCapacity[int](20)
	if set == nil {
		t.Error("NewHashSetWithCapacity returned nil")
	}
	if !set.IsEmpty() {
		t.Error("New HashSet should be empty")
	}

	// Test constructor with negative capacity
	set = NewHashSetWithCapacity[int](-1)
	if set == nil {
		t.Error("NewHashSetWithCapacity returned nil")
	}
	if !set.IsEmpty() {
		t.Error("New HashSet should be empty")
	}
}

func TestHashSet_Add(t *testing.T) {
	set := NewHashSet[int]()

	// Test adding new element
	if !set.Add(1) {
		t.Error("Add(1) should return true")
	}
	if !set.Contains(1) {
		t.Error("Set should contain 1")
	}

	// Test adding duplicate element
	if set.Add(1) {
		t.Error("Add(1) again should return false")
	}

	// Test adding multiple elements
	if !set.Add(2) {
		t.Error("Add(2) should return true")
	}
	if !set.Add(3) {
		t.Error("Add(3) should return true")
	}
	if set.Size() != 3 {
		t.Errorf("Size() = %d; want 3", set.Size())
	}
}

func TestHashSet_AddAll(t *testing.T) {
	set := NewHashSet[int]()
	list := lists.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Test adding collection
	if !set.AddAll(list) {
		t.Error("AddAll should return true")
	}
	if set.Size() != 3 {
		t.Errorf("Size() = %d; want 3", set.Size())
	}

	// Test adding nil collection
	if set.AddAll(nil) {
		t.Error("AddAll(nil) should return false")
	}

	// Test adding empty collection
	emptyList := lists.NewArrayList[int]()
	if set.AddAll(emptyList) {
		t.Error("AddAll(empty) should return false")
	}

	// Test adding collection with duplicates
	list.Add(1)
	if set.AddAll(list) {
		t.Error("AddAll with duplicates should return false")
	}
}

func TestHashSet_AddAllBatch(t *testing.T) {
	set := NewHashSet[int]().(*HashSet[int])

	// Test adding batch
	if !set.AddAllBatch([]int{1, 2, 3}) {
		t.Error("AddAllBatch should return true")
	}
	if set.Size() != 3 {
		t.Errorf("Size() = %d; want 3", set.Size())
	}

	// Test adding nil batch
	if set.AddAllBatch(nil) {
		t.Error("AddAllBatch(nil) should return false")
	}

	// Test adding empty batch
	if set.AddAllBatch([]int{}) {
		t.Error("AddAllBatch(empty) should return false")
	}

	// Test adding batch with duplicates
	if set.AddAllBatch([]int{1, 2}) {
		t.Error("AddAllBatch with duplicates should return false")
	}
}

func TestHashSet_Clear(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Clear()
	if !set.IsEmpty() {
		t.Error("Set should be empty after Clear()")
	}
	if set.Size() != 0 {
		t.Errorf("Size() = %d; want 0", set.Size())
	}
}

func TestHashSet_Contains(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)

	// Test containing element
	if !set.Contains(1) {
		t.Error("Set should contain 1")
	}

	// Test not containing element
	if set.Contains(2) {
		t.Error("Set should not contain 2")
	}
}

func TestHashSet_ContainsAll(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	list := lists.NewArrayList[int]()
	list.Add(1)
	list.Add(2)

	// Test containing all elements
	contains, err := set.ContainsAll(list)
	if err != nil {
		t.Errorf("ContainsAll returned error: %v", err)
	}
	if !contains {
		t.Error("Set should contain all elements")
	}

	// Test not containing all elements
	list.Add(4)
	contains, err = set.ContainsAll(list)
	if err != nil {
		t.Errorf("ContainsAll returned error: %v", err)
	}
	if contains {
		t.Error("Set should not contain all elements")
	}

	// Test nil collection
	contains, err = set.ContainsAll(nil)
	if err == nil {
		t.Error("ContainsAll(nil) should return error")
	}
	if contains {
		t.Error("ContainsAll(nil) should return false")
	}

	// Test empty collection
	emptyList := lists.NewArrayList[int]()
	contains, err = set.ContainsAll(emptyList)
	if err != nil {
		t.Errorf("ContainsAll returned error: %v", err)
	}
	if contains {
		t.Error("ContainsAll(empty) should return false")
	}
}

func TestHashSet_Equals(t *testing.T) {
	set1 := NewHashSet[int]()
	set2 := NewHashSet[int]()
	set3 := NewHashSet[int]()

	// Test empty sets
	if !set1.Equals(set2) {
		t.Error("Empty sets should be equal")
	}

	// Test sets with same elements
	set1.Add(1)
	set1.Add(2)
	set2.Add(1)
	set2.Add(2)
	if !set1.Equals(set2) {
		t.Error("Sets with same elements should be equal")
	}

	// Test sets with different elements
	set3.Add(1)
	set3.Add(3)
	if set1.Equals(set3) {
		t.Error("Sets with different elements should not be equal")
	}

	// Test sets with different sizes
	set3.Add(2)
	set3.Add(4)
	if set1.Equals(set3) {
		t.Error("Sets with different sizes should not be equal")
	}

	// Test nil collection
	if set1.Equals(nil) {
		t.Error("Set should not equal nil")
	}
}

func TestHashSet_IsEmpty(t *testing.T) {
	set := NewHashSet[int]()

	// Test empty set
	if !set.IsEmpty() {
		t.Error("New set should be empty")
	}

	// Test non-empty set
	set.Add(1)
	if set.IsEmpty() {
		t.Error("Set with elements should not be empty")
	}

	// Test after clear
	set.Clear()
	if !set.IsEmpty() {
		t.Error("Set should be empty after Clear()")
	}
}

func TestHashSet_Iterator(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test iterator
	iterator := set.Iterator()
	count := 0
	seen := make(map[int]bool)
	for iterator.HasNext() {
		val, err := iterator.Next()
		if err != nil {
			t.Errorf("Next() returned error: %v", err)
		}
		if val == nil {
			t.Error("Next() returned nil value")
		}
		seen[*val] = true
		count++
	}
	if count != 3 {
		t.Errorf("Iterator returned %d elements; want 3", count)
	}
	if !seen[1] || !seen[2] || !seen[3] {
		t.Error("Iterator did not return all elements")
	}

	// Test iterator after end
	if iterator.HasNext() {
		t.Error("HasNext() should return false after end")
	}
	val, err := iterator.Next()
	if err == nil {
		t.Error("Next() should return error after end")
	}
	if val != nil {
		t.Error("Next() should return nil after end")
	}
}

func TestHashSet_Remove(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test removing existing element
	if !set.Remove(2) {
		t.Error("Remove(2) should return true")
	}
	if set.Contains(2) {
		t.Error("Set should not contain 2 after removal")
	}

	// Test removing non-existent element
	if set.Remove(4) {
		t.Error("Remove(4) should return false")
	}

	// Test removing last element
	if !set.Remove(1) {
		t.Error("Remove(1) should return true")
	}
	if !set.Remove(3) {
		t.Error("Remove(3) should return true")
	}
	if !set.IsEmpty() {
		t.Error("Set should be empty after removing all elements")
	}
}

func TestHashSet_RemoveAll(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	list := lists.NewArrayList[int]()
	list.Add(2)
	list.Add(4)

	// Test removing elements
	if !set.RemoveAll(list) {
		t.Error("RemoveAll should return true")
	}
	if set.Size() != 2 {
		t.Errorf("Size() = %d; want 2", set.Size())
	}
	if set.Contains(2) || set.Contains(4) {
		t.Error("Set should not contain removed elements")
	}

	// Test removing nil collection
	if set.RemoveAll(nil) {
		t.Error("RemoveAll(nil) should return false")
	}

	// Test removing empty collection
	emptyList := lists.NewArrayList[int]()
	if set.RemoveAll(emptyList) {
		t.Error("RemoveAll(empty) should return false")
	}

	// Test removing non-existent elements
	list.Clear()
	list.Add(5)
	list.Add(6)
	if set.RemoveAll(list) {
		t.Error("RemoveAll with non-existent elements should return false")
	}
}

func TestHashSet_RemoveAllBatch(t *testing.T) {
	set := NewHashSet[int]().(*HashSet[int])
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	// Test removing batch
	if !set.RemoveAllBatch([]int{2, 4}) {
		t.Error("RemoveAllBatch should return true")
	}
	if set.Size() != 2 {
		t.Errorf("Size() = %d; want 2", set.Size())
	}
	if set.Contains(2) || set.Contains(4) {
		t.Error("Set should not contain removed elements")
	}

	// Test removing nil batch
	if set.RemoveAllBatch(nil) {
		t.Error("RemoveAllBatch(nil) should return false")
	}

	// Test removing empty batch
	if set.RemoveAllBatch([]int{}) {
		t.Error("RemoveAllBatch(empty) should return false")
	}

	// Test removing non-existent elements
	if set.RemoveAllBatch([]int{5, 6}) {
		t.Error("RemoveAllBatch with non-existent elements should return false")
	}
}

func TestHashSet_RetainAll(t *testing.T) {
	set := NewHashSet[int]().(*HashSet[int])
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	list := lists.NewArrayList[int]()
	list.Add(2)
	list.Add(4)

	// Test retaining elements
	if !set.RetainAll(list) {
		t.Error("RetainAll should return true")
	}
	if set.Size() != 2 {
		t.Errorf("Size() = %d; want 2", set.Size())
	}
	if !set.Contains(2) || !set.Contains(4) {
		t.Error("Set should contain retained elements")
	}
	if set.Contains(1) || set.Contains(3) {
		t.Error("Set should not contain non-retained elements")
	}

	// Test retaining nil collection
	if set.RetainAll(nil) {
		t.Error("RetainAll(nil) should return false")
	}

	// Test retaining empty collection
	emptyList := lists.NewArrayList[int]()
	if !set.RetainAll(emptyList) {
		t.Error("RetainAll(empty) should return true")
	}
	if !set.IsEmpty() {
		t.Error("Set should be empty after retaining empty collection")
	}

	// Test retaining all elements
	set.Add(1)
	set.Add(2)
	set.Add(3)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if set.RetainAll(list) {
		t.Error("RetainAll with all elements should return false")
	}

	// Test retaining empty collection from empty set
	set.Clear()
	if set.RetainAll(emptyList) {
		t.Error("RetainAll(empty) on empty set should return false")
	}
	if !set.IsEmpty() {
		t.Error("Empty set should remain empty after retaining empty collection")
	}
}

func TestHashSet_Size(t *testing.T) {
	set := NewHashSet[int]()

	// Test empty set
	if set.Size() != 0 {
		t.Errorf("Size() = %d; want 0", set.Size())
	}

	// Test after adding elements
	set.Add(1)
	set.Add(2)
	set.Add(3)
	if set.Size() != 3 {
		t.Errorf("Size() = %d; want 3", set.Size())
	}

	// Test after removing elements
	set.Remove(2)
	if set.Size() != 2 {
		t.Errorf("Size() = %d; want 2", set.Size())
	}

	// Test after clear
	set.Clear()
	if set.Size() != 0 {
		t.Errorf("Size() = %d; want 0", set.Size())
	}
}

func TestHashSet_ToArray(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test array conversion
	array := set.ToArray()
	if len(array) != 3 {
		t.Errorf("ToArray() length = %d; want 3", len(array))
	}
	seen := make(map[int]bool)
	for _, val := range array {
		seen[val] = true
	}
	if !seen[1] || !seen[2] || !seen[3] {
		t.Error("ToArray() did not return all elements")
	}

	// Test empty set
	set.Clear()
	array = set.ToArray()
	if len(array) != 0 {
		t.Errorf("ToArray() length = %d; want 0", len(array))
	}
}

func TestHashSet_RemoveIf(t *testing.T) {
	set := NewHashSet[int]().(*HashSet[int])
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	// Test removing even numbers
	if !set.RemoveIf(func(x int) bool { return x%2 == 0 }) {
		t.Error("RemoveIf should return true")
	}
	if set.Size() != 2 {
		t.Errorf("Size() = %d; want 2", set.Size())
	}
	if set.Contains(2) || set.Contains(4) {
		t.Error("Set should not contain removed elements")
	}
	if !set.Contains(1) || !set.Contains(3) {
		t.Error("Set should contain non-removed elements")
	}

	// Test nil predicate
	if set.RemoveIf(nil) {
		t.Error("RemoveIf(nil) should return false")
	}

	// Test removing all elements
	if !set.RemoveIf(func(x int) bool { return true }) {
		t.Error("RemoveIf(all) should return true")
	}
	if !set.IsEmpty() {
		t.Error("Set should be empty after removing all elements")
	}

	// Test removing no elements
	set.Add(1)
	set.Add(2)
	if set.RemoveIf(func(x int) bool { return x > 10 }) {
		t.Error("RemoveIf(none) should return false")
	}
	if set.Size() != 2 {
		t.Errorf("Size() = %d; want 2", set.Size())
	}
}

func TestHashSet_ForEach(t *testing.T) {
	set := NewHashSet[int]().(*HashSet[int])
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test forEach
	sum := 0
	set.ForEach(func(x int) { sum += x })
	if sum != 6 {
		t.Errorf("ForEach sum = %d; want 6", sum)
	}

	// Test nil action
	set.ForEach(nil) // Should not panic

	// Test empty set
	set.Clear()
	sum = 0
	set.ForEach(func(x int) { sum += x })
	if sum != 0 {
		t.Errorf("ForEach sum = %d; want 0", sum)
	}
}

func TestHashSet_Clone(t *testing.T) {
	set := NewHashSet[int]().(*HashSet[int])
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test clone
	clone := set.Clone()
	if clone == nil {
		t.Error("Clone() returned nil")
	}
	if clone == set {
		t.Error("Clone() should return new instance")
	}
	if !clone.Equals(set) {
		t.Error("Clone() should have same elements")
	}

	// Test independence
	clone.Add(4)
	if set.Contains(4) {
		t.Error("Original set should not be affected by clone")
	}
	set.Add(5)
	if clone.Contains(5) {
		t.Error("Clone should not be affected by original set")
	}

	// Test empty set
	set.Clear()
	clone = set.Clone()
	if !clone.IsEmpty() {
		t.Error("Clone of empty set should be empty")
	}
}

func TestNewHashSetFromCollection(t *testing.T) {
	// Test nil collection
	set := NewHashSetFromCollection[int](nil)
	if set == nil {
		t.Error("NewHashSetFromCollection(nil) returned nil")
	}
	if !set.IsEmpty() {
		t.Error("NewHashSetFromCollection(nil) should return empty set")
	}

	// Test empty collection
	emptyList := lists.NewArrayList[int]()
	set = NewHashSetFromCollection[int](emptyList)
	if set == nil {
		t.Error("NewHashSetFromCollection(empty) returned nil")
	}
	if !set.IsEmpty() {
		t.Error("NewHashSetFromCollection(empty) should return empty set")
	}

	// Test collection with elements
	list := lists.NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	set = NewHashSetFromCollection[int](list)
	if set == nil {
		t.Error("NewHashSetFromCollection returned nil")
	}
	if set.Size() != 3 {
		t.Errorf("Size() = %d; want 3", set.Size())
	}
	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Error("Set should contain all elements from collection")
	}
}

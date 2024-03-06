package sets

import (
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if !set.Contains(1) {
		t.Errorf("Expected set to contain element 1")
	}

	if !set.Contains(2) {
		t.Errorf("Expected set to contain element 2")
	}

	if !set.Contains(3) {
		t.Errorf("Expected set to contain element 3")
	}

	if set.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", set.Size())
	}

	set.Add(1)

	if set.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", set.Size())
	}

}

func TestHashSet_AddAll(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	otherSet := NewHashSet[int]()
	otherSet.Add(4)
	otherSet.Add(5)
	otherSet.Add(6)

	set.AddAll(otherSet)

	if !set.Contains(4) {
		t.Errorf("Expected set to contain element 4")
	}

	if !set.Contains(5) {
		t.Errorf("Expected set to contain element 5")
	}

	if !set.Contains(6) {
		t.Errorf("Expected set to contain element 6")
	}

	if set.Size() != 6 {
		t.Errorf("Expected size to be 6, got %d", set.Size())
	}
}

func TestHashSet_Remove(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Remove(2)

	if set.Contains(2) {
		t.Errorf("Expected set not to contain element 2")
	}

	if set.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", set.Size())
	}

	if set.Remove(5) {
		t.Errorf("Expected false")
	}

}

func TestHashSet_RemoveAll(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	otherSet := NewHashSet[int]()
	otherSet.Add(2)
	otherSet.Add(3)

	set.RemoveAll(otherSet)

	if set.Contains(2) {
		t.Errorf("Expected set not to contain element 2")
	}

	if set.Contains(3) {
		t.Errorf("Expected set not to contain element 3")
	}

	if set.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", set.Size())
	}

	if otherSet.RemoveAll(nil) {
		t.Errorf("Expected false")
	}

	emptySet := NewHashSet[int]()
	if set.RemoveAll(emptySet) {
		t.Errorf("Expected false")
	}

}

func TestHashSet_Contains(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if !set.Contains(1) {
		t.Errorf("Expected set to contain element 1")
	}

	if !set.Contains(2) {
		t.Errorf("Expected set to contain element 2")
	}

	if !set.Contains(3) {
		t.Errorf("Expected set to contain element 3")
	}

	if set.Contains(4) {
		t.Errorf("Expected set not to contain element 4")
	}
}

func TestHashSet_ContainsAll(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	otherSet := NewHashSet[int]()
	otherSet.Add(1)
	otherSet.Add(2)

	_, err := set.ContainsAll(otherSet)
	if err != nil {
		t.Errorf("Expected set to contain all elements")
	}

	otherSet.Add(4)

	res, err := set.ContainsAll(otherSet)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if res {
		t.Errorf("Expected set not to contain all elements")
	}

	_, err = set.ContainsAll(nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	emptySet := NewHashSet[int]()
	res, err = set.ContainsAll(emptySet)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if res {
		t.Errorf("Expected set not to contain all elements")
	}

}

func TestHashSet_Equals(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	otherSet := NewHashSet[int]()
	otherSet.Add(1)
	otherSet.Add(2)
	otherSet.Add(3)

	if !set.Equals(otherSet) {
		t.Errorf("Expected sets to be equal")
	}
	set.Add(5)

	otherSet.Add(4)

	if set.Equals(otherSet) {
		t.Errorf("Expected sets not to be equal")
	}

	otherSet.Clear()

	if set.Equals(otherSet) {
		t.Errorf("Expected sets not to be equal")
	}

	if !set.Equals(set) {
		t.Errorf("Expected set to be equal to itself")
	}

	if set.Equals(nil) {
		t.Errorf("Expected sets not to be equal")
	}

}

func TestHashSet_Iterator(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	iterator := set.Iterator()
	for iterator.HasNext() {
		val, _ := iterator.Next()
		if !set.Contains(*val) {
			t.Errorf("Expected set to contain element %d", val)
		}
	}

	val, err := iterator.Next()
	if val != nil {
		t.Errorf("Expected nil, got %d", *val)
	}
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestHashSet_IsEmpty(t *testing.T) {
	set := NewHashSet[int]()
	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}

	set.Add(1)
	if set.IsEmpty() {
		t.Errorf("Expected set not to be empty")
	}
}

func TestHashSet_Size(t *testing.T) {
	set := NewHashSet[int]()
	if set.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", set.Size())
	}

	set.Add(1)
	if set.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", set.Size())
	}

	set.Add(2)
	if set.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", set.Size())
	}
}

func TestHashSet_Clear(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", set.Size())
	}

	set.Clear()

	if set.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", set.Size())
	}
}

func TestHashSet_ToArray(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	arr := set.ToArray()

	if len(arr) != 3 {
		t.Errorf("Expected array length to be 3, got %d", len(arr))
	}

	if !contains(arr, 1) {
		t.Errorf("Expected array to contain element 1")
	}

	if !contains(arr, 2) {
		t.Errorf("Expected array to contain element 2")
	}

	if !contains(arr, 3) {
		t.Errorf("Expected array to contain element 3")
	}
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

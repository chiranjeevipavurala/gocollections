package main

import (
	"testing"

	"github.com/chiranjeevipavurala/gocollections/sets"
)

// HashSet Benchmarks

func BenchmarkHashSetAdd(b *testing.B) {
	set := sets.NewHashSet[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkHashSetContains(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i % MediumSize)
	}
}

func BenchmarkHashSetRemove(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(i % MediumSize)
	}
}

func BenchmarkHashSetAddAll(b *testing.B) {
	set := sets.NewHashSet[int]()
	otherSet := sets.NewHashSet[int]()
	// Pre-populate other set
	for i := 0; i < SmallSize; i++ {
		otherSet.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.AddAll(otherSet)
	}
}

// Additional HashSet benchmarks for missing methods

func BenchmarkHashSetClear(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
		// Re-populate for next iteration
		for j := 0; j < MediumSize; j++ {
			set.Add(j)
		}
	}
}

func BenchmarkHashSetIsEmpty(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.IsEmpty()
	}
}

func BenchmarkHashSetSize(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Size()
	}
}

func BenchmarkHashSetToArray(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.ToArray()
	}
}

func BenchmarkHashSetRemoveAll(b *testing.B) {
	set := sets.NewHashSet[int]()
	otherSet := sets.NewHashSet[int]()
	// Pre-populate both sets
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
		if i%2 == 0 {
			otherSet.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.RemoveAll(otherSet)
	}
}

func BenchmarkHashSetRetainAll(b *testing.B) {
	otherSet := sets.NewHashSet[int]()
	// Pre-populate other set
	for i := 0; i < MediumSize; i++ {
		if i%2 == 0 {
			otherSet.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh set for each iteration
		set := sets.NewHashSet[int]()
		// Pre-populate set
		for j := 0; j < MediumSize; j++ {
			set.Add(j)
		}

		// Type assert to access RetainAll
		if hashSet, ok := set.(*sets.HashSet[int]); ok {
			hashSet.RetainAll(otherSet)
		}
	}
}

func BenchmarkHashSetContainsAll(b *testing.B) {
	set := sets.NewHashSet[int]()
	otherSet := sets.NewHashSet[int]()
	// Pre-populate both sets
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
		if i%2 == 0 {
			otherSet.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.ContainsAll(otherSet)
	}
}

func BenchmarkHashSetEquals(b *testing.B) {
	set := sets.NewHashSet[int]()
	otherSet := sets.NewHashSet[int]()
	// Pre-populate both sets identically
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
		otherSet.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Equals(otherSet)
	}
}

func BenchmarkHashSetClone(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	// Type assert to access Clone
	if hashSet, ok := set.(*sets.HashSet[int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hashSet.Clone()
		}
	}
}

func BenchmarkHashSetRemoveIf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh set for each iteration
		set := sets.NewHashSet[int]()
		// Pre-populate with data
		for j := 0; j < MediumSize; j++ {
			set.Add(j)
		}

		// Type assert to access RemoveIf
		if hashSet, ok := set.(*sets.HashSet[int]); ok {
			hashSet.RemoveIf(func(element int) bool {
				return element%2 == 0
			})
		}
	}
}

func BenchmarkHashSetForEach(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	// Type assert to access ForEach
	if hashSet, ok := set.(*sets.HashSet[int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hashSet.ForEach(func(element int) {
				// Do nothing, just iterate
			})
		}
	}
}

func BenchmarkHashSetAddAllBatch(b *testing.B) {
	elements := make([]int, SmallSize)
	for i := 0; i < SmallSize; i++ {
		elements[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh set for each iteration
		set := sets.NewHashSet[int]()
		// Type assert to access AddAllBatch
		if hashSet, ok := set.(*sets.HashSet[int]); ok {
			hashSet.AddAllBatch(elements)
		}
	}
}

func BenchmarkHashSetRemoveAllBatch(b *testing.B) {
	elements := make([]int, SmallSize)
	// Pre-populate elements array
	for i := 0; i < SmallSize; i++ {
		elements[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh set for each iteration
		set := sets.NewHashSet[int]()
		// Pre-populate set with MediumSize elements
		for j := 0; j < MediumSize; j++ {
			set.Add(j)
		}

		// Type assert to access RemoveAllBatch
		if hashSet, ok := set.(*sets.HashSet[int]); ok {
			hashSet.RemoveAllBatch(elements)
		}
	}
}

// LinkedHashSet Benchmarks

func BenchmarkLinkedHashSetAdd(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkLinkedHashSetContains(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i % MediumSize)
	}
}

func BenchmarkLinkedHashSetRemove(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh set for each iteration
		set := sets.NewLinkedHashSet[int]()
		// Pre-populate with data
		for j := 0; j < MediumSize; j++ {
			set.Add(j)
		}

		set.Remove(i % MediumSize)
	}
}

func BenchmarkLinkedHashSetAddAll(b *testing.B) {
	otherSet := sets.NewLinkedHashSet[int]()
	// Pre-populate other set
	for i := 0; i < SmallSize; i++ {
		otherSet.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh set for each iteration
		set := sets.NewLinkedHashSet[int]()
		set.AddAll(otherSet)
	}
}

// Additional LinkedHashSet benchmarks for missing methods

func BenchmarkLinkedHashSetClear(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
		// Re-populate for next iteration
		for j := 0; j < MediumSize; j++ {
			set.Add(j)
		}
	}
}

func BenchmarkLinkedHashSetIsEmpty(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.IsEmpty()
	}
}

func BenchmarkLinkedHashSetSize(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Size()
	}
}

func BenchmarkLinkedHashSetToArray(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.ToArray()
	}
}

func BenchmarkLinkedHashSetRemoveAll(b *testing.B) {
	otherSet := sets.NewLinkedHashSet[int]()
	// Pre-populate other set
	for i := 0; i < MediumSize; i++ {
		if i%2 == 0 {
			otherSet.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh set for each iteration
		set := sets.NewLinkedHashSet[int]()
		// Pre-populate set
		for j := 0; j < MediumSize; j++ {
			set.Add(j)
		}

		set.RemoveAll(otherSet)
	}
}

func BenchmarkLinkedHashSetContainsAll(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	otherSet := sets.NewLinkedHashSet[int]()
	// Pre-populate both sets
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
		if i%2 == 0 {
			otherSet.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.ContainsAll(otherSet)
	}
}

func BenchmarkLinkedHashSetEquals(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	otherSet := sets.NewLinkedHashSet[int]()
	// Pre-populate both sets identically
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
		otherSet.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Equals(otherSet)
	}
}

// Iterator performance benchmarks

func BenchmarkHashSetIterator(b *testing.B) {
	set := sets.NewHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := set.Iterator()
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}

func BenchmarkLinkedHashSetIterator(b *testing.B) {
	set := sets.NewLinkedHashSet[int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := set.Iterator()
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}

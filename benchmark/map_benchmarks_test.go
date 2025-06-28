package main

import (
	"testing"

	"github.com/chiranjeevipavurala/gocollections/maps"
)

// HashMap Benchmarks

func BenchmarkHashMapPut(b *testing.B) {
	hashMap := maps.NewHashMap[string, int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		hashMap.Put(key, i)
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	hashMap := maps.NewHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		hashMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		hashMap.Get(key)
	}
}

func BenchmarkHashMapRemove(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh map for each iteration
		hashMap := maps.NewHashMap[string, int]()
		// Pre-populate with data
		for j := 0; j < MediumSize; j++ {
			key := string(rune(j%26 + 'a'))
			hashMap.Put(key, j)
		}
		// Remove all elements
		for j := 0; j < MediumSize; j++ {
			key := string(rune(j%26 + 'a'))
			hashMap.Remove(key)
		}
	}
}

func BenchmarkHashMapHasKey(b *testing.B) {
	hashMap := maps.NewHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		hashMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		hashMap.HasKey(key)
	}
}

// LinkedHashMap Benchmarks

func BenchmarkLinkedHashMapPut(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}
}

func BenchmarkLinkedHashMapGet(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Get(key)
	}
}

func BenchmarkLinkedHashMapRemove(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh map for each iteration
		linkedHashMap := maps.NewLinkedHashMap[string, int]()
		// Pre-populate with data
		for j := 0; j < MediumSize; j++ {
			key := string(rune(j%26 + 'a'))
			linkedHashMap.Put(key, j)
		}
		// Remove all elements
		for j := 0; j < MediumSize; j++ {
			key := string(rune(j%26 + 'a'))
			linkedHashMap.Remove(key)
		}
	}
}

func BenchmarkLinkedHashMapHasKey(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.HasKey(key)
	}
}

// LinkedHashMap-specific operations

func BenchmarkLinkedHashMapPutIfAbsent(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.PutIfAbsent(key, i)
	}
}

func BenchmarkLinkedHashMapComputeIfAbsent(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Type assert to access ComputeIfAbsent
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			key := string(rune(i%26 + 'a'))
			lhm.ComputeIfAbsent(key, func(k string) int {
				return len(k) * 10
			})
		}
	}
}

func BenchmarkLinkedHashMapForEachEntry(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	// Type assert to access ForEachEntry
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lhm.ForEachEntry(func(key string, value int) {
				// Do nothing, just iterate
			})
		}
	}
}

// Additional LinkedHashMap benchmarks for missing methods

func BenchmarkLinkedHashMapHasValue(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	// Type assert to access HasValue
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lhm.HasValue(i % MediumSize)
		}
	}
}

func BenchmarkLinkedHashMapEntrySet(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedHashMap.EntrySet()
	}
}

func BenchmarkLinkedHashMapKeySet(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedHashMap.KeySet()
	}
}

func BenchmarkLinkedHashMapValues(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedHashMap.Values()
	}
}

func BenchmarkLinkedHashMapGetOrDefault(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	// Type assert to access GetOrDefault
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := string(rune(i%26 + 'a'))
			lhm.GetOrDefault(key, -1)
		}
	}
}

func BenchmarkLinkedHashMapRemoveKeyWithValue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh map for each iteration
		linkedHashMap := maps.NewLinkedHashMap[string, int]()
		// Pre-populate with data
		for j := 0; j < MediumSize; j++ {
			key := string(rune(j%26 + 'a'))
			linkedHashMap.Put(key, j)
		}

		if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
			// Remove all elements
			for j := 0; j < MediumSize; j++ {
				key := string(rune(j%26 + 'a'))
				lhm.RemoveKeyWithValue(key, j)
			}
		}
	}
}

func BenchmarkLinkedHashMapReplace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh map for each iteration
		linkedHashMap := maps.NewLinkedHashMap[string, int]()
		// Pre-populate with data
		for j := 0; j < MediumSize; j++ {
			key := string(rune(j%26 + 'a'))
			linkedHashMap.Put(key, j)
		}

		if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
			// Replace all elements
			for j := 0; j < MediumSize; j++ {
				key := string(rune(j%26 + 'a'))
				lhm.Replace(key, j*2)
			}
		}
	}
}

func BenchmarkLinkedHashMapReplaceKeyWithValue(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh map for each iteration
		linkedHashMap := maps.NewLinkedHashMap[string, int]()
		// Pre-populate with data
		for j := 0; j < MediumSize; j++ {
			key := string(rune(j%26 + 'a'))
			linkedHashMap.Put(key, j)
		}

		if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
			// Replace all elements
			for j := 0; j < MediumSize; j++ {
				key := string(rune(j%26 + 'a'))
				lhm.ReplaceKeyWithValue(key, j, j*2)
			}
		}
	}
}

func BenchmarkLinkedHashMapClear(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	// Type assert to access Clear
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lhm.Clear()
			// Re-populate for next iteration
			for j := 0; j < MediumSize; j++ {
				key := string(rune(j%26 + 'a'))
				lhm.Put(key, j)
			}
		}
	}
}

func BenchmarkLinkedHashMapIsEmpty(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	// Type assert to access IsEmpty
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lhm.IsEmpty()
		}
	}
}

func BenchmarkLinkedHashMapSize(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
	}

	// Type assert to access Size
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lhm.Size()
		}
	}
}

func BenchmarkLinkedHashMapEquals(b *testing.B) {
	linkedHashMap := maps.NewLinkedHashMap[string, int]()
	otherMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate both maps identically
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		linkedHashMap.Put(key, i)
		otherMap.Put(key, i)
	}

	// Type assert to access Equals
	if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			lhm.Equals(otherMap)
		}
	}
}

func BenchmarkLinkedHashMapPutAll(b *testing.B) {
	otherMap := maps.NewLinkedHashMap[string, int]()
	// Pre-populate other map
	for i := 0; i < SmallSize; i++ {
		key := string(rune(i%26 + 'a'))
		otherMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh map for each iteration
		linkedHashMap := maps.NewLinkedHashMap[string, int]()
		if lhm, ok := linkedHashMap.(*maps.LinkedHashMap[string, int]); ok {
			lhm.PutAll(otherMap)
		}
	}
}

// TreeMap Benchmarks (if available)

func BenchmarkTreeMapPut(b *testing.B) {
	comparator := &StringComparator{}
	treeMap := maps.NewTreeMap[string, int](comparator)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		treeMap.Put(key, i)
	}
}

func BenchmarkTreeMapGet(b *testing.B) {
	comparator := &StringComparator{}
	treeMap := maps.NewTreeMap[string, int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		treeMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		treeMap.Get(key)
	}
}

func BenchmarkTreeMapRemove(b *testing.B) {
	comparator := &StringComparator{}
	treeMap := maps.NewTreeMap[string, int](comparator)
	// Pre-populate with data
	for i := 0; i < MediumSize; i++ {
		key := string(rune(i%26 + 'a'))
		treeMap.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := string(rune(i%26 + 'a'))
		treeMap.Remove(key)
	}
}

// StringComparator for TreeMap
type StringComparator struct{}

func (c *StringComparator) Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

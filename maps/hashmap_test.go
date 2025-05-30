package maps

import (
	"sync"
	"testing"

	"github.com/chiranjeevipavurala/gocollections/collections"
)

func TestNewHashMap(t *testing.T) {
	// Test default constructor
	hm := NewHashMap[string, int]()
	if hm == nil {
		t.Error("NewHashMap returned nil")
	}
	if !hm.IsEmpty() {
		t.Error("New HashMap should be empty")
	}

	// Test constructor with capacity
	hm = NewHashMapWithCapacity[string, int](20)
	if hm == nil {
		t.Error("NewHashMapWithCapacity returned nil")
	}
	if !hm.IsEmpty() {
		t.Error("New HashMap should be empty")
	}

	// Test constructor with invalid capacity
	hm = NewHashMapWithCapacity[string, int](-1)
	if hm == nil {
		t.Error("NewHashMapWithCapacity returned nil")
	}
	if !hm.IsEmpty() {
		t.Error("New HashMap should be empty")
	}
}

func TestHashMap_BasicOperations(t *testing.T) {
	hm := NewHashMap[string, int]()

	// Test Put and Get
	hm.Put("one", 1)
	hm.Put("two", 2)
	hm.Put("three", 3)

	if val := hm.Get("one"); val == nil || *val != 1 {
		t.Errorf("Get('one') = %v; want 1", val)
	}

	// Test Size
	if size := hm.Size(); size != 3 {
		t.Errorf("Size() = %d; want 3", size)
	}

	// Test HasKey
	if !hm.HasKey("one") {
		t.Error("HasKey('one') = false; want true")
	}
	if hm.HasKey("four") {
		t.Error("HasKey('four') = true; want false")
	}

	// Test HasValue
	if !hm.HasValue(1) {
		t.Error("HasValue(1) = false; want true")
	}
	if hm.HasValue(4) {
		t.Error("HasValue(4) = true; want false")
	}

	// Test Remove
	if val := hm.Remove("one"); val != 1 {
		t.Errorf("Remove('one') = %v; want 1", val)
	}
	if hm.HasKey("one") {
		t.Error("HasKey('one') = true after removal; want false")
	}

	// Test Clear
	hm.Clear()
	if !hm.IsEmpty() {
		t.Error("IsEmpty() = false after Clear(); want true")
	}
}

func TestHashMap_EntrySet(t *testing.T) {
	hm := NewHashMap[string, int]()
	hm.Put("one", 1)
	hm.Put("two", 2)

	entrySet := hm.EntrySet()
	if entrySet.Size() != 2 {
		t.Errorf("EntrySet.Size() = %d; want 2", entrySet.Size())
	}

	// Test entry set contents
	entries := entrySet.ToArray()
	found := make(map[string]bool)
	for _, entry := range entries {
		key := entry.GetKey()
		value := entry.GetValue()
		if key == "one" && value == 1 {
			found["one"] = true
		} else if key == "two" && value == 2 {
			found["two"] = true
		}
	}
	if !found["one"] || !found["two"] {
		t.Error("EntrySet missing expected entries")
	}
}

func TestHashMap_KeySet(t *testing.T) {
	hm := NewHashMap[string, int]()
	hm.Put("one", 1)
	hm.Put("two", 2)

	keySet := hm.KeySet()
	if keySet.Size() != 2 {
		t.Errorf("KeySet.Size() = %d; want 2", keySet.Size())
	}

	// Test key set contents
	keys := keySet.ToArray()
	found := make(map[string]bool)
	for _, key := range keys {
		found[key] = true
	}
	if !found["one"] || !found["two"] {
		t.Error("KeySet missing expected keys")
	}
}

func TestHashMap_Values(t *testing.T) {
	hm := NewHashMap[string, int]()
	hm.Put("one", 1)
	hm.Put("two", 2)

	values := hm.Values()
	if values.Size() != 2 {
		t.Errorf("Values.Size() = %d; want 2", values.Size())
	}

	// Test values collection contents
	vals := values.ToArray()
	found := make(map[int]bool)
	for _, val := range vals {
		found[val] = true
	}
	if !found[1] || !found[2] {
		t.Error("Values missing expected values")
	}
}

func TestHashMap_PutIfAbsent(t *testing.T) {
	hm := NewHashMap[string, int]()

	// Test putting new value
	if val := hm.PutIfAbsent("one", 1); val != 0 {
		t.Errorf("PutIfAbsent('one', 1) = %v; want 0", val)
	}

	// Test putting existing value
	if val := hm.PutIfAbsent("one", 2); val != 1 {
		t.Errorf("PutIfAbsent('one', 2) = %v; want 1", val)
	}
}

func TestHashMap_Replace(t *testing.T) {
	hm := NewHashMap[string, int]()
	hm.Put("one", 1)

	// Test replacing existing value
	if val := hm.Replace("one", 2); val != 1 {
		t.Errorf("Replace('one', 2) = %v; want 1", val)
	}

	// Test replacing non-existent value
	if val := hm.Replace("two", 2); val != 0 {
		t.Errorf("Replace('two', 2) = %v; want 0", val)
	}
}

func TestHashMap_ReplaceKeyWithValue(t *testing.T) {
	hm := NewHashMap[string, int]()
	hm.Put("one", 1)

	// Test replacing with correct old value
	if !hm.ReplaceKeyWithValue("one", 1, 2) {
		t.Error("ReplaceKeyWithValue('one', 1, 2) = false; want true")
	}

	// Test replacing with incorrect old value
	if hm.ReplaceKeyWithValue("one", 1, 3) {
		t.Error("ReplaceKeyWithValue('one', 1, 3) = true; want false")
	}

	// Test replacing non-existent key
	if hm.ReplaceKeyWithValue("two", 1, 2) {
		t.Error("ReplaceKeyWithValue('two', 1, 2) = true; want false")
	}
}

func TestHashMap_RemoveKeyWithValue(t *testing.T) {
	hm := NewHashMap[string, int]()
	hm.Put("one", 1)

	// Test removing with correct value
	if !hm.RemoveKeyWithValue("one", 1) {
		t.Error("RemoveKeyWithValue('one', 1) = false; want true")
	}

	// Test removing with incorrect value
	hm.Put("one", 1)
	if hm.RemoveKeyWithValue("one", 2) {
		t.Error("RemoveKeyWithValue('one', 2) = true; want false")
	}

	// Test removing non-existent key
	if hm.RemoveKeyWithValue("two", 1) {
		t.Error("RemoveKeyWithValue('two', 1) = true; want false")
	}
}

func TestHashMap_PutAll(t *testing.T) {
	hm1 := NewHashMap[string, int]()
	hm2 := NewHashMap[string, int]()

	// Test putting into empty map
	hm2.Put("one", 1)
	hm2.Put("two", 2)
	hm1.PutAll(hm2)
	if hm1.Size() != 2 {
		t.Errorf("PutAll size = %d; want 2", hm1.Size())
	}
	if val := hm1.Get("one"); val == nil || *val != 1 {
		t.Errorf("Get('one') = %v; want 1", val)
	}

	// Test putting into non-empty map
	hm3 := NewHashMap[string, int]()
	hm3.Put("three", 3)
	hm1.PutAll(hm3)
	if hm1.Size() != 3 {
		t.Errorf("PutAll size = %d; want 3", hm1.Size())
	}

	// Test putting nil map
	hm1.PutAll(nil)
	if hm1.Size() != 3 {
		t.Errorf("PutAll(nil) size = %d; want 3", hm1.Size())
	}

	// Test putting map with overlapping keys
	hm4 := NewHashMap[string, int]()
	hm4.Put("one", 10)
	hm1.PutAll(hm4)
	if val := hm1.Get("one"); val == nil || *val != 10 {
		t.Errorf("Get('one') after overlap = %v; want 10", val)
	}
}

func TestHashMap_Concurrent(t *testing.T) {
	hm := NewHashMap[string, int]()
	var wg sync.WaitGroup
	iterations := 1000
	goroutines := 10

	// Test concurrent puts and gets
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := "key" + string(rune(id))
				hm.Put(key, j)
				hm.Get(key)
			}
		}(i)
	}
	wg.Wait()

	// Test concurrent removes
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := "key" + string(rune(id))
				hm.Remove(key)
			}
		}(i)
	}
	wg.Wait()
}

func TestHashMap_Equals(t *testing.T) {
	hm1 := NewHashMap[string, int]()
	hm2 := NewHashMap[string, int]()
	hm3 := NewHashMap[string, int]()
	hm4 := NewHashMap[string, int]()

	// Test empty maps
	if !hm1.Equals(hm2) {
		t.Error("Empty maps should be equal")
	}

	// Test maps with same entries
	hm1.Put("one", 1)
	hm1.Put("two", 2)
	hm2.Put("one", 1)
	hm2.Put("two", 2)
	if !hm1.Equals(hm2) {
		t.Error("Maps with same entries should be equal")
	}

	// Test maps with different sizes
	hm3.Put("one", 1)
	if hm1.Equals(hm3) {
		t.Error("Maps with different sizes should not be equal")
	}

	// Test maps with different entries
	hm4.Put("one", 1)
	hm4.Put("two", 3)
	if hm1.Equals(hm4) {
		t.Error("Maps with different entries should not be equal")
	}

	// Test nil map
	if hm1.Equals(nil) {
		t.Error("Map should not equal nil")
	}

	// Test different type
	if hm1.Equals("not a map") {
		t.Error("Map should not equal different type")
	}
}

func TestHashMap_GetOrDefault(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	hm.Put("one", 1)

	if val := hm.GetOrDefault("one", 0); val != 1 {
		t.Errorf("GetOrDefault('one', 0) = %v; want 1", val)
	}
	if val := hm.GetOrDefault("two", 42); val != 42 {
		t.Errorf("GetOrDefault('two', 42) = %v; want 42", val)
	}
}

func TestHashMap_ForEachEntry(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	hm.Put("a", 1)
	hm.Put("b", 2)
	hm.Put("c", 3)
	sum := 0
	hm.ForEachEntry(func(k string, v int) {
		sum += v
	})
	if sum != 6 {
		t.Errorf("ForEachEntry sum = %d; want 6", sum)
	}
	// Edge: empty map
	hm.Clear()
	sum = 0
	hm.ForEachEntry(func(k string, v int) {
		sum += v
	})
	if sum != 0 {
		t.Errorf("ForEachEntry on empty map sum = %d; want 0", sum)
	}
}

func TestHashMap_ComputeIfAbsent(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	val, err := hm.ComputeIfAbsent("foo", func(k string) int { return 100 })
	if err != nil || val != 100 {
		t.Errorf("ComputeIfAbsent('foo') = %v, %v; want 100, nil", val, err)
	}
	val, err = hm.ComputeIfAbsent("foo", func(k string) int { return 200 })
	if err != nil || val != 100 {
		t.Errorf("ComputeIfAbsent('foo') again = %v, %v; want 100, nil", val, err)
	}
	// Edge: nil mappingFunction
	val, err = hm.ComputeIfAbsent("bar", nil)
	if err == nil {
		t.Error("ComputeIfAbsent with nil mappingFunction should return error")
	}
	if val != 0 {
		t.Errorf("ComputeIfAbsent with nil mappingFunction should return zero value, got %v", val)
	}
	if hm.HasKey("bar") {
		t.Error("ComputeIfAbsent with nil mappingFunction should not add key to map")
	}
}

func TestHashMap_PutAllBatch(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	entries := []collections.MapEntry[string, int]{
		collections.NewHashMapEntry("x", 10),
		collections.NewHashMapEntry("y", 20),
	}
	err := hm.PutAllBatch(entries)
	if err != nil {
		t.Errorf("PutAllBatch() = %v; want nil", err)
	}
	if hm.Size() != 2 {
		t.Errorf("Size() = %d; want 2", hm.Size())
	}
	// Edge: nil entries
	err = hm.PutAllBatch(nil)
	if err == nil {
		t.Error("PutAllBatch(nil) = nil; want error")
	}
}

func TestHashMap_RemoveAllBatch(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	hm.Put("a", 1)
	hm.Put("b", 2)
	hm.Put("c", 3)
	err := hm.RemoveAllBatch([]string{"a", "b"})
	if err != nil {
		t.Errorf("RemoveAllBatch() = %v; want nil", err)
	}
	if hm.Size() != 1 {
		t.Errorf("Size() = %d; want 1", hm.Size())
	}
	// Edge: nil keys
	err = hm.RemoveAllBatch(nil)
	if err == nil {
		t.Error("RemoveAllBatch(nil) = nil; want error")
	}
}

func TestHashMap_Clone(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	hm.Put("a", 1)
	hm.Put("b", 2)
	clone := hm.Clone().(*HashMap[string, int])
	if clone.Size() != 2 {
		t.Errorf("Clone.Size() = %d; want 2", clone.Size())
	}
	clone.Put("c", 3)
	if hm.Size() == clone.Size() {
		t.Error("Clone should be independent of original")
	}
	// Edge: clone of empty map
	hm.Clear()
	clone = hm.Clone().(*HashMap[string, int])
	if !clone.IsEmpty() {
		t.Error("Clone of empty map should be empty")
	}
}

func TestHashMap_RemoveIf(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	hm.Put("a", 1)
	hm.Put("b", 2)
	hm.Put("c", 3)
	removed := hm.RemoveIf(func(k string, v int) bool { return v%2 == 0 })
	if !removed {
		t.Error("RemoveIf() = false; want true")
	}
	if hm.Size() != 2 {
		t.Errorf("Size() = %d; want 2", hm.Size())
	}
	// Edge: remove all
	removed = hm.RemoveIf(func(k string, v int) bool { return true })
	if !removed || !hm.IsEmpty() {
		t.Error("RemoveIf(all) did not remove all elements")
	}
	// Edge: nil predicate
	hm.Put("a", 1)
	hm.Put("b", 2)
	removed = hm.RemoveIf(nil)
	if removed {
		t.Error("RemoveIf(nil) = true; want false")
	}
	if hm.Size() != 2 {
		t.Errorf("RemoveIf(nil) should not modify map, size = %d; want 2", hm.Size())
	}
	if val := hm.Get("a"); val == nil || *val != 1 {
		t.Errorf("RemoveIf(nil) should not modify values, got %v; want 1", val)
	}
	if val := hm.Get("b"); val == nil || *val != 2 {
		t.Errorf("RemoveIf(nil) should not modify values, got %v; want 2", val)
	}
}

func TestHashMap_ReplaceAll(t *testing.T) {
	hm := NewHashMap[string, int]().(*HashMap[string, int])
	hm.Put("a", 1)
	hm.Put("b", 2)
	hm.ReplaceAll(func(k string, v int) int { return v * 10 })
	if val := hm.Get("a"); val == nil || *val != 10 {
		t.Errorf("ReplaceAll('a') = %v; want 10", val)
	}
	if val := hm.Get("b"); val == nil || *val != 20 {
		t.Errorf("ReplaceAll('b') = %v; want 20", val)
	}
	// Edge: empty map
	hm.Clear()
	hm.ReplaceAll(func(k string, v int) int { return v * 2 })
	if !hm.IsEmpty() {
		t.Error("ReplaceAll on empty map should keep it empty")
	}
	// Edge: nil operator
	hm.Put("a", 1)
	hm.Put("b", 2)
	hm.ReplaceAll(nil)
	if val := hm.Get("a"); val == nil || *val != 1 {
		t.Errorf("ReplaceAll with nil operator should not modify values, got %v; want 1", val)
	}
	if val := hm.Get("b"); val == nil || *val != 2 {
		t.Errorf("ReplaceAll with nil operator should not modify values, got %v; want 2", val)
	}
}

func TestHashMap_Get(t *testing.T) {
	hm := NewHashMap[string, int]()

	// Test getting non-existent key
	if val := hm.Get("nonexistent"); val != nil {
		t.Errorf("Get('nonexistent') = %v; want nil", val)
	}

	// Test getting existing key
	hm.Put("one", 1)
	if val := hm.Get("one"); val == nil || *val != 1 {
		t.Errorf("Get('one') = %v; want 1", val)
	}

	// Test getting key after removal
	hm.Remove("one")
	if val := hm.Get("one"); val != nil {
		t.Errorf("Get('one') after removal = %v; want nil", val)
	}
}

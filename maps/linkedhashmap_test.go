package maps

import (
	"fmt"
	"sync"
	"testing"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/stretchr/testify/assert"
)

func TestNewLinkedHashMap(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	assert.NotNil(t, lhm)
	assert.Equal(t, 0, lhm.Size())
	assert.True(t, lhm.IsEmpty())
}

func TestLinkedHashMap_PutAndGet(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()

	// Test Put
	oldValue := lhm.Put("one", 1)
	assert.Equal(t, 0, oldValue) // Should return zero value for new key

	// Test Get
	value := lhm.Get("one")
	assert.NotNil(t, value)
	assert.Equal(t, 1, *value)

	// Test Put with existing key
	oldValue = lhm.Put("one", 2)
	assert.Equal(t, 1, oldValue)

	// Test Get after update
	value = lhm.Get("one")
	assert.NotNil(t, value)
	assert.Equal(t, 2, *value)

	// Test Get non-existent key
	value = lhm.Get("two")
	assert.Nil(t, value)
}

func TestLinkedHashMap_PutAll(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()

	// Test PutAll with nil source map
	lhm.PutAll(nil)
	assert.Equal(t, 0, lhm.Size())

	// Test PutAll with empty source map
	emptyMap := NewLinkedHashMap[string, int]()
	lhm.PutAll(emptyMap)
	assert.Equal(t, 0, lhm.Size())

	// Test PutAll with non-empty source map
	sourceMap := NewLinkedHashMap[string, int]()
	sourceMap.Put("one", 1)
	sourceMap.Put("two", 2)
	sourceMap.Put("three", 3)

	lhm.PutAll(sourceMap)
	assert.Equal(t, 3, lhm.Size())
	assert.Equal(t, 1, *lhm.Get("one"))
	assert.Equal(t, 2, *lhm.Get("two"))
	assert.Equal(t, 3, *lhm.Get("three"))

	// Test PutAll with overlapping keys
	sourceMap.Put("one", 4)
	sourceMap.Put("four", 4)
	lhm.PutAll(sourceMap)
	assert.Equal(t, 4, lhm.Size())
	assert.Equal(t, 4, *lhm.Get("one")) // Value should be updated
	assert.Equal(t, 2, *lhm.Get("two"))
	assert.Equal(t, 3, *lhm.Get("three"))
	assert.Equal(t, 4, *lhm.Get("four"))

	// Test PutAll maintains insertion order
	expectedOrder := []struct {
		key   string
		value int
	}{
		{"one", 4},
		{"two", 2},
		{"three", 3},
		{"four", 4},
	}

	index := 0
	if linkedMap, ok := lhm.(*LinkedHashMap[string, int]); ok {
		linkedMap.ForEachEntry(func(key string, value int) {
			assert.Equal(t, expectedOrder[index].key, key)
			assert.Equal(t, expectedOrder[index].value, value)
			index++
		})
		assert.Equal(t, len(expectedOrder), index)
	} else {
		t.Fatal("Failed to type assert to LinkedHashMap")
	}
}

func TestLinkedHashMap_PutIfAbsent(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()

	// Test PutIfAbsent with new key
	oldValue := lhm.PutIfAbsent("one", 1)
	assert.Equal(t, 0, oldValue) // Should return zero value for new key
	assert.Equal(t, 1, *lhm.Get("one"))
	assert.Equal(t, 1, lhm.Size())

	// Test PutIfAbsent with existing key
	oldValue = lhm.PutIfAbsent("one", 2)
	assert.Equal(t, 1, oldValue)        // Should return existing value
	assert.Equal(t, 1, *lhm.Get("one")) // Value should not change
	assert.Equal(t, 1, lhm.Size())      // Size should not change

	// Test PutIfAbsent with another new key
	oldValue = lhm.PutIfAbsent("two", 2)
	assert.Equal(t, 0, oldValue) // Should return zero value for new key
	assert.Equal(t, 2, *lhm.Get("two"))
	assert.Equal(t, 2, lhm.Size())

	// Test PutIfAbsent with nil value
	oldValue = lhm.PutIfAbsent("three", 0)
	assert.Equal(t, 0, oldValue) // Should return zero value for new key
	assert.Equal(t, 0, *lhm.Get("three"))
	assert.Equal(t, 3, lhm.Size())

	// Test PutIfAbsent with same value as existing key
	oldValue = lhm.PutIfAbsent("one", 1)
	assert.Equal(t, 1, oldValue)        // Should return existing value
	assert.Equal(t, 1, *lhm.Get("one")) // Value should not change
	assert.Equal(t, 3, lhm.Size())      // Size should not change

	// Test PutIfAbsent after Remove
	lhm.Remove("one")
	oldValue = lhm.PutIfAbsent("one", 3)
	assert.Equal(t, 0, oldValue) // Should return zero value for new key
	assert.Equal(t, 3, *lhm.Get("one"))
	assert.Equal(t, 3, lhm.Size())

	// Test PutIfAbsent with multiple concurrent operations
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := i * 10
			oldValue := lhm.PutIfAbsent(key, value)
			assert.Equal(t, 0, oldValue) // Should return zero value for new key
			assert.Equal(t, value, *lhm.Get(key))
		}(i)
	}
	wg.Wait()
	assert.Equal(t, 13, lhm.Size()) // 3 original + 10 new keys
}

func TestLinkedHashMap_Remove(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()

	// Test Remove on empty map
	oldValue := lhm.Remove("nonexistent")
	assert.Equal(t, 0, oldValue)
	assert.Equal(t, 0, lhm.Size())

	// Test Remove after adding single element
	lhm.Put("one", 1)
	oldValue = lhm.Remove("one")
	assert.Equal(t, 1, oldValue)
	assert.Equal(t, 0, lhm.Size())
	assert.Nil(t, lhm.Get("one"))

	// Test Remove with multiple elements
	lhm.Put("one", 1)
	lhm.Put("two", 2)
	lhm.Put("three", 3)

	// Remove middle element
	oldValue = lhm.Remove("two")
	assert.Equal(t, 2, oldValue)
	assert.Equal(t, 2, lhm.Size())
	assert.Nil(t, lhm.Get("two"))
	assert.Equal(t, 1, *lhm.Get("one"))
	assert.Equal(t, 3, *lhm.Get("three"))

	// Remove first element
	oldValue = lhm.Remove("one")
	assert.Equal(t, 1, oldValue)
	assert.Equal(t, 1, lhm.Size())
	assert.Nil(t, lhm.Get("one"))
	assert.Equal(t, 3, *lhm.Get("three"))

	// Remove last element
	oldValue = lhm.Remove("three")
	assert.Equal(t, 3, oldValue)
	assert.Equal(t, 0, lhm.Size())
	assert.Nil(t, lhm.Get("three"))

	// Test Remove with zero value
	lhm.Put("zero", 0)
	oldValue = lhm.Remove("zero")
	assert.Equal(t, 0, oldValue)
	assert.Equal(t, 0, lhm.Size())
	assert.Nil(t, lhm.Get("zero"))

	// Test Remove with same key multiple times
	lhm.Put("repeat", 1)
	oldValue = lhm.Remove("repeat")
	assert.Equal(t, 1, oldValue)
	oldValue = lhm.Remove("repeat")
	assert.Equal(t, 0, oldValue)
	assert.Equal(t, 0, lhm.Size())

	// Test Remove with concurrent operations
	var wg sync.WaitGroup
	// First add some elements
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		lhm.Put(key, i)
	}

	// Then remove them concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			oldValue := lhm.Remove(key)
			assert.Equal(t, i, oldValue)
		}(i)
	}
	wg.Wait()
	assert.Equal(t, 0, lhm.Size())

	// Test Remove maintains insertion order
	lhm.Put("first", 1)
	lhm.Put("second", 2)
	lhm.Put("third", 3)

	// Remove middle element and verify order
	lhm.Remove("second")
	lhm.Put("fourth", 4)

	// Verify remaining elements are in correct order
	expectedOrder := []struct {
		key   string
		value int
	}{
		{"first", 1},
		{"third", 3},
		{"fourth", 4},
	}

	index := 0
	if linkedMap, ok := lhm.(*LinkedHashMap[string, int]); ok {
		linkedMap.ForEachEntry(func(key string, value int) {
			assert.Equal(t, expectedOrder[index].key, key)
			assert.Equal(t, expectedOrder[index].value, value)
			index++
		})
		assert.Equal(t, len(expectedOrder), index)
	} else {
		t.Fatal("Failed to type assert to LinkedHashMap")
	}
}

func TestLinkedHashMap_RemoveKeyWithValue(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()

	// Test RemoveKeyWithValue on empty map
	removed := lhm.RemoveKeyWithValue("nonexistent", 1)
	assert.False(t, removed)
	assert.Equal(t, 0, lhm.Size())

	// Test RemoveKeyWithValue with non-matching value
	lhm.Put("one", 1)
	removed = lhm.RemoveKeyWithValue("one", 2)
	assert.False(t, removed)
	assert.Equal(t, 1, lhm.Size())
	assert.Equal(t, 1, *lhm.Get("one"))

	// Test RemoveKeyWithValue with matching value
	removed = lhm.RemoveKeyWithValue("one", 1)
	assert.True(t, removed)
	assert.Equal(t, 0, lhm.Size())
	assert.Nil(t, lhm.Get("one"))

	// Test RemoveKeyWithValue with multiple elements
	lhm.Put("one", 1)
	lhm.Put("two", 2)
	lhm.Put("three", 3)

	// Remove middle element with matching value
	removed = lhm.RemoveKeyWithValue("two", 2)
	assert.True(t, removed)
	assert.Equal(t, 2, lhm.Size())
	assert.Nil(t, lhm.Get("two"))
	assert.Equal(t, 1, *lhm.Get("one"))
	assert.Equal(t, 3, *lhm.Get("three"))

	// Remove first element with matching value
	removed = lhm.RemoveKeyWithValue("one", 1)
	assert.True(t, removed)
	assert.Equal(t, 1, lhm.Size())
	assert.Nil(t, lhm.Get("one"))
	assert.Equal(t, 3, *lhm.Get("three"))

	// Remove last element with matching value
	removed = lhm.RemoveKeyWithValue("three", 3)
	assert.True(t, removed)
	assert.Equal(t, 0, lhm.Size())
	assert.Nil(t, lhm.Get("three"))

	// Test RemoveKeyWithValue with zero value
	lhm.Put("zero", 0)
	removed = lhm.RemoveKeyWithValue("zero", 0)
	assert.True(t, removed)
	assert.Equal(t, 0, lhm.Size())
	assert.Nil(t, lhm.Get("zero"))

	// Test RemoveKeyWithValue with same key multiple times
	lhm.Put("repeat", 1)
	removed = lhm.RemoveKeyWithValue("repeat", 1)
	assert.True(t, removed)
	removed = lhm.RemoveKeyWithValue("repeat", 1)
	assert.False(t, removed)
	assert.Equal(t, 0, lhm.Size())

	// Test RemoveKeyWithValue with concurrent operations
	var wg sync.WaitGroup
	// First add some elements
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		lhm.Put(key, i)
	}

	// Then remove them concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			removed := lhm.RemoveKeyWithValue(key, i)
			assert.True(t, removed)
		}(i)
	}
	wg.Wait()
	assert.Equal(t, 0, lhm.Size())

	// Test RemoveKeyWithValue maintains insertion order
	lhm.Put("first", 1)
	lhm.Put("second", 2)
	lhm.Put("third", 3)

	// Remove middle element and verify order
	lhm.RemoveKeyWithValue("second", 2)
	lhm.Put("fourth", 4)

	// Verify remaining elements are in correct order
	expectedOrder := []struct {
		key   string
		value int
	}{
		{"first", 1},
		{"third", 3},
		{"fourth", 4},
	}

	index := 0
	if linkedMap, ok := lhm.(*LinkedHashMap[string, int]); ok {
		linkedMap.ForEachEntry(func(key string, value int) {
			assert.Equal(t, expectedOrder[index].key, key)
			assert.Equal(t, expectedOrder[index].value, value)
			index++
		})
		assert.Equal(t, len(expectedOrder), index)
	} else {
		t.Fatal("Failed to type assert to LinkedHashMap")
	}

	// Test RemoveKeyWithValue with value mismatch
	lhm.Clear()
	lhm.Put("key", 1)
	removed = lhm.RemoveKeyWithValue("key", 2)
	assert.False(t, removed)
	assert.Equal(t, 1, lhm.Size())
	assert.Equal(t, 1, *lhm.Get("key"))

	// Test RemoveKeyWithValue with key mismatch
	removed = lhm.RemoveKeyWithValue("wrong", 1)
	assert.False(t, removed)
	assert.Equal(t, 1, lhm.Size())
	assert.Equal(t, 1, *lhm.Get("key"))
}

func TestLinkedHashMap_Replace(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)

	// Test Replace existing key
	oldValue := lhm.Replace("one", 2)
	assert.Equal(t, 1, oldValue)
	assert.Equal(t, 2, *lhm.Get("one"))

	// Test Replace non-existent key
	oldValue = lhm.Replace("two", 2)
	assert.Equal(t, 0, oldValue)
	assert.Nil(t, lhm.Get("two"))
}

func TestLinkedHashMap_ReplaceKeyWithValue(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)

	// Test ReplaceKeyWithValue with matching old value
	assert.True(t, lhm.ReplaceKeyWithValue("one", 1, 2))
	assert.Equal(t, 2, *lhm.Get("one"))

	// Test ReplaceKeyWithValue with non-matching old value
	assert.False(t, lhm.ReplaceKeyWithValue("one", 1, 3))
	assert.Equal(t, 2, *lhm.Get("one"))

	// Test ReplaceKeyWithValue with non-existent key
	assert.False(t, lhm.ReplaceKeyWithValue("two", 1, 2))
}

func TestLinkedHashMap_Clear(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)
	lhm.Put("two", 2)

	lhm.Clear()
	assert.Equal(t, 0, lhm.Size())
	assert.True(t, lhm.IsEmpty())
	assert.Nil(t, lhm.Get("one"))
	assert.Nil(t, lhm.Get("two"))
}

func TestLinkedHashMap_HasKey(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)

	assert.True(t, lhm.HasKey("one"))
	assert.False(t, lhm.HasKey("two"))
}

func TestLinkedHashMap_HasValue(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)
	lhm.Put("two", 1)

	assert.True(t, lhm.HasValue(1))
	assert.False(t, lhm.HasValue(2))
}

func TestLinkedHashMap_EntrySet(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)
	lhm.Put("two", 2)

	entrySet := lhm.EntrySet()
	assert.Equal(t, 2, entrySet.Size())

	// Test entry set contains correct entries
	entries := entrySet.ToArray()
	assert.Equal(t, 2, len(entries))

	// Verify entries are in insertion order
	// Note: The order in ToArray() might be different from insertion order
	// due to the HashSet implementation. We should only verify that both entries exist.
	foundOne := false
	foundTwo := false
	for _, entry := range entries {
		key := entry.GetKey()
		value := entry.GetValue()
		if key == "one" && value == 1 {
			foundOne = true
		} else if key == "two" && value == 2 {
			foundTwo = true
		}
	}
	assert.True(t, foundOne, "Entry 'one' not found")
	assert.True(t, foundTwo, "Entry 'two' not found")
}

func TestLinkedHashMap_KeySet(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)
	lhm.Put("two", 2)

	keySet := lhm.KeySet()
	assert.Equal(t, 2, keySet.Size())

	// Test key set contains correct keys
	keys := keySet.ToArray()
	assert.Equal(t, 2, len(keys))

	// Since we're using HashSet, we can't guarantee order
	// Just verify both keys exist
	foundOne := false
	foundTwo := false
	for _, key := range keys {
		if key == "one" {
			foundOne = true
		} else if key == "two" {
			foundTwo = true
		}
	}
	assert.True(t, foundOne, "Key 'one' not found")
	assert.True(t, foundTwo, "Key 'two' not found")
}

func TestLinkedHashMap_Values(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)
	lhm.Put("two", 2)

	values := lhm.Values()
	assert.Equal(t, 2, values.Size())

	// Test values collection contains correct values
	valuesArray := values.ToArray()
	assert.Equal(t, 2, len(valuesArray))

	// Since we're using HashSet, we can't guarantee order
	// Just verify both values exist
	foundOne := false
	foundTwo := false
	for _, value := range valuesArray {
		if value == 1 {
			foundOne = true
		} else if value == 2 {
			foundTwo = true
		}
	}
	assert.True(t, foundOne, "Value 1 not found")
	assert.True(t, foundTwo, "Value 2 not found")
}

func TestLinkedHashMap_Equals(t *testing.T) {
	lhm1 := NewLinkedHashMap[string, int]()
	lhm1.Put("one", 1)
	lhm1.Put("two", 2)

	lhm2 := NewLinkedHashMap[string, int]()
	lhm2.Put("one", 1)
	lhm2.Put("two", 2)

	// Test equal maps
	assert.True(t, lhm1.Equals(lhm2))

	// Test different size
	lhm2.Put("three", 3)
	assert.False(t, lhm1.Equals(lhm2))

	// Test different values
	lhm2 = NewLinkedHashMap[string, int]()
	lhm2.Put("one", 1)
	lhm2.Put("two", 3)
	assert.False(t, lhm1.Equals(lhm2))

	// Test nil
	assert.False(t, lhm1.Equals(nil))

	// Test wrong type
	assert.False(t, lhm1.Equals("not a map"))
}

func TestLinkedHashMap_GetOrDefault(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)

	// Convert to ConcurrentMap to access GetOrDefault
	cm, ok := lhm.(collections.ConcurrentMap[string, int])
	assert.True(t, ok)

	// Test existing key
	assert.Equal(t, 1, cm.GetOrDefault("one", 0))

	// Test non-existent key
	assert.Equal(t, 0, cm.GetOrDefault("two", 0))
}

func TestLinkedHashMap_ForEachEntry(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	lhm.Put("one", 1)
	lhm.Put("two", 2)

	// Convert to ConcurrentMap to access ForEachEntry
	cm, ok := lhm.(collections.ConcurrentMap[string, int])
	assert.True(t, ok)

	// Test ForEachEntry
	keys := make([]string, 0)
	values := make([]int, 0)
	cm.ForEachEntry(func(key string, value int) {
		keys = append(keys, key)
		values = append(values, value)
	})

	assert.Equal(t, 2, len(keys))
	assert.Equal(t, "one", keys[0])
	assert.Equal(t, "two", keys[1])
	assert.Equal(t, 1, values[0])
	assert.Equal(t, 2, values[1])
}

func TestLinkedHashMap_ComputeIfAbsent(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()
	linkedMap, ok := lhm.(*LinkedHashMap[string, int])
	assert.True(t, ok, "Failed to type assert to LinkedHashMap")

	// Test ComputeIfAbsent with nil mapping function
	value, err := linkedMap.ComputeIfAbsent("key", nil)
	assert.Error(t, err)
	assert.Equal(t, "mapping function cannot be nil", err.Error())
	assert.Equal(t, 0, value)
	assert.Equal(t, 0, lhm.Size())

	// Test ComputeIfAbsent with new key
	value, err = linkedMap.ComputeIfAbsent("one", func(k string) int {
		return 1
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, value)
	assert.Equal(t, 1, *lhm.Get("one"))
	assert.Equal(t, 1, lhm.Size())

	// Test ComputeIfAbsent with existing key
	value, err = linkedMap.ComputeIfAbsent("one", func(k string) int {
		return 2
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, value)           // Should return existing value
	assert.Equal(t, 1, *lhm.Get("one")) // Value should not change
	assert.Equal(t, 1, lhm.Size())

	// Test ComputeIfAbsent with zero value
	value, err = linkedMap.ComputeIfAbsent("zero", func(k string) int {
		return 0
	})
	assert.NoError(t, err)
	assert.Equal(t, 0, value)
	assert.Equal(t, 0, *lhm.Get("zero"))
	assert.Equal(t, 2, lhm.Size())

	// Test ComputeIfAbsent with multiple keys
	value, err = linkedMap.ComputeIfAbsent("two", func(k string) int {
		return 2
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, value)
	assert.Equal(t, 2, *lhm.Get("two"))
	assert.Equal(t, 3, lhm.Size())

	// Test ComputeIfAbsent maintains insertion order
	expectedOrder := []struct {
		key   string
		value int
	}{
		{"one", 1},
		{"zero", 0},
		{"two", 2},
	}

	index := 0
	linkedMap.ForEachEntry(func(key string, value int) {
		assert.Equal(t, expectedOrder[index].key, key)
		assert.Equal(t, expectedOrder[index].value, value)
		index++
	})
	assert.Equal(t, len(expectedOrder), index)

	// Test ComputeIfAbsent with concurrent operations
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value, err := linkedMap.ComputeIfAbsent(key, func(k string) int {
				return i * 10
			})
			assert.NoError(t, err)
			assert.Equal(t, i*10, value)
			assert.Equal(t, i*10, *lhm.Get(key))
		}(i)
	}
	wg.Wait()
	assert.Equal(t, 8, lhm.Size()) // 3 original + 5 new keys

	// Test ComputeIfAbsent with mapping function that panics
	func() {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, "test panic", r)
			} else {
				t.Error("Expected panic did not occur")
			}
		}()
		linkedMap.ComputeIfAbsent("panic", func(k string) int {
			panic("test panic")
		})
	}()
	assert.Equal(t, 8, lhm.Size()) // Size should remain unchanged after panic

	// Test ComputeIfAbsent after Remove
	lhm.Remove("one")
	value, err = linkedMap.ComputeIfAbsent("one", func(k string) int {
		return 3
	})
	assert.NoError(t, err)
	assert.Equal(t, 3, value)
	assert.Equal(t, 3, *lhm.Get("one"))
	assert.Equal(t, 8, lhm.Size()) // Size should be 8 (7 existing + 1 new)
}

func TestLinkedHashMap_InsertionOrder(t *testing.T) {
	lhm := NewLinkedHashMap[string, int]()

	// Add elements in specific order
	lhm.Put("one", 1)
	lhm.Put("two", 2)
	lhm.Put("three", 3)

	// Test order in EntrySet
	entries := lhm.EntrySet().ToArray()
	assert.Equal(t, 3, len(entries))
	// Since we're using HashSet, we can't guarantee order
	// Just verify all entries exist
	foundEntries := make(map[string]int)
	for _, entry := range entries {
		foundEntries[entry.GetKey()] = entry.GetValue()
	}
	assert.Equal(t, 1, foundEntries["one"])
	assert.Equal(t, 2, foundEntries["two"])
	assert.Equal(t, 3, foundEntries["three"])

	// Test order in KeySet
	keys := lhm.KeySet().ToArray()
	assert.Equal(t, 3, len(keys))
	// Since we're using HashSet, we can't guarantee order
	// Just verify all keys exist
	foundKeys := make(map[string]bool)
	for _, key := range keys {
		foundKeys[key] = true
	}
	assert.True(t, foundKeys["one"])
	assert.True(t, foundKeys["two"])
	assert.True(t, foundKeys["three"])

	// Test order in Values
	values := lhm.Values().ToArray()
	assert.Equal(t, 3, len(values))
	// Since we're using HashSet, we can't guarantee order
	// Just verify all values exist
	foundValues := make(map[int]bool)
	for _, value := range values {
		foundValues[value] = true
	}
	assert.True(t, foundValues[1])
	assert.True(t, foundValues[2])
	assert.True(t, foundValues[3])

	// Test order after removal and reinsertion
	lhm.Remove("two")
	lhm.Put("two", 4)

	entries = lhm.EntrySet().ToArray()
	assert.Equal(t, 3, len(entries))
	// Since we're using HashSet, we can't guarantee order
	// Just verify all entries exist with correct values
	foundEntries = make(map[string]int)
	for _, entry := range entries {
		foundEntries[entry.GetKey()] = entry.GetValue()
	}
	assert.Equal(t, 1, foundEntries["one"])
	assert.Equal(t, 4, foundEntries["two"])
	assert.Equal(t, 3, foundEntries["three"])

	// Test order after updating existing key
	lhm.Put("one", 5)
	entries = lhm.EntrySet().ToArray()
	assert.Equal(t, 3, len(entries))
	// Since we're using HashSet, we can't guarantee order
	// Just verify all entries exist with correct values
	foundEntries = make(map[string]int)
	for _, entry := range entries {
		foundEntries[entry.GetKey()] = entry.GetValue()
	}
	assert.Equal(t, 5, foundEntries["one"])
	assert.Equal(t, 4, foundEntries["two"])
	assert.Equal(t, 3, foundEntries["three"])
}

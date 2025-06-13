package maps

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashTable(t *testing.T) {
	ht := NewHashTable[int, string]()
	assert.NotNil(t, ht, "HashTable should not be nil")
	assert.True(t, ht.IsEmpty(), "New HashTable should be empty")
}

func TestHashTable_BasicOperations(t *testing.T) {
	ht := NewHashTable[string, int]()

	// Test Put and Get
	ht.Put("one", 1)
	ht.Put("two", 2)
	ht.Put("three", 3)

	assert.Equal(t, 1, *ht.Get("one"))
	assert.Equal(t, 2, *ht.Get("two"))
	assert.Equal(t, 3, *ht.Get("three"))
	assert.Nil(t, ht.Get("four"))

	// Test ContainsKey
	assert.True(t, ht.ContainsKey("one"))
	assert.True(t, ht.ContainsKey("two"))
	assert.True(t, ht.ContainsKey("three"))
	assert.False(t, ht.ContainsKey("four"))

	// Test ContainsValue
	assert.True(t, ht.ContainsValue(1))
	assert.True(t, ht.ContainsValue(2))
	assert.True(t, ht.ContainsValue(3))
	assert.False(t, ht.ContainsValue(4))

	// Test Size
	assert.Equal(t, 3, ht.Size())

	// Test IsEmpty
	assert.False(t, ht.IsEmpty())

	// Test Remove
	oldValue := ht.Remove("two")
	assert.Equal(t, 2, oldValue)
	assert.Nil(t, ht.Get("two"))
	assert.Equal(t, 2, ht.Size())

	// Test Clear
	ht.Clear()
	assert.True(t, ht.IsEmpty())
	assert.Equal(t, 0, ht.Size())
}

func TestHashTable_EntrySet(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)
	ht.Put("two", 2)
	ht.Put("three", 3)

	entrySet := ht.EntrySet()
	assert.Equal(t, 3, entrySet.Size())

	// Convert to array for easier testing
	entries := entrySet.ToArray()
	assert.Equal(t, 3, len(entries))

	// Create a map for easier value lookup
	entryMap := make(map[string]int)
	for _, entry := range entries {
		entryMap[entry.GetKey()] = entry.GetValue()
	}

	assert.Equal(t, 1, entryMap["one"])
	assert.Equal(t, 2, entryMap["two"])
	assert.Equal(t, 3, entryMap["three"])
}

func TestHashTable_KeySet(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)
	ht.Put("two", 2)
	ht.Put("three", 3)

	keySet := ht.KeySet()
	assert.Equal(t, 3, keySet.Size())

	// Convert to array for easier testing
	keys := keySet.ToArray()
	assert.Equal(t, 3, len(keys))

	// Create a map for easier lookup
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}

	assert.True(t, keyMap["one"])
	assert.True(t, keyMap["two"])
	assert.True(t, keyMap["three"])
}

func TestHashTable_Values(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)
	ht.Put("two", 2)
	ht.Put("three", 3)

	values := ht.Values()
	assert.Equal(t, 3, values.Size())

	// Convert to array for easier testing
	valueArray := values.ToArray()
	assert.Equal(t, 3, len(valueArray))

	// Create a map for easier lookup
	valueMap := make(map[int]bool)
	for _, value := range valueArray {
		valueMap[value] = true
	}

	assert.True(t, valueMap[1])
	assert.True(t, valueMap[2])
	assert.True(t, valueMap[3])
}

func TestHashTable_PutIfAbsent(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)

	// Try to put a new key-value pair
	oldValue := ht.PutIfAbsent("two", 2)
	assert.Equal(t, 0, oldValue) // Should return zero value for V
	assert.Equal(t, 2, *ht.Get("two"))

	// Try to put an existing key-value pair
	oldValue = ht.PutIfAbsent("one", 10)
	assert.Equal(t, 1, oldValue)       // Should return the old value
	assert.Equal(t, 1, *ht.Get("one")) // Value should not change
}

func TestHashTable_PutIfAbsent_ExistingKey(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("foo", 42)
	ret := ht.PutIfAbsent("foo", 99)
	assert.Equal(t, 42, ret, "PutIfAbsent should return the existing value if key exists")
	assert.Equal(t, 42, *ht.Get("foo"), "PutIfAbsent should not overwrite the existing value")
}

func TestHashTable_Replace(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)

	// Replace existing key
	oldValue := ht.Replace("one", 10)
	assert.Equal(t, 1, oldValue)
	assert.Equal(t, 10, *ht.Get("one"))

	// Try to replace non-existing key
	oldValue = ht.Replace("two", 2)
	assert.Equal(t, 0, oldValue) // Should return zero value for V
	assert.Nil(t, ht.Get("two"))
}

func TestHashTable_ReplaceKeyWithValue(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)

	// Replace with matching old value
	success := ht.ReplaceKeyWithValue("one", 1, 10)
	assert.True(t, success)
	assert.Equal(t, 10, *ht.Get("one"))

	// Try to replace with non-matching old value
	success = ht.ReplaceKeyWithValue("one", 1, 20)
	assert.False(t, success)
	assert.Equal(t, 10, *ht.Get("one"))

	// Try to replace non-existing key
	success = ht.ReplaceKeyWithValue("two", 1, 2)
	assert.False(t, success)
	assert.Nil(t, ht.Get("two"))
}

func TestHashTable_RemoveKeyWithValue(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)

	// Remove with matching value
	success := ht.RemoveKeyWithValue("one", 1)
	assert.True(t, success)
	assert.Nil(t, ht.Get("one"))

	// Try to remove with non-matching value
	ht.Put("two", 2)
	success = ht.RemoveKeyWithValue("two", 1)
	assert.False(t, success)
	assert.Equal(t, 2, *ht.Get("two"))

	// Try to remove non-existing key
	success = ht.RemoveKeyWithValue("three", 1)
	assert.False(t, success)
}

func TestHashTable_PutAll(t *testing.T) {
	ht1 := NewHashTable[string, int]()
	ht1.Put("one", 1)
	ht1.Put("two", 2)

	ht2 := NewHashTable[string, int]()
	ht2.Put("three", 3)
	ht2.Put("four", 4)

	ht1.PutAll(ht2)

	assert.Equal(t, 4, ht1.Size())
	assert.Equal(t, 3, *ht1.Get("three"))
	assert.Equal(t, 4, *ht1.Get("four"))

	// Test with nil map
	ht1.PutAll(nil)
	assert.Equal(t, 4, ht1.Size())
}

func TestHashTable_HasKeyAndHasValue(t *testing.T) {
	ht := NewHashTable[string, int]()
	ht.Put("one", 1)
	ht.Put("two", 2)

	assert.True(t, ht.HasKey("one"))
	assert.True(t, ht.HasKey("two"))
	assert.False(t, ht.HasKey("three"))

	assert.True(t, ht.HasValue(1))
	assert.True(t, ht.HasValue(2))
	assert.False(t, ht.HasValue(3))
}

func TestHashTable_Concurrent(t *testing.T) {
	ht := NewHashTable[int, int]()
	var wg sync.WaitGroup
	iterations := 1000
	goroutines := 10

	// Test concurrent puts and gets
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := id*iterations + j
				ht.Put(key, j)
				val := ht.Get(key)
				if val == nil || *val != j {
					t.Errorf("Concurrent test: Get returned wrong value for key %d, got %v, want %d", key, val, j)
				}
			}
		}(i)
	}

	wg.Wait()

	// Verify the final state
	for i := 0; i < goroutines; i++ {
		for j := 0; j < iterations; j++ {
			key := i*iterations + j
			val := ht.Get(key)
			if val == nil || *val != j {
				t.Errorf("Final state: Wrong value for key %d, got %v, want %d", key, val, j)
			}
		}
	}

	// Test concurrent removes
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := id*iterations + j
				val := ht.Remove(key)
				if val != j {
					t.Errorf("Concurrent test: Remove returned wrong value for key %d, got %d, want %d", key, val, j)
				}
			}
		}(i)
	}

	wg.Wait()
	assert.True(t, ht.IsEmpty(), "Map should be empty after concurrent removes")
}

func TestHashTable_Equals(t *testing.T) {
	ht1 := NewHashTable[string, int]()
	ht1.Put("one", 1)
	ht1.Put("two", 2)
	ht1.Put("three", 3)

	ht2 := NewHashTable[string, int]()
	ht2.Put("one", 1)
	ht2.Put("two", 2)
	ht2.Put("three", 3)

	ht3 := NewHashTable[string, int]()
	ht3.Put("one", 1)
	ht3.Put("two", 2)
	ht3.Put("three", 4)

	assert.True(t, ht1.Equals(ht2))
	assert.False(t, ht1.Equals(ht3))
	assert.False(t, ht1.Equals(nil))
	assert.False(t, ht1.Equals("not a map"))
}

func TestHashTable_Equals_DifferentSizes(t *testing.T) {
	ht1 := NewHashTable[string, int]()
	ht2 := NewHashTable[string, int]()

	ht1.Put("a", 1)
	ht1.Put("b", 2)
	ht2.Put("a", 1)
	// ht2 has only one entry, ht1 has two

	assert.False(t, ht1.Equals(ht2), "Equals should return false when sizes differ")
	assert.False(t, ht2.Equals(ht1), "Equals should return false when sizes differ (reverse)")
}

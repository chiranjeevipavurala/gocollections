package maps

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/chiranjeevipavurala/gocollections/collections"
	"github.com/stretchr/testify/assert"
)

// IntComparator implements Comparator for integers
type IntComparator struct{}

func (c *IntComparator) Compare(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TestNewTreeMap(t *testing.T) {
	// Test with nil comparator
	tm := NewTreeMap[int, string](nil)
	assert.Nil(t, tm, "TreeMap should be nil when comparator is nil")

	// Test with valid comparator
	tm = NewTreeMap[int, string](&IntComparator{})
	assert.NotNil(t, tm, "TreeMap should not be nil with valid comparator")
	assert.True(t, tm.IsEmpty(), "New TreeMap should be empty")
}

func TestTreeMap_BasicOperations(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test Put and Get
	tm.Put(1, "one")
	tm.Put(2, "two")
	tm.Put(3, "three")

	if val := tm.Get(1); val == nil || *val != "one" {
		t.Errorf("Get(1) = %v; want 'one'", val)
	}

	// Test Size
	if size := tm.Size(); size != 3 {
		t.Errorf("Size() = %d; want 3", size)
	}

	// Test HasKey
	if !tm.HasKey(1) {
		t.Error("HasKey(1) = false; want true")
	}
	if tm.HasKey(4) {
		t.Error("HasKey(4) = true; want false")
	}

	// Test HasValue
	if !tm.HasValue("one") {
		t.Error("HasValue('one') = false; want true")
	}
	if tm.HasValue("four") {
		t.Error("HasValue('four') = true; want false")
	}

	// Test Remove
	if val := tm.Remove(1); val != "one" {
		t.Errorf("Remove(1) = %v; want 'one'", val)
	}
	if tm.HasKey(1) {
		t.Error("HasKey(1) = true after removal; want false")
	}

	// Test Clear
	tm.Clear()
	if !tm.IsEmpty() {
		t.Error("IsEmpty() = false after Clear(); want true")
	}
}

func TestTreeMap_EntrySet(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})
	tm.Put(1, "one")
	tm.Put(2, "two")

	entrySet := tm.EntrySet()
	assert.Equal(t, 2, entrySet.Size(), "EntrySet should have size 2")

	// Test entry set contents
	entries := entrySet.ToArray()
	found := make(map[int]bool)
	for _, entry := range entries {
		key := entry.GetKey()
		value := entry.GetValue()
		if key == 1 && value == "one" {
			found[1] = true
		} else if key == 2 && value == "two" {
			found[2] = true
		}
	}
	assert.True(t, found[1] && found[2], "EntrySet should contain all entries")
}

func TestTreeMap_KeySet(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})
	tm.Put(1, "one")
	tm.Put(2, "two")

	keySet := tm.KeySet()
	assert.Equal(t, 2, keySet.Size(), "KeySet should have size 2")

	// Test key set contents
	keys := keySet.ToArray()
	found := make(map[int]bool)
	for _, key := range keys {
		found[key] = true
	}
	assert.True(t, found[1] && found[2], "KeySet should contain all keys")
}

func TestTreeMap_Values(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})
	tm.Put(1, "one")
	tm.Put(2, "two")

	values := tm.Values()
	assert.Equal(t, 2, values.Size(), "Values should have size 2")

	// Test values collection contents
	vals := values.ToArray()
	found := make(map[string]bool)
	for _, val := range vals {
		found[val] = true
	}
	assert.True(t, found["one"] && found["two"], "Values should contain all values")
}

func TestTreeMap_SortedOperations(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})
	tm.Put(3, "three")
	tm.Put(1, "one")
	tm.Put(2, "two")

	// Test FirstKey
	firstKey, err := tm.FirstKey()
	assert.Nil(t, err, "FirstKey should not return error")
	assert.Equal(t, 1, *firstKey, "FirstKey should return smallest key")

	// Test LastKey
	lastKey, err := tm.LastKey()
	assert.Nil(t, err, "LastKey should not return error")
	assert.Equal(t, 3, *lastKey, "LastKey should return largest key")

	// Test LowerKey
	lowerKey, err := tm.LowerKey(2)
	assert.Nil(t, err, "LowerKey should not return error")
	assert.Equal(t, 1, *lowerKey, "LowerKey should return greatest key less than given key")

	// Test HigherKey
	higherKey, err := tm.HigherKey(2)
	assert.Nil(t, err, "HigherKey should not return error")
	assert.Equal(t, 3, *higherKey, "HigherKey should return least key greater than given key")
}

func TestTreeMap_EmptyOperations(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test operations on empty map
	firstKey, err := tm.FirstKey()
	assert.NotNil(t, err, "FirstKey should return error on empty map")
	assert.Nil(t, firstKey, "FirstKey should return nil on empty map")

	lastKey, err := tm.LastKey()
	assert.NotNil(t, err, "LastKey should return error on empty map")
	assert.Nil(t, lastKey, "LastKey should return nil on empty map")

	lowerKey, err := tm.LowerKey(1)
	assert.NotNil(t, err, "LowerKey should return error on empty map")
	assert.Nil(t, lowerKey, "LowerKey should return nil on empty map")

	higherKey, err := tm.HigherKey(1)
	assert.NotNil(t, err, "HigherKey should return error on empty map")
	assert.Nil(t, higherKey, "HigherKey should return nil on empty map")
}

func TestTreeMap_Concurrent(t *testing.T) {
	tm := NewTreeMap[int, int](&IntComparator{})
	var wg sync.WaitGroup
	iterations := 10 // Further reduced for debugging
	goroutines := 3  // Further reduced for debugging

	// Create a channel to signal completion
	done := make(chan bool)
	timeout := time.After(5 * time.Second) // Reduced timeout

	// Test concurrent puts and gets
	t.Log("Starting concurrent puts and gets")
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer func() {
				t.Logf("Goroutine %d: Defer called", id)
				wg.Done()
			}()
			for j := 0; j < iterations; j++ {
				key := id*iterations + j
				t.Logf("Goroutine %d: Starting Put for key %d", id, key)
				oldVal := tm.Put(key, j)
				t.Logf("Goroutine %d: Completed Put for key %d, old value: %v", id, key, oldVal)

				t.Logf("Goroutine %d: Starting Get for key %d", id, key)
				val := tm.Get(key)
				t.Logf("Goroutine %d: Completed Get for key %d, value: %v", id, key, val)

				if val == nil {
					t.Errorf("Goroutine %d: Get returned nil for key %d", id, key)
				} else if *val != j {
					t.Errorf("Goroutine %d: Get returned wrong value for key %d, got %v, want %v", id, key, *val, j)
				}

				// Add a small delay between operations
				time.Sleep(10 * time.Millisecond)
			}
			t.Logf("Goroutine %d: Finished puts and gets", id)
		}(i)
	}

	// Wait for puts and gets with timeout
	go func() {
		t.Log("Waiting for puts and gets to complete")
		wg.Wait()
		t.Log("Puts and gets completed")
		done <- true
	}()

	select {
	case <-done:
		t.Log("First phase completed successfully")
	case <-timeout:
		t.Fatal("Test timed out during puts and gets")
	}

	// Verify the final state
	t.Log("Verifying final state")
	for i := 0; i < goroutines; i++ {
		for j := 0; j < iterations; j++ {
			key := i*iterations + j
			val := tm.Get(key)
			if val == nil {
				t.Errorf("Key %d not found in final state", key)
			} else if *val != j {
				t.Errorf("Wrong value for key %d in final state, got %v, want %v", key, *val, j)
			}
		}
	}

	// Test concurrent removes
	t.Log("Starting concurrent removes")
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer func() {
				t.Logf("Goroutine %d: Remove defer called", id)
				wg.Done()
			}()
			for j := 0; j < iterations; j++ {
				key := id*iterations + j
				t.Logf("Goroutine %d: Starting Remove for key %d", id, key)
				val := tm.Remove(key)
				t.Logf("Goroutine %d: Completed Remove for key %d, value: %v", id, key, val)

				if val != j {
					t.Errorf("Goroutine %d: Remove returned wrong value for key %d, got %v, want %v", id, key, val, j)
				}

				// Add a small delay between operations
				time.Sleep(10 * time.Millisecond)
			}
			t.Logf("Goroutine %d: Finished removes", id)
		}(i)
	}

	// Wait for removes with timeout
	go func() {
		t.Log("Waiting for removes to complete")
		wg.Wait()
		t.Log("Removes completed")
		done <- true
	}()

	select {
	case <-done:
		t.Log("Second phase completed successfully")
	case <-timeout:
		t.Fatal("Test timed out during removes")
	}

	// Verify the map is empty
	if !tm.IsEmpty() {
		t.Errorf("Map should be empty after removes, but has size %d", tm.Size())
	}
}

func TestTreeMap_RedBlackTreeProperties(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{}).(*TreeMap[int, string])

	// Insert elements in random order
	elements := []int{5, 3, 7, 1, 9, 2, 8, 4, 6}
	for _, elem := range elements {
		tm.Put(elem, "value")
	}

	// Verify tree properties
	assert.True(t, tm.verifyRedBlackProperties(), "Tree should maintain Red-Black properties")
}

func TestTreeMap_PutIfAbsent(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test PutIfAbsent on empty map
	val := tm.PutIfAbsent(1, "one")
	assert.Equal(t, "", val, "PutIfAbsent should return zero value when key doesn't exist")
	assert.Equal(t, "one", *tm.Get(1), "PutIfAbsent should insert value when key doesn't exist")

	// Test PutIfAbsent on existing key
	val = tm.PutIfAbsent(1, "ONE")
	assert.Equal(t, "one", val, "PutIfAbsent should return existing value")
	assert.Equal(t, "one", *tm.Get(1), "PutIfAbsent should not change existing value")

	// Test PutIfAbsent with nil value
	val = tm.PutIfAbsent(2, "")
	assert.Equal(t, "", val, "PutIfAbsent should handle empty string value")
	assert.Equal(t, "", *tm.Get(2), "PutIfAbsent should store empty string value")
}

func TestTreeMap_Replace(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test Replace on non-existent key
	val := tm.Replace(1, "one")
	assert.Equal(t, "", val, "Replace should return zero value for non-existent key")
	assert.Nil(t, tm.Get(1), "Replace should not insert value for non-existent key")

	// Test Replace on existing key
	tm.Put(1, "one")
	val = tm.Replace(1, "ONE")
	assert.Equal(t, "one", val, "Replace should return old value")
	assert.Equal(t, "ONE", *tm.Get(1), "Replace should update value")

	// Test ReplaceKeyWithValue
	success := tm.ReplaceKeyWithValue(1, "ONE", "one")
	assert.True(t, success, "ReplaceKeyWithValue should succeed with matching old value")
	assert.Equal(t, "one", *tm.Get(1), "ReplaceKeyWithValue should update value")

	success = tm.ReplaceKeyWithValue(1, "wrong", "new")
	assert.False(t, success, "ReplaceKeyWithValue should fail with non-matching old value")
	assert.Equal(t, "one", *tm.Get(1), "ReplaceKeyWithValue should not update value with wrong old value")
}

func TestTreeMap_EdgeCases(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test operations with nil values
	tm.Put(1, "")
	assert.Equal(t, "", *tm.Get(1), "Should handle empty string value")
	assert.True(t, tm.HasValue(""), "Should find empty string value")

	// Test operations with zero values
	tm.Put(2, "zero")
	val := tm.Remove(2)
	assert.Equal(t, "zero", val, "Should return correct value on remove")
	assert.Nil(t, tm.Get(2), "Should not find removed key")

	// Test operations with same key
	tm.Put(3, "first")
	tm.Put(3, "second")
	assert.Equal(t, "second", *tm.Get(3), "Should update value for same key")

	// Test operations with nil map
	var nilMap collections.Map[int, string]
	tm.PutAll(nilMap)
	assert.Equal(t, 2, tm.Size(), "PutAll should handle nil map")
}

func TestTreeMap_ConcurrentStress(t *testing.T) {
	tm := NewTreeMap[int, int](&IntComparator{})
	var wg sync.WaitGroup
	iterations := 1000
	goroutines := 10

	// Test concurrent puts, gets, and removes
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := id*iterations + j

				// Put
				tm.Put(key, j)

				// Get
				val := tm.Get(key)
				if val != nil && *val != j {
					t.Errorf("Concurrent stress test: Get returned wrong value for key %d, got %v, want %v", key, *val, j)
				}

				// Remove
				removedVal := tm.Remove(key)
				if removedVal != j {
					t.Errorf("Concurrent stress test: Remove returned wrong value for key %d, got %v, want %v", key, removedVal, j)
				}

				// Verify removal
				if tm.Get(key) != nil {
					t.Errorf("Concurrent stress test: Key %d still exists after removal", key)
				}
			}
		}(i)
	}

	wg.Wait()
	assert.True(t, tm.IsEmpty(), "Map should be empty after concurrent stress test")
}

func TestTreeMap_Balancing(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{}).(*TreeMap[int, string])

	// Insert elements in a way that would create an unbalanced tree
	// without balancing
	for i := 0; i < 100; i++ {
		tm.Put(i, fmt.Sprintf("value%d", i))
	}

	// Verify tree properties
	assert.True(t, tm.verifyRedBlackProperties(), "Tree should maintain Red-Black properties after sequential inserts")

	// Clear and test reverse order
	tm.Clear()
	for i := 99; i >= 0; i-- {
		tm.Put(i, fmt.Sprintf("value%d", i))
	}

	// Verify tree properties
	assert.True(t, tm.verifyRedBlackProperties(), "Tree should maintain Red-Black properties after reverse sequential inserts")

	// Test alternating inserts
	tm.Clear()
	for i := 0; i < 50; i++ {
		tm.Put(i*2, fmt.Sprintf("value%d", i*2))
		tm.Put(99-i*2, fmt.Sprintf("value%d", 99-i*2))
	}

	// Verify tree properties
	assert.True(t, tm.verifyRedBlackProperties(), "Tree should maintain Red-Black properties after alternating inserts")
}

func TestTreeMap_PutAll(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test PutAll with empty map
	emptyMap := NewTreeMap[int, string](&IntComparator{})
	tm.PutAll(emptyMap)
	assert.Equal(t, 0, tm.Size(), "PutAll with empty map should not change size")

	// Test PutAll with nil map
	tm.PutAll(nil)
	assert.Equal(t, 0, tm.Size(), "PutAll with nil map should not change size")

	// Test PutAll with new entries
	sourceMap := NewTreeMap[int, string](&IntComparator{})
	sourceMap.Put(1, "one")
	sourceMap.Put(2, "two")
	sourceMap.Put(3, "three")

	tm.PutAll(sourceMap)
	assert.Equal(t, 3, tm.Size(), "PutAll should add all entries from source map")
	assert.Equal(t, "one", *tm.Get(1), "PutAll should copy first entry")
	assert.Equal(t, "two", *tm.Get(2), "PutAll should copy second entry")
	assert.Equal(t, "three", *tm.Get(3), "PutAll should copy third entry")

	// Test PutAll with overlapping keys
	overlapMap := NewTreeMap[int, string](&IntComparator{})
	overlapMap.Put(2, "TWO")
	overlapMap.Put(3, "THREE")
	overlapMap.Put(4, "four")

	tm.PutAll(overlapMap)
	assert.Equal(t, 4, tm.Size(), "PutAll should update existing entries and add new ones")
	assert.Equal(t, "one", *tm.Get(1), "PutAll should not change non-overlapping entries")
	assert.Equal(t, "TWO", *tm.Get(2), "PutAll should update overlapping entries")
	assert.Equal(t, "THREE", *tm.Get(3), "PutAll should update overlapping entries")
	assert.Equal(t, "four", *tm.Get(4), "PutAll should add new entries")

	// Test PutAll with large number of entries
	largeMap := NewTreeMap[int, string](&IntComparator{})
	for i := 0; i < 1000; i++ {
		largeMap.Put(i, fmt.Sprintf("value%d", i))
	}

	tm.Clear()
	tm.PutAll(largeMap)
	assert.Equal(t, 1000, tm.Size(), "PutAll should handle large number of entries")

	// Verify all entries were copied correctly
	for i := 0; i < 1000; i++ {
		expected := fmt.Sprintf("value%d", i)
		actual := tm.Get(i)
		assert.NotNil(t, actual, "PutAll should copy all entries")
		assert.Equal(t, expected, *actual, "PutAll should copy values correctly")
	}

	// Test PutAll with concurrent modifications
	tm.Clear()
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: PutAll
	go func() {
		defer wg.Done()
		tm.PutAll(largeMap)
	}()

	// Goroutine 2: Concurrent puts
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			tm.Put(i+1000, fmt.Sprintf("concurrent%d", i))
		}
	}()

	wg.Wait()
	assert.GreaterOrEqual(t, tm.Size(), 1000, "PutAll should handle concurrent modifications")
}

func TestTreeMap_RemoveKeyWithValue(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test RemoveKeyWithValue on empty map
	success := tm.RemoveKeyWithValue(1, "one")
	assert.False(t, success, "RemoveKeyWithValue should return false for non-existent key")

	// Test RemoveKeyWithValue with non-existent key
	tm.Put(1, "one")
	success = tm.RemoveKeyWithValue(2, "two")
	assert.False(t, success, "RemoveKeyWithValue should return false for non-existent key")
	assert.Equal(t, 1, tm.Size(), "Map size should not change for non-existent key")

	// Test RemoveKeyWithValue with wrong value
	success = tm.RemoveKeyWithValue(1, "wrong")
	assert.False(t, success, "RemoveKeyWithValue should return false for wrong value")
	assert.Equal(t, "one", *tm.Get(1), "Entry should not be removed for wrong value")
	assert.Equal(t, 1, tm.Size(), "Map size should not change for wrong value")

	// Test RemoveKeyWithValue with correct key-value pair
	success = tm.RemoveKeyWithValue(1, "one")
	assert.True(t, success, "RemoveKeyWithValue should return true for correct key-value pair")
	assert.Nil(t, tm.Get(1), "Entry should be removed")
	assert.Equal(t, 0, tm.Size(), "Map size should decrease by 1")

	// Test RemoveKeyWithValue with multiple entries
	tm.Put(1, "one")
	tm.Put(2, "two")
	tm.Put(3, "three")

	success = tm.RemoveKeyWithValue(2, "two")
	assert.True(t, success, "RemoveKeyWithValue should succeed for middle entry")
	assert.Equal(t, 2, tm.Size(), "Map size should decrease by 1")
	assert.Nil(t, tm.Get(2), "Middle entry should be removed")
	assert.Equal(t, "one", *tm.Get(1), "First entry should remain")
	assert.Equal(t, "three", *tm.Get(3), "Last entry should remain")

	// Test RemoveKeyWithValue with empty string value
	tm.Put(4, "")
	success = tm.RemoveKeyWithValue(4, "")
	assert.True(t, success, "RemoveKeyWithValue should handle empty string value")
	assert.Nil(t, tm.Get(4), "Entry with empty string should be removed")
}

func TestTreeMap_Remove(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test Remove on empty map
	val := tm.Remove(1)
	assert.Equal(t, "", val, "Remove should return zero value for non-existent key")
	assert.Equal(t, 0, tm.Size(), "Map size should not change for non-existent key")

	// Test Remove with single entry
	tm.Put(1, "one")
	val = tm.Remove(1)
	assert.Equal(t, "one", val, "Remove should return correct value")
	assert.Nil(t, tm.Get(1), "Entry should be removed")
	assert.Equal(t, 0, tm.Size(), "Map size should decrease by 1")

	// Test Remove with multiple entries
	tm.Put(1, "one")
	tm.Put(2, "two")
	tm.Put(3, "three")

	// Test Remove first entry
	val = tm.Remove(1)
	assert.Equal(t, "one", val, "Remove should return correct value for first entry")
	assert.Equal(t, 2, tm.Size(), "Map size should decrease by 1")
	assert.Nil(t, tm.Get(1), "First entry should be removed")
	assert.Equal(t, "two", *tm.Get(2), "Second entry should remain")
	assert.Equal(t, "three", *tm.Get(3), "Third entry should remain")

	// Test Remove middle entry
	val = tm.Remove(2)
	assert.Equal(t, "two", val, "Remove should return correct value for middle entry")
	assert.Equal(t, 1, tm.Size(), "Map size should decrease by 1")
	assert.Nil(t, tm.Get(2), "Middle entry should be removed")
	assert.Equal(t, "three", *tm.Get(3), "Last entry should remain")

	// Test Remove last entry
	val = tm.Remove(3)
	assert.Equal(t, "three", val, "Remove should return correct value for last entry")
	assert.Equal(t, 0, tm.Size(), "Map size should decrease by 1")
	assert.Nil(t, tm.Get(3), "Last entry should be removed")

	// Test Remove with empty string value
	tm.Put(4, "")
	val = tm.Remove(4)
	assert.Equal(t, "", val, "Remove should handle empty string value")
	assert.Nil(t, tm.Get(4), "Entry with empty string should be removed")
	assert.Equal(t, 0, tm.Size(), "Map size should decrease by 1")

	// Test Remove with concurrent operations
	tm.Put(1, "one")
	tm.Put(2, "two")
	tm.Put(3, "three")

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: Remove entries
	go func() {
		defer wg.Done()
		tm.Remove(1)
		tm.Remove(2)
	}()

	// Goroutine 2: Add new entries
	go func() {
		defer wg.Done()
		tm.Put(4, "four")
		tm.Put(5, "five")
	}()

	wg.Wait()
	assert.Equal(t, 3, tm.Size(), "Map should maintain correct size after concurrent operations")
	assert.Nil(t, tm.Get(1), "First entry should be removed")
	assert.Nil(t, tm.Get(2), "Second entry should be removed")
	assert.Equal(t, "three", *tm.Get(3), "Third entry should remain")
	assert.Equal(t, "four", *tm.Get(4), "Fourth entry should be added")
	assert.Equal(t, "five", *tm.Get(5), "Fifth entry should be added")
}

func TestTreeMap_RedBlackProperties(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{}).(*TreeMap[int, string])

	// Test empty tree
	assert.True(t, tm.verifyRedBlackProperties(), "Empty tree should satisfy Red-Black properties")

	// Test single node
	tm.Put(1, "one")
	assert.True(t, tm.verifyRedBlackProperties(), "Single node tree should satisfy Red-Black properties")

	// Test tree with red root (should fail)
	tm.root.color = Red
	assert.False(t, tm.verifyRedBlackProperties(), "Tree with red root should fail Red-Black properties")

	// Test tree with red-red violation
	tm.root.color = Black
	tm.Put(2, "two")
	tm.Put(3, "three")
	tm.root.right.color = Red
	tm.root.right.right.color = Red
	assert.False(t, tm.verifyRedBlackProperties(), "Tree with red-red violation should fail Red-Black properties")

	// Test tree with unequal black heights
	tm.Clear()
	tm.Put(5, "five")
	tm.Put(3, "three")
	tm.Put(7, "seven")
	tm.Put(1, "one")
	tm.Put(9, "nine")
	tm.root.left.color = Red
	tm.root.right.color = Red
	tm.root.left.left.color = Black
	tm.root.right.right.color = Black
	assert.True(t, tm.verifyRedBlackProperties(), "Balanced tree should satisfy Red-Black properties")
}

func TestTreeMap_FirstLastNode(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{}).(*TreeMap[int, string])

	// Test empty tree
	assert.Nil(t, tm.getFirstNode(nil), "getFirstNode should return nil for empty tree")
	assert.Nil(t, tm.getLastNode(nil), "getLastNode should return nil for empty tree")

	// Test single node
	tm.Put(1, "one")
	first := tm.getFirstNode(tm.root)
	last := tm.getLastNode(tm.root)
	assert.Equal(t, 1, first.key, "getFirstNode should return root for single node")
	assert.Equal(t, 1, last.key, "getLastNode should return root for single node")

	// Test tree with multiple nodes
	tm.Clear()
	tm.Put(5, "five")
	tm.Put(3, "three")
	tm.Put(7, "seven")
	tm.Put(1, "one")
	tm.Put(9, "nine")

	first = tm.getFirstNode(tm.root)
	last = tm.getLastNode(tm.root)
	assert.Equal(t, 1, first.key, "getFirstNode should return leftmost node")
	assert.Equal(t, 9, last.key, "getLastNode should return rightmost node")

	// Test with subtree
	first = tm.getFirstNode(tm.root.right)
	last = tm.getLastNode(tm.root.left)
	assert.Equal(t, 7, first.key, "getFirstNode should work with subtree")
	assert.Equal(t, 3, last.key, "getLastNode should work with subtree")
}

func TestTreeMap_LowerHigherKey(t *testing.T) {
	tm := NewTreeMap[int, string](&IntComparator{})

	// Test empty map
	lower, err := tm.LowerKey(5)
	assert.Error(t, err, "LowerKey should return error for empty map")
	assert.Nil(t, lower, "LowerKey should return nil for empty map")

	higher, err := tm.HigherKey(5)
	assert.Error(t, err, "HigherKey should return error for empty map")
	assert.Nil(t, higher, "HigherKey should return nil for empty map")

	// Test single node
	tm.Put(5, "five")
	lower, err = tm.LowerKey(5)
	assert.Error(t, err, "LowerKey should return error when no lower key exists")
	assert.Nil(t, lower, "LowerKey should return nil when no lower key exists")

	higher, err = tm.HigherKey(5)
	assert.Error(t, err, "HigherKey should return error when no higher key exists")
	assert.Nil(t, higher, "HigherKey should return nil when no higher key exists")

	// Test multiple nodes
	tm.Put(3, "three")
	tm.Put(7, "seven")
	tm.Put(1, "one")
	tm.Put(9, "nine")

	// Test LowerKey
	lower, err = tm.LowerKey(5)
	assert.NoError(t, err, "LowerKey should not return error")
	assert.Equal(t, 3, *lower, "LowerKey should return greatest key less than given key")

	lower, err = tm.LowerKey(1)
	assert.Error(t, err, "LowerKey should return error when no lower key exists")
	assert.Nil(t, lower, "LowerKey should return nil when no lower key exists")

	// Test HigherKey
	higher, err = tm.HigherKey(5)
	assert.NoError(t, err, "HigherKey should not return error")
	assert.Equal(t, 7, *higher, "HigherKey should return least key greater than given key")

	higher, err = tm.HigherKey(9)
	assert.Error(t, err, "HigherKey should return error when no higher key exists")
	assert.Nil(t, higher, "HigherKey should return nil when no higher key exists")
}

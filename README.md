# Go Collections

A comprehensive collection of data structures and algorithms implemented in Go, providing efficient and type-safe implementations of common data structures.

## Features

- **Type Safety**: All collections are implemented using Go's generics, ensuring type safety at compile time.
- **Thread Safety**: All collections are thread-safe, using mutex locks for concurrent access.
- **Comprehensive Testing**: Extensive test coverage for all collections and their operations.
- **Error Handling**: Proper error handling using custom error codes for all operations that can fail.
- **Documentation**: Well-documented code with examples and usage patterns.

## Collections

### Lists

#### ArrayList
A dynamic array implementation that automatically resizes when needed.

**Features:**
- O(1) random access
- O(n) insertion/deletion at arbitrary positions
- Automatic resizing
- Thread-safe operations
- Batch operations support
- Fast operations for common use cases

**Key Methods:**
- `Add(element E) bool`
- `AddAtIndex(index int, element E) error`
- `Remove(element E) bool`
- `RemoveAtIndex(index int) (*E, error)`
- `Get(index int) (*E, error)`
- `Set(index int, element E) (*E, error)`
- `Size() int`
- `IsEmpty() bool`
- `Clear()`
- `Contains(element E) bool`
- `IndexOf(element E) int`
- `LastIndexOf(element E) int`
- `ToArray() []E`
- `AddAll(collection Collection[E]) bool`
- `AddAllBatch(elements []E) bool`
- `RemoveAll(collection Collection[E]) bool`
- `RemoveAllBatch(elements []E) bool`
- `RetainAll(collection Collection[E]) bool`
- `SubList(fromIndex int, toIndex int) (List[E], error)`
- `Sort(comparator Comparator[E])`
- `Reversed() Collection[E]`
- `Clone() List[E]`
- `Shuffle()`
- `FindFirst(predicate func(E) bool) (*E, error)`
- `FindAll(predicate func(E) bool) []E`
- `Filter(predicate func(E) bool) List[E]`
- `ForEach(consumer func(E))`
- `AddAllFirst(elements []E) bool`
- `AddAllLast(elements []E) bool`
- `RemoveIf(predicate func(E) bool) bool`
- `ReplaceAll(operator func(E) E)`
- `RemoveRange(fromIndex int, toIndex int) error`

#### LinkedList
A doubly-linked list implementation with efficient insertion and deletion operations.

**Features:**
- O(1) insertion/deletion at both ends
- O(n) random access
- Thread-safe operations
- Batch operations support
- Queue and Deque operations
- Stack operations

**Key Methods:**
- All ArrayList methods plus:
- `AddFirst(element E)`
- `AddLast(element E)`
- `RemoveFirst() (*E, error)`
- `RemoveLast() (*E, error)`
- `GetFirst() (*E, error)`
- `GetLast() (*E, error)`
- `RemoveFirstOccurrence(val E) bool`
- `RemoveLastOccurrence(val E) bool`
- `DescendingIterator() Iterator[E]`
- `Element() (*E, error)`
- `Offer(val E) bool`
- `Peek() (*E, error)`
- `Poll() (*E, error)`
- `OfferFirst(val E) bool`
- `OfferLast(val E) bool`
- `PeekFirst() (*E, error)`
- `PeekLast() (*E, error)`
- `PollFirst() (*E, error)`
- `PollLast() (*E, error)`
- `Pop() (*E, error)`
- `Push(val E)`

### Stack
A LIFO (Last-In-First-Out) stack implementation.

**Features:**
- O(1) push and pop operations
- Thread-safe operations
- Type-safe implementation
- Error handling for empty stack operations

**Key Methods:**
- `Push(element E)`
- `Pop() (*E, error)`
- `Peek() (*E, error)`
- `IsEmpty() bool`
- `Size() int`
- `Clear()`
- `Contains(element E) bool`
- `AddAll(collection Collection[E]) bool`

### PriorityQueue
A priority queue implementation based on a binary heap.

**Features:**
- O(log n) insertion and removal
- O(1) peek operation
- Thread-safe operations
- Customizable comparator for ordering
- Automatic capacity management
- Batch operations support
- Type-safe implementation
- Error handling for empty queue operations

**Key Methods:**
- `Add(element E) bool`
- `Offer(element E) bool`
- `Poll() (*E, error)`
- `Peek() (*E, error)`
- `Element() (*E, error)`
- `RemoveHead() (*E, error)`
- `Size() int`
- `IsEmpty() bool`
- `Clear()`
- `Contains(element E) bool`
- `Remove(element E) bool`
- `AddAll(collection Collection[E]) bool`
- `RemoveAll(collection Collection[E]) bool`
- `ContainsAll(collection Collection[E]) (bool, error)`
- `Equals(collection Collection[E]) bool`
- `Iterator() Iterator[E]`
- `ToArray() []E`
- `ToArrayWithType(arrayType reflect.Type) (interface{}, error)`
- `RemoveIf(predicate func(E) bool) bool`
- `ForEach(action func(E))`
- `GetComparator() Comparator[E]`

**Factory Methods:**
- `NewPriorityQueue[E comparable](comparator Comparator[E]) Queue[E]`
- `NewPriorityQueueWithCapacity[E comparable](initialCapacity int, comparator Comparator[E]) Queue[E]`
- `NewPriorityQueueFromCollection[E comparable](collection Collection[E], comparator Comparator[E]) Queue[E]`
- `NewPriorityQueueFromSortedSet[E constraints.Ordered](sortedSet SortedSet[E]) Queue[E]`

### TreeMap
A Red-Black tree-based map implementation that maintains sorted order of keys.

**Features:**
- O(log n) insertion, deletion, and lookup operations
- Thread-safe operations
- Type-safe implementation
- Comprehensive test coverage, including edge cases and balancing logic

**Key Methods:**
- `Put(key K, value V) V`
- `Get(key K) *V`
- `Remove(key K) V`
- `ContainsKey(key K) bool`
- `ContainsValue(value V) bool`
- `Size() int`
- `IsEmpty() bool`
- `Clear()`
- `FirstKey() (*K, error)`
- `LastKey() (*K, error)`
- `LowerKey(key K) (*K, error)`
- `HigherKey(key K) (*K, error)`
- `EntrySet() collections.Set[collections.MapEntry[K, V]]`
- `KeySet() collections.Set[K]`
- `Values() collections.Collection[V]`
- `Equals(other any) bool`
- `PutIfAbsent(key K, value V) V`
- `Replace(key K, value V) V`
- `ReplaceKeyWithValue(key K, oldValue V, newValue V) bool`
- `RemoveKeyWithValue(key K, value V) bool`
- `PutAll(m collections.Map[K, V])`

#### LinkedHashMap
A hash table and linked list implementation of the Map interface that maintains insertion order.

**Features:**
- O(1) average time complexity for basic operations
- Maintains insertion order of key-value pairs
- Thread-safe operations using read-write mutex locks
- Type-safe implementation using generics
- Predictable iteration order
- Automatic resizing
- Deadlock prevention through minimal lock duration
- Comprehensive test coverage

**Key Methods:**
- `Put(key K, value V) V`
- `Get(key K) *V`
- `GetOrDefault(key K, defaultValue V) V`
- `PutIfAbsent(key K, value V) V`
- `Remove(key K) V`
- `RemoveKeyWithValue(key K, value V) bool`
- `Replace(key K, value V) V`
- `ReplaceKeyWithValue(key K, oldValue V, newValue V) bool`
- `ComputeIfAbsent(key K, mappingFunction func(K) V) (V, error)`
- `PutAll(m collections.Map[K, V])`
- `Clear()`
- `Size() int`
- `IsEmpty() bool`
- `HasKey(key K) bool`
- `HasValue(value V) bool`
- `EntrySet() collections.Set[collections.MapEntry[K, V]]`
- `KeySet() collections.Set[K]`
- `Values() collections.Collection[V]`
- `Equals(obj any) bool`
- `ForEachEntry(action func(key K, value V))`

**Use Cases:**
- When you need a map that remembers the order in which keys were inserted
- LRU (Least Recently Used) cache implementations
- Maintaining order in configuration maps
- Any scenario where insertion order matters

### Sets

#### HashSet
A hash table-based set implementation that provides O(1) average time complexity for basic operations.

**Features:**
- O(1) average time complexity for add, remove, and contains operations
- Thread-safe operations using mutex locks
- Type-safe implementation using generics
- No duplicate elements allowed
- No ordering guarantees
- Automatic resizing

**Key Methods:**
- `Add(element E) bool`
- `Remove(element E) bool`
- `Contains(element E) bool`
- `Size() int`
- `IsEmpty() bool`
- `Clear()`
- `AddAll(collection Collection[E]) bool`
- `RemoveAll(collection Collection[E]) bool`
- `RetainAll(collection Collection[E]) bool`
- `ContainsAll(collection Collection[E]) (bool, error)`
- `Equals(collection Collection[E]) bool`
- `ToArray() []E`
- `Iterator() Iterator[E]`

#### LinkedHashSet
A hash table and linked list implementation of the Set interface, with predictable iteration order.

**Features:**
- O(1) average time complexity for add, remove, and contains operations
- Maintains insertion order
- Thread-safe operations using mutex locks
- Type-safe implementation using generics
- No duplicate elements allowed
- Predictable iteration order
- Automatic resizing

**Key Methods:**
- All HashSet methods plus:
- `GetFirst() (*E, error)`
- `GetLast() (*E, error)`
- `RemoveFirst() (*E, error)`
- `RemoveLast() (*E, error)`

## Error Handling

The library uses custom error codes for better error handling and identification. Common error codes include:

- `IndexOutOfBoundsError`: When accessing an index outside the valid range
- `NoSuchElementError`: When trying to access elements from an empty collection
- `NullPointerError`: When passing nil values to methods that don't accept them
- `EmptyStackError`: When performing operations on an empty stack
- `EmptyQueueError`: When performing operations on an empty queue

## Usage

### ArrayList Example
```go
// Create a new ArrayList
list := lists.NewArrayList[int]()

// Add elements
list.Add(1)
list.Add(2)
list.Add(3)

// Add elements in batch
list.AddAllBatch([]int{4, 5, 6})

// Remove elements
list.Remove(3)
list.RemoveAtIndex(0)

// Get elements
element, err := list.Get(0)
if err != nil {
    // Handle error
}

// Check if list contains an element
if list.Contains(5) {
    // Do something
}

// Get the size of the list
size := list.Size()

// Clear the list
list.Clear()
```

### LinkedList Example
```go
// Create a new LinkedList
list := lists.NewLinkedList[int]()

// Add elements
list.Add(1)
list.Add(2)
list.Add(3)

// Add elements at specific positions
list.AddFirst(0)
list.AddLast(4)

// Remove elements
list.RemoveFirst()
list.RemoveLast()

// Get elements
first, err := list.GetFirst()
if err != nil {
    // Handle error
}

// Use as a queue
list.Offer(5)
element, err := list.Poll()

// Use as a stack
list.Push(6)
element, err = list.Pop()
```

### Stack Example
```go
// Create a new Stack
stack := lists.NewStack[int]()

// Push elements
stack.Push(1)
stack.Push(2)
stack.Push(3)

// Pop elements
element, err := stack.Pop()
if err != nil {
    // Handle error
}

// Peek at the top element
top, err := stack.Peek()
if err != nil {
    // Handle error
}
```

### PriorityQueue Example
```go
// Create a new PriorityQueue with a custom comparator
comparator := &IntComparator[int]{}
pq := queues.NewPriorityQueue[int](comparator)

// Add elements
pq.Add(5)
pq.Add(3)
pq.Add(7)

// Poll elements (they will come out in order)
element, err := pq.Poll()
if err != nil {
    // Handle error
}

// Peek at the next element without removing it
next, err := pq.Peek()
if err != nil {
    // Handle error
}

// Remove specific elements
pq.Remove(3)

// Check if queue contains an element
if pq.Contains(5) {
    // Do something
}
```

### LinkedHashMap Example
```go
// Create a new LinkedHashMap
lhm := maps.NewLinkedHashMap[string, int]()

// Add elements (order will be maintained)
lhm.Put("first", 1)
lhm.Put("second", 2)
lhm.Put("third", 3)

// Get elements
value := lhm.Get("second")
if value != nil {
    // Use the value
}

// Put if absent (won't overwrite existing value)
oldValue := lhm.PutIfAbsent("first", 10)
// oldValue will be 1, and the map still contains {"first": 1}

// Compute if absent (creates value if key doesn't exist)
newValue, err := lhm.ComputeIfAbsent("fourth", func(k string) int {
    return len(k) * 10
})
if err == nil {
    // newValue will be 60 (len("fourth") * 10)
}

// Remove elements
removedValue := lhm.Remove("second")

// Remove with value check
removed := lhm.RemoveKeyWithValue("first", 1)

// Iterate in insertion order
lhm.ForEachEntry(func(key string, value int) {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
})

// Get all keys in insertion order
keys := lhm.KeySet()

// Get all values in insertion order
values := lhm.Values()

// Check if map contains key or value
if lhm.HasKey("third") {
    // Key exists
}

if lhm.HasValue(3) {
    // Value exists
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Inspired by Java's Collections Framework
- Built with Go's generics support
- Thread-safe implementations using sync package

## Thread Safety

All collections in this library are thread-safe by default, using mutex locks for concurrent access. The implementation details are as follows:

### Read-Write Locks
- Read operations (e.g., `Get`, `Contains`, `Size`) use `RLock()` for better concurrency
- Write operations (e.g., `Add`, `Remove`, `Clear`) use `Lock()` for exclusive access
- Nested locks are avoided to prevent deadlocks
- Lock granularity is optimized for common use cases

### Concurrent Operations
- Batch operations are atomic
- Iterator operations are protected by read locks
- Collection views (e.g., `EntrySet`, `KeySet`) maintain thread safety
- Concurrent modifications are detected and handled appropriately

## Performance Characteristics

### Time Complexity
- **ArrayList**: O(1) random access, O(n) insertion/deletion
- **LinkedList**: O(1) insertion/deletion at ends, O(n) random access
- **HashSet/LinkedHashSet**: O(1) average case for basic operations
- **TreeMap**: O(log n) for all operations
- **LinkedHashMap**: O(1) average case for basic operations
- **PriorityQueue**: O(log n) insertion/removal, O(1) peek
- **Stack**: O(1) for all operations

### Space Complexity
- **ArrayList**: O(n) with automatic resizing
- **LinkedList**: O(n) with constant overhead per element
- **HashSet/LinkedHashSet**: O(n) with load factor considerations
- **TreeMap**: O(n) with balanced tree structure
- **LinkedHashMap**: O(n) with hash table and linked list overhead
- **PriorityQueue**: O(n) with heap structure
- **Stack**: O(n) with dynamic resizing

## Best Practices

### Choosing the Right Collection
- Use **ArrayList** for random access and when size is known
- Use **LinkedList** for frequent insertions/deletions at ends
- Use **HashSet** for unique elements with no ordering requirements
- Use **LinkedHashSet** for unique elements with insertion order
- Use **TreeMap** for sorted key-value pairs
- Use **LinkedHashMap** for key-value pairs with insertion order
- Use **PriorityQueue** for priority-based processing
- Use **Stack** for LIFO operations

### Thread Safety Considerations
- Prefer read operations when possible for better concurrency
- Batch operations when modifying collections
- Use appropriate collection views for iteration
- Consider using sync.Map for highly concurrent scenarios
- Avoid nested locks and long-held locks

### Memory Management
- Use appropriate initial capacities to avoid resizing
- Clear collections when no longer needed
- Consider using object pools for frequently created collections
- Monitor memory usage in long-running applications

### Development Setup
```bash
# Clone the repository
git clone https://github.com/chiranjeevipavurala/gocollections.git

# Install dependencies
go mod download

# Run tests
go test ./...

# Run benchmarks
go test -bench=. ./...
```

### Code Style
- Follow Go's standard formatting guidelines
- Use meaningful variable and function names
- Add comments for complex logic
- Write comprehensive tests
- Update documentation for new features

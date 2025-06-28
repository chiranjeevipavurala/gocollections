# Go Collections Benchmarking Suite

This directory contains comprehensive benchmarks for all data structures in the Go Collections library. The benchmarks help you understand the performance characteristics of different data structures and choose the right one for your use case.

## What is Benchmarking?

Benchmarking is the process of measuring and comparing the performance of different implementations. In Go, benchmarking helps us:

- **Measure Performance**: How fast operations execute
- **Compare Implementations**: Which data structure is better for specific use cases
- **Identify Bottlenecks**: Find performance issues in code
- **Optimize Code**: Make informed decisions about improvements
- **Validate Changes**: Ensure optimizations actually improve performance

## Running Benchmarks

### Quick Start

```bash
# Run all benchmarks
go test -bench=. ./benchmark/

# Run specific benchmark categories
go test -bench=BenchmarkArrayList ./benchmark/
go test -bench=BenchmarkHashMap ./benchmark/
go test -bench=BenchmarkPriorityQueue ./benchmark/

# Run with more iterations for better accuracy
go test -bench=. -benchtime=5s ./benchmark/

# Run with memory allocation profiling
go test -bench=. -benchmem ./benchmark/
```

### Running Individual Benchmark Files

```bash
# List benchmarks
go test -bench=BenchmarkArrayListAdd ./benchmark/

# Run specific benchmark
go test -bench=BenchmarkArrayListAdd -benchtime=1s ./benchmark/
```

### Benchmark Output Explanation

```
BenchmarkArrayListAdd-8         1000000              1234 ns/op
BenchmarkLinkedListAdd-8         500000              2468 ns/op
```

- **BenchmarkArrayListAdd-8**: Benchmark name and number of CPU cores
- **1000000**: Number of iterations
- **1234 ns/op**: Average time per operation (nanoseconds)
- **2468 ns/op**: Average time per operation for comparison

## Benchmark Categories

### 📋 List Benchmarks (`list_benchmarks.go`)

Compares ArrayList vs LinkedList performance:

- **Add Operations**: O(1) for both, but different constant factors
- **Get Operations**: O(1) ArrayList vs O(n) LinkedList
- **Insert at Index**: Performance varies by position
- **Remove Operations**: Performance varies by position
- **Contains Operations**: Linear search for both

**Key Insights:**
- ArrayList is faster for random access
- LinkedList is faster for insertions/deletions at ends
- ArrayList has better cache locality

### 🔗 Set Benchmarks (`set_benchmarks.go`)

Compares HashSet vs LinkedHashSet performance:

- **Add Operations**: O(1) average case for both
- **Contains Operations**: O(1) average case for both
- **Remove Operations**: O(1) average case for both
- **Iterator Performance**: LinkedHashSet maintains order

**Key Insights:**
- HashSet is slightly faster for basic operations
- LinkedHashSet maintains insertion order
- Both have similar performance characteristics

### 🗺️ Map Benchmarks (`map_benchmarks.go`)

Compares HashMap vs LinkedHashMap vs TreeMap performance:

- **Put Operations**: O(1) HashMap/LinkedHashMap vs O(log n) TreeMap
- **Get Operations**: O(1) HashMap/LinkedHashMap vs O(log n) TreeMap
- **Remove Operations**: O(1) HashMap/LinkedHashMap vs O(log n) TreeMap
- **Special Operations**: PutIfAbsent, ComputeIfAbsent, ForEachEntry

**Key Insights:**
- HashMap is fastest for basic operations
- LinkedHashMap maintains insertion order with minimal overhead
- TreeMap provides sorted keys but slower operations

### 📊 Queue Benchmarks (`queue_benchmarks.go`)

Compares PriorityQueue, Stack, and LinkedList as Queue:

- **Add/Offer Operations**: Performance varies by implementation
- **Poll/Pop Operations**: Performance varies by implementation
- **Peek Operations**: O(1) for all implementations
- **Priority Operations**: PriorityQueue maintains heap property

**Key Insights:**
- PriorityQueue is slower but provides priority ordering
- Stack and LinkedList are fast for basic queue operations
- Choose based on ordering requirements

## Performance Characteristics Summary

### Time Complexity Comparison

| Operation | ArrayList | LinkedList | HashSet | LinkedHashSet | HashMap | LinkedHashMap | TreeMap | PriorityQueue |
|-----------|-----------|------------|---------|---------------|---------|---------------|---------|---------------|
| Add/Put   | O(1)      | O(1)       | O(1)    | O(1)          | O(1)    | O(1)          | O(log n)| O(log n)      |
| Get       | O(1)      | O(n)       | O(1)    | O(1)          | O(1)    | O(1)          | O(log n)| O(1)          |
| Remove    | O(n)      | O(n)       | O(1)    | O(1)          | O(1)    | O(1)          | O(log n)| O(n)          |
| Contains  | O(n)      | O(n)       | O(1)    | O(1)          | O(1)    | O(1)          | O(log n)| O(n)          |

### Memory Usage

- **ArrayList**: O(n) with automatic resizing
- **LinkedList**: O(n) with constant overhead per element
- **HashSet/LinkedHashSet**: O(n) with load factor considerations
- **HashMap/LinkedHashMap**: O(n) with hash table overhead
- **TreeMap**: O(n) with balanced tree structure
- **PriorityQueue**: O(n) with heap structure

## Choosing the Right Data Structure

### Use ArrayList when:
- You need random access to elements
- You know the approximate size in advance
- You perform mostly read operations
- Cache locality is important

### Use LinkedList when:
- You frequently insert/delete at the beginning or end
- You need to use it as a queue or stack
- You don't need random access
- Memory overhead is not a concern

### Use HashSet when:
- You need unique elements
- Order doesn't matter
- You need fast contains operations
- You don't need to maintain insertion order

### Use LinkedHashSet when:
- You need unique elements
- You need to maintain insertion order
- You need predictable iteration order
- You're willing to pay a small performance cost

### Use HashMap when:
- You need key-value pairs
- Order doesn't matter
- You need fast get/put operations
- You don't need sorted keys

### Use LinkedHashMap when:
- You need key-value pairs
- You need to maintain insertion order
- You need predictable iteration order
- You're willing to pay a small performance cost

### Use TreeMap when:
- You need sorted keys
- You need range queries
- You need ordered iteration
- You're willing to pay O(log n) performance cost

### Use PriorityQueue when:
- You need priority-based processing
- You need to always get the highest/lowest priority element
- You need heap operations
- You're willing to pay O(log n) performance cost

## Benchmark Configuration

The benchmarks use different data sizes to test various scenarios:

- **SmallSize (100)**: Tests with small collections
- **MediumSize (1000)**: Tests with medium collections
- **LargeSize (10000)**: Tests with large collections

## Tips for Running Benchmarks

1. **Run multiple times**: Performance can vary due to system load
2. **Use consistent environment**: Close other applications
3. **Warm up the system**: Run a few iterations first
4. **Use appropriate benchtime**: Longer runs give more accurate results
5. **Monitor system resources**: CPU, memory, and disk usage

## Contributing

When adding new data structures or operations:

1. Add benchmark functions to the appropriate file
2. Follow the naming convention: `Benchmark<Structure><Operation>`
3. Include both small and large data size tests
4. Add comparisons with similar data structures
5. Update this README with new insights

## Example Benchmark Results

```
BenchmarkArrayListAdd-8         1000000              1234 ns/op
BenchmarkLinkedListAdd-8         500000              2468 ns/op
BenchmarkHashSetAdd-8            800000              1567 ns/op
BenchmarkLinkedHashSetAdd-8      600000              1890 ns/op
BenchmarkHashMapPut-8            900000              1345 ns/op
BenchmarkLinkedHashMapPut-8      700000              1678 ns/op
BenchmarkPriorityQueueAdd-8      300000              4123 ns/op
```

These results show that:
- ArrayList is ~2x faster than LinkedList for add operations
- HashSet is ~20% faster than LinkedHashSet
- HashMap is ~25% faster than LinkedHashMap
- PriorityQueue is significantly slower due to heap maintenance 
#!/bin/bash

# Go Collections Benchmark Runner
# This script runs comprehensive benchmarks for all data structures

echo "🚀 Go Collections Benchmark Suite"
echo "=================================="
echo ""

# Function to run benchmarks with nice formatting
run_benchmark() {
    local name="$1"
    local pattern="$2"
    local description="$3"
    
    echo "📊 $name"
    echo "   $description"
    echo "   Running: go test -bench=$pattern -benchtime=1s ./benchmark/"
    echo ""
    
    go test -bench="$pattern" -benchtime=1s ./benchmark/
    echo ""
    echo "----------------------------------------"
    echo ""
}

# Function to run memory benchmarks
run_memory_benchmark() {
    local name="$1"
    local pattern="$2"
    
    echo "💾 $name (Memory Allocation)"
    echo "   Running: go test -bench=$pattern -benchmem -benchtime=1s ./benchmark/"
    echo ""
    
    go test -bench="$pattern" -benchmem -benchtime=1s ./benchmark/
    echo ""
    echo "----------------------------------------"
    echo ""
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed or not in PATH"
    exit 1
fi

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ Error: Please run this script from the project root directory"
    exit 1
fi

echo "🔧 Environment:"
echo "   Go version: $(go version)"
echo "   Architecture: $(go env GOARCH)"
echo "   OS: $(go env GOOS)"
echo "   CPU cores: $(nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo "unknown")"
echo ""

# Run all benchmarks
echo "🎯 Running All Benchmarks..."
go test -bench=. -benchtime=1s ./benchmark/
echo ""

echo "📋 Detailed Benchmark Categories:"
echo ""

# List benchmarks
run_benchmark "ArrayList Benchmarks" "BenchmarkArrayList" "Dynamic array implementation with O(1) random access"

# LinkedList benchmarks
run_benchmark "LinkedList Benchmarks" "BenchmarkLinkedList" "Doubly-linked list with O(1) insertion/deletion at ends"

# Set benchmarks
run_benchmark "HashSet Benchmarks" "BenchmarkHashSet" "Hash table-based set with O(1) average operations"

# LinkedHashSet benchmarks
run_benchmark "LinkedHashSet Benchmarks" "BenchmarkLinkedHashSet" "Hash set that maintains insertion order"

# Map benchmarks
run_benchmark "HashMap Benchmarks" "BenchmarkHashMap" "Hash table-based map with O(1) average operations"

# LinkedHashMap benchmarks
run_benchmark "LinkedHashMap Benchmarks" "BenchmarkLinkedHashMap" "Hash map that maintains insertion order"

# TreeMap benchmarks
run_benchmark "TreeMap Benchmarks" "BenchmarkTreeMap" "Red-Black tree-based map with O(log n) operations"

# Queue benchmarks
run_benchmark "PriorityQueue Benchmarks" "BenchmarkPriorityQueue" "Heap-based priority queue with O(log n) operations"

# Stack benchmarks
run_benchmark "Stack Benchmarks" "BenchmarkStack" "LIFO stack implementation"

# Memory allocation benchmarks
echo "💾 Memory Allocation Analysis:"
echo ""

run_memory_benchmark "ArrayList Memory" "BenchmarkArrayList"
run_memory_benchmark "HashMap Memory" "BenchmarkHashMap"
run_memory_benchmark "PriorityQueue Memory" "BenchmarkPriorityQueue"

echo "✅ Benchmark suite completed!"
echo ""
echo "📈 Tips for interpreting results:"
echo "   - Lower ns/op = faster performance"
echo "   - Lower B/op = less memory allocation"
echo "   - Lower allocs/op = fewer allocations"
echo ""
echo "🔍 For more detailed analysis, run:"
echo "   go test -bench=. -benchmem -benchtime=5s ./benchmark/ > benchmark_results.txt"
echo ""
echo "📖 See benchmark/README.md for detailed explanations and performance characteristics." 
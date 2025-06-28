# Benchmark Results

This directory contains benchmark results for the Go Collections library. Results are automatically generated and saved here for tracking performance over time.

## File Naming Convention

Results are saved with the following naming pattern:
```
benchmark_results_<OS>_<ARCH>_<TIMESTAMP>.txt
```

For example:
- `benchmark_results_Darwin_arm64_2024-01-15_14-30-25.txt`
- `benchmark_results_Linux_x86_64_2024-01-15_14-30-25.txt`

## How to Generate Results

### Local Development
```bash
# Run the save results script
./benchmark/save_results.sh

# Or run benchmarks manually
go test -bench=. -benchtime=2s ./benchmark/ > benchmark_results.txt
go test -bench=. -benchmem -benchtime=2s ./benchmark/ > benchmark_memory_results.txt
```

### CI/CD (GitHub Actions)
Benchmarks are automatically run on:
- Every push to main/master branch
- Every pull request
- Manual workflow dispatch

Results are saved as GitHub Actions artifacts and can be downloaded from the Actions tab.

## Interpreting Results

### Performance Metrics
- **ns/op**: Nanoseconds per operation (lower is better)
- **B/op**: Bytes allocated per operation (lower is better)
- **allocs/op**: Number of allocations per operation (lower is better)

### Performance Categories
- **Excellent**: < 20 ns/op
- **Good**: 20-50 ns/op
- **Acceptable**: 50-200 ns/op
- **Slow**: > 200 ns/op (may be expected for O(n) operations)

### Key Performance Insights

#### ArrayList vs LinkedList
- **ArrayList.Get**: ~14 ns/op (O(1) - excellent)
- **LinkedList.Get**: ~1000 ns/op (O(n) - expected)

#### HashMap vs LinkedHashMap
- **HashMap.Get**: ~33 ns/op
- **LinkedHashMap.Get**: ~28 ns/op (minimal overhead!)

#### HashSet vs LinkedHashSet
- **HashSet.Add**: ~160 ns/op
- **LinkedHashSet.Add**: ~231 ns/op (~44% overhead for ordering)

## Performance Trends

### Recent Performance Highlights
- LinkedHashMap shows minimal overhead for maintaining insertion order
- ArrayList random access is extremely fast (~14 ns/op)
- PriorityQueue operations are very efficient (~14-21 ns/op)
- TreeMap provides good performance for sorted data (~36-42 ns/op)

### Recommendations
1. **Use ArrayList** for random access and when size is known
2. **Use LinkedList** for frequent insertions/deletions at ends
3. **Use HashMap** for maximum performance without ordering
4. **Use LinkedHashMap** when you need insertion order (minimal overhead)
5. **Use TreeMap** when you need sorted keys
6. **Use PriorityQueue** for priority-based processing

## Comparing Results

### Between Different Systems
Results may vary significantly between different:
- CPU architectures (ARM64 vs x86_64)
- Operating systems (macOS vs Linux vs Windows)
- Go versions
- System load and available resources

### Over Time
Compare results from different timestamps to:
- Track performance improvements
- Identify performance regressions
- Validate optimization efforts

## Latest Results

The `latest_results.txt` file is a symlink to the most recent benchmark results.

```bash
# View latest results
cat latest_results.txt

# Compare with previous results
diff benchmark_results_*.txt
```

## Contributing

When adding new data structures or optimizations:
1. Run benchmarks before and after changes
2. Save results using `./benchmark/save_results.sh`
3. Document any significant performance changes
4. Update this README with new insights

## Performance Goals

### Target Performance (ns/op)
- **O(1) operations**: < 50 ns/op
- **O(log n) operations**: < 100 ns/op
- **O(n) operations**: < 1000 ns/op (for small n)
- **Memory allocation**: Minimize B/op and allocs/op

### Current Status
✅ All O(1) operations meet targets
✅ O(log n) operations meet targets
✅ Memory usage is reasonable
✅ Thread safety overhead is minimal 
# Contributing Guidelines

## Error Handling Guidelines

### 1. Redundant Error Checks
- Avoid redundant error checks when the underlying method already handles the error cases
- Example of redundant check:
```go
// Don't do this
val, err := list.RemoveAtIndex(index)
if err != nil {
    return nil, err
}
return val, nil

// Do this instead
val, _ := list.RemoveAtIndex(index)
return val, nil
```

### 2. Error Check Elimination Criteria
A redundant error check can be eliminated if ALL of these conditions are met:
1. The method is called within a mutex lock
2. The index/position is guaranteed to be valid
3. The underlying method already handles the error case
4. The operation cannot fail due to the current state

### 3. Required Error Checks
Always keep error checks for:
1. Input validation (nil checks, bounds checks)
2. State validation (empty collection, invalid state)
3. External operations (I/O, network calls)
4. Operations that can fail independently of the current state

## Thread Safety Guidelines

### 1. Mutex Usage
- Use `sync.RWMutex` for read-write operations
- Use `sync.Mutex` for write-only operations
- Always use `defer` with mutex unlocks
- Keep critical sections as small as possible

### 2. Lock Ordering
- Always acquire locks in a consistent order to prevent deadlocks
- Document lock ordering in complex operations
- Avoid nested locks when possible

### 3. Thread-Safe Operations
- All public methods must be thread-safe
- Internal methods can assume the caller holds the appropriate lock
- Document thread-safety guarantees in method comments

### 4. Deadlock Prevention
- Never acquire multiple locks without a consistent ordering
- Common deadlock scenarios to avoid:
  ```go
  // DON'T: Inconsistent lock ordering
  func (a *TypeA) Method1(b *TypeB) {
      a.mu.Lock()
      b.mu.Lock()  // Potential deadlock if b.Method2(a) is called
      // ...
  }
  
  func (b *TypeB) Method2(a *TypeA) {
      b.mu.Lock()
      a.mu.Lock()  // Deadlock with Method1
      // ...
  }
  
  // DO: Consistent lock ordering
  func (a *TypeA) Method1(b *TypeB) {
      a.mu.Lock()
      b.mu.Lock()
      // ...
  }
  
  func (b *TypeB) Method2(a *TypeA) {
      a.mu.Lock()  // Always lock A before B
      b.mu.Lock()
      // ...
  }
  ```
- Best practices:
  1. Document lock ordering in method comments
  2. Use lock ordering constants for complex types
  3. Consider using `sync.Map` for simple key-value operations
  4. Use `RLock()` when possible instead of `Lock()`
  5. Keep critical sections as small as possible
  6. Avoid calling external methods while holding locks
  7. Use timeouts for lock acquisition when appropriate
- Lock hierarchy:
  1. Define a clear hierarchy of locks
  2. Always acquire locks in the same order
  3. Document the hierarchy in package documentation
  4. Use lock ordering verification tools
- Common pitfalls:
  1. Recursive locking (same mutex)
  2. Inconsistent lock ordering
  3. Holding locks during I/O operations
  4. Locking during callbacks
  5. Nested locks without clear ordering

## Code Organization

### 1. Method Ordering
1. Constructor methods
2. Public interface methods
3. Private helper methods
4. Internal utility methods

### 2. Method Documentation
- Document thread-safety guarantees
- Document error conditions
- Document preconditions and postconditions
- Include examples for complex operations

### 3. Error Codes
- Use predefined error codes from `errorcodes` package
- Document new error codes when added
- Keep error messages consistent

## Testing Guidelines

### 1. Test Coverage
- Test all public methods
- Test error conditions
- Test thread-safety
- Test edge cases

### 2. Test Organization
- Group related tests together
- Use descriptive test names
- Include setup and teardown when needed
- Test both success and failure cases

### 3. Concurrent Testing
- Test concurrent operations
- Test race conditions
- Test deadlock prevention
- Test performance under load

### 4. Code Coverage Requirements
- Aim for 100% line coverage for critical paths
- Minimum coverage requirements:
  - 100% for public interface methods
  - 100% for error handling paths
  - 100% for thread-safety critical sections
  - 90% for utility/helper methods
- Coverage exclusions:
  - Generated code
  - Debug-only code
  - Platform-specific code
  - Unreachable code paths
- Coverage verification:
  - Run coverage analysis before submitting PRs
  - Document any intentional coverage gaps
  - Justify any coverage below requirements
- Coverage tools:
  - Use `go test -cover` for basic coverage
  - Use `go test -coverprofile` for detailed analysis
  - Use `go tool cover` for coverage visualization
- Coverage reporting:
  - Include coverage report in PRs
  - Track coverage trends over time
  - Set up CI/CD coverage checks

## Performance Guidelines

### 1. Memory Management
- Pre-allocate slices when size is known
- Use appropriate initial capacities
- Avoid unnecessary allocations
- Document memory usage patterns

### 2. Algorithm Complexity
- Document time complexity
- Document space complexity
- Optimize for common cases
- Consider trade-offs between memory and speed

### 3. Resource Management
- Release resources promptly
- Use defer for cleanup
- Handle resource exhaustion
- Document resource requirements

## Code Review Checklist

1. Error Handling
   - [ ] No redundant error checks
   - [ ] All error cases handled
   - [ ] Error messages are clear
   - [ ] Error codes are appropriate

2. Thread Safety
   - [ ] Mutex usage is correct
   - [ ] No deadlock potential
   - [ ] Critical sections are minimal
   - [ ] Thread-safety is documented

3. Code Organization
   - [ ] Methods are properly ordered
   - [ ] Documentation is complete
   - [ ] Code is readable
   - [ ] Naming is consistent

4. Testing
   - [ ] Tests cover all cases
   - [ ] Concurrent tests included
   - [ ] Edge cases tested
   - [ ] Performance tested

5. Performance
   - [ ] No unnecessary allocations
   - [ ] Complexity is documented
   - [ ] Resources are managed
   - [ ] Optimizations are justified

## Java Collections Framework Compatibility

### 1. Interface Hierarchy
- Follow Java's collection interface hierarchy:
  - `Collection` → `List`/`Set`/`Queue`
  - `List` → `ArrayList`/`LinkedList`
  - `Set` → `HashSet`/`TreeSet`
  - `Queue` → `PriorityQueue`/`BlockingQueue`
- Maintain consistent method signatures with Java equivalents
- Preserve Java's method behavior and error handling patterns

### 2. Method Naming and Behavior
- Use Java's method names where applicable:
  - `Add`/`Remove`/`Contains` for basic operations
  - `AddAll`/`RemoveAll`/`RetainAll` for bulk operations
  - `Iterator`/`ForEach` for iteration
- Maintain Java's method semantics:
  - Return values should match Java's behavior
  - Error conditions should be similar
  - Thread safety guarantees should be equivalent

### 3. Implementation Patterns
- Follow Java's implementation patterns:
  - Use similar internal data structures
  - Maintain comparable performance characteristics
  - Preserve Java's memory usage patterns
- Document any deviations from Java's behavior
- Justify any Go-specific optimizations

### 4. Error Handling
- Map Java exceptions to Go errors:
  - `NullPointerException` → `NullPointerError`
  - `IndexOutOfBoundsException` → `IndexOutOfBoundsError`
  - `IllegalArgumentException` → `IllegalArgumentError`
  - `NoSuchElementException` → `NoSuchElementError`
- Maintain similar error conditions and messages
- Document any Go-specific error cases

### 5. Thread Safety
- Match Java's thread safety guarantees:
  - `ArrayList` → Thread-safe with explicit synchronization
  - `Vector` → Thread-safe by default
  - `ConcurrentHashMap` → Thread-safe with fine-grained locking
- Document thread safety guarantees clearly
- Use appropriate synchronization mechanisms

### 6. Performance Characteristics
- Maintain similar time complexity:
  - `ArrayList`: O(1) random access, O(n) insertion/deletion
  - `LinkedList`: O(n) random access, O(1) insertion/deletion
  - `HashSet`: O(1) average case operations
  - `TreeSet`: O(log n) operations
- Document any performance differences
- Justify any Go-specific optimizations

### 7. Testing Compatibility
- Test against Java's behavior:
  - Verify method return values
  - Check error conditions
  - Validate thread safety
  - Measure performance characteristics
- Document any behavioral differences
- Include Java compatibility tests 
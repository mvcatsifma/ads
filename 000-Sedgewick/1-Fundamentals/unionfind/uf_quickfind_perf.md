# Quick-Find Performance Analysis

## Benchmark Results

```
BenchmarkQuickFind_Find-12     1,000,000,000    0.4173 ns/op
BenchmarkQuickFind_Union-12    1,000,000        1,052 ns/op
```

**Test Environment:**
- CPU: Apple M3 Pro
- Architecture: arm64 (darwin)
- Array Size: 1,000 elements

## Find Operation

**Performance:** 0.4173 nanoseconds per operation

**Characteristics:**
- **Time Complexity:** O(1) - constant time
- **Operations per second:** ~2.4 billion
- **Implementation:** Direct array access `return u.id[p]`
- **Iterations:** 1 billion successful benchmark runs

**Analysis:**
The find operation is extremely fast because it's a single array lookup. No loops, no recursion - just direct memory access. This is the primary advantage of the Quick-Find algorithm.

## Union Operation

**Performance:** 1,052 nanoseconds per operation (1.052 microseconds)

**Characteristics:**
- **Time Complexity:** O(n) - linear time
- **Operations per second:** ~950,000
- **Cost per element:** ~1.05 ns/element (1,052 ns ÷ 1,000 elements)
- **Iterations:** 1 million successful benchmark runs

**Analysis:**
The union operation scans the entire array to update all elements in the component being merged. For an array of 1,000 elements, this means touching every element on each union operation.

## Performance Ratio

**Union vs Find:** 2,520x slower
- Find: 0.4173 ns
- Union: 1,052 ns
- Ratio: 1,052 ÷ 0.4173 ≈ 2,520

## Scaling Behavior

**Expected performance at different sizes:**

| Array Size | Find Time | Union Time | Union/Find Ratio |
|------------|-----------|------------|------------------|
| 100        | ~0.4 ns   | ~100 ns    | ~250x            |
| 1,000      | ~0.4 ns   | ~1,000 ns  | ~2,500x          |
| 10,000     | ~0.4 ns   | ~10,000 ns | ~25,000x         |
| 100,000    | ~0.4 ns   | ~100 μs    | ~250,000x        |
| 1,000,000  | ~0.4 ns   | ~1 ms      | ~2,500,000x      |

**Key Observation:** While find time remains constant, union time grows linearly with array size. The performance gap widens dramatically as the data structure grows.

## Use Case Suitability

**Quick-Find is optimal when:**
- Find operations vastly outnumber union operations
- The data structure is small (< 1,000 elements)
- Constant-time lookups are critical
- Union operations are rare

**Quick-Find is problematic when:**
- Many union operations are needed
- Working with large datasets (> 10,000 elements)
- Building the union-find structure requires many merges
- Total runtime is dominated by union operations

## Memory Characteristics

**Space Complexity:** O(n)
- Single integer array of size n
- No additional data structures
- Minimal memory overhead

**Cache Performance:**
- Find: Excellent (single cache line access)
- Union: Poor (scans entire array, many cache misses)

## Conclusion

Quick-Find demonstrates excellent find performance (0.4 ns) but suffers from poor union performance (1,052 ns for 1,000 elements). The O(n) union operation makes it impractical for applications requiring frequent merges, despite its O(1) find advantage.
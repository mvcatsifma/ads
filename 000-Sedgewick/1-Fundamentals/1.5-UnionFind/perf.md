# Weighted Quick-Union Benchmark Results Summary

## Performance Overview

```
BenchmarkWeightedQuickUnion_Find-12     1,000,000,000    0.8189 ns/op
BenchmarkWeightedQuickUnion_Union-12      322,218,870    3.428 ns/op
```

**Test Environment:**
- CPU: Apple M3 Pro
- Architecture: arm64 (darwin)
- Array Size: 10,000 elements (find), 1,000,000 elements (union)

## Understanding the Numbers

**Benchmark output format:** `[iterations] [time per operation]`

### Find Operation
- **Iterations run:** 1,000,000,000 (1 billion times)
- **Time per operation:** 0.8189 nanoseconds
- **Throughput:** 1,000,000,000 ns/sec ÷ 0.8189 ns/op = **1.22 billion ops/sec**

### Union Operation
- **Iterations run:** 322,218,870 (322 million times)
- **Time per operation:** 3.428 nanoseconds
- **Throughput:** 1,000,000,000 ns/sec ÷ 3.428 ns/op = **292 million ops/sec**

**Note:** The iteration count (322M) is how many times Go ran the benchmark loop to get stable timing. The throughput (292M ops/sec) is the actual performance metric.

## Results Analysis

### Find: 0.82 ns/op
- Extremely fast tree traversal
- Can perform **1.22 billion finds per second**

### Union: 3.43 ns/op
- Two finds + one link operation
- Can perform **292 million unions per second**
- Only **4.2x slower** than find (reasonable since union does 2 finds)

## Comparison Across Algorithms

| Algorithm | Find | Union | Find Speed | Union Speed |
|-----------|------|-------|------------|-------------|
| **Quick-Find** | 0.42 ns | 1,052 ns | 2.4B ops/sec | 950K ops/sec |
| **Quick-Union** | 6,293 ns | 847 ns | 159K ops/sec | 1.2M ops/sec |
| **Weighted Quick-Union** | 0.82 ns | 3.43 ns | 1.2B ops/sec | 292M ops/sec |

## Key Insights

### 1. Balanced Performance ✅
Weighted Quick-Union achieves **excellent balance**:
- Find is nearly as fast as Quick-Find (0.82 ns vs 0.42 ns)
- Union is **307x faster** than Quick-Find (3.43 ns vs 1,052 ns)
- Both operations are in the **nanosecond range**

### 2. Logarithmic Guarantee Working
Despite 10,000 sequential unions (worst case for unweighted), find remains **0.82 ns**:
- Tree height is kept at O(log n) ≈ log₂(10,000) ≈ 13-14 levels
- Compare to unweighted Quick-Union: 6,293 ns (**7,685x slower!**)

### 3. Union Efficiency
Union at **3.43 ns** is remarkably fast:
- Includes two find operations (~1.6 ns total)
- Plus size comparison and link (~1.8 ns)
- Minimal overhead

## Conclusion

Weighted Quick-Union delivers **best-of-both-worlds performance**:
- **0.82 ns find** (1.2B ops/sec) - nearly as fast as Quick-Find
- **3.43 ns union** (292M ops/sec) - 307x faster than Quick-Find
- **Logarithmic guarantee** - prevents degenerate O(n) trees

This makes it the **practical choice** for most union-find applications, offering excellent performance for both operations without the extreme trade-offs of Quick-Find or unweighted Quick-Union.
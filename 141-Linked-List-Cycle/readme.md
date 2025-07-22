## 141. Linked List Cycle

https://leetcode.com/problems/linked-list-cycle/

## Benchmark results

```text
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: linked-list-cycle
cpu: Apple M3 Pro
Benchmark_hasCycle/NoCycle_Size_10-12           416487948                2.754 ns/op           0 B/op          0 allocs/op
Benchmark_hasCycle/WithCycle_Size_10-12         412422508                2.891 ns/op           0 B/op          0 allocs/op
Benchmark_hasCycle/NoCycle_Size_100-12          16643155                71.57 ns/op            0 B/op          0 allocs/op
Benchmark_hasCycle/WithCycle_Size_100-12        16663638                71.58 ns/op            0 B/op          0 allocs/op
Benchmark_hasCycle/NoCycle_Size_1000-12          1213450               991.0 ns/op             0 B/op          0 allocs/op
Benchmark_hasCycle/WithCycle_Size_1000-12        1209652               990.2 ns/op             0 B/op          0 allocs/op
Benchmark_hasCycle/NoCycle_Size_10000-12           96322             10578 ns/op               0 B/op          0 allocs/op
Benchmark_hasCycle/WithCycle_Size_10000-12        110558             10976 ns/op               0 B/op          0 allocs/op
PASS
ok      linked-list-cycle       12.522s
```

Let's break down these benchmark results:

Performance Analysis:
1. Time Complexity:
    - Size 10: ~2.7-2.9 ns/op
    - Size 100: ~71.6 ns/op
    - Size 1000: ~990 ns/op
    - Size 10000: ~10.5-11 μs/op
      Shows clear linear O(n) scaling

2. Memory Usage:
    - `0 B/op`: No bytes allocated per operation
    - `0 allocs/op`: No heap allocations per operation
      This is excellent - function only uses existing memory

3. Cyclic vs Non-cyclic comparison:
    - Small lists (10): ~5% overhead for cyclic
    - Medium lists (100): negligible difference
    - Large lists (1000): negligible difference
    - Very large lists (10000): ~4% overhead for cyclic

4. Iterations (reliability):
    - Small lists: ~400M iterations
    - Medium lists: ~16M iterations
    - Large lists: ~1.2M iterations
    - Very large lists: ~100K iterations
      Shows good statistical significance

Overall: This is an excellent implementation with:
- Linear time complexity
- Zero memory allocation
- Minimal overhead for cycle detection
- Consistent performance across sizes
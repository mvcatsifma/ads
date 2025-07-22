## 143. Reorder List

https://leetcode.com/problems/reorder-list/description/

## Benchmark results

Let's analyze these benchmark results in detail:

1. Time Complexity (ns/op):
    - Size 4: 110.3 ns
    - Size 10: 410.6 ns
    - Size 100: 5,638 ns
    - Size 1000: 47,006 ns
    - Size 10000: 828,209 ns
      Shows roughly O(n) scaling

2. Memory Usage:
    - Size 4: 64 B/op, 4 allocs/op
    - Size 10: 450 B/op, 11 allocs/op
    - Size 100: 6.9 KB/op, 107 allocs/op
    - Size 1000: 98.4 KB/op, 1011 allocs/op
    - Size 10000: 820 KB/op, 10099 allocs/op
      Shows linear memory growth due to map usage

3. Allocations Analysis:
    - Approximately n+3 allocations per operation
    - One allocation for map
    - One for head0
    - n allocations for map entries
    - Memory usage roughly 82 bytes per node

Potential Optimizations:
1. Consider using slice instead of map for storing nodes
2. Could potentially reduce allocations by reusing structures
3. Memory usage might be improved with array-based implementation

Overall Performance:
- Linear time complexity as expected
- Memory usage scales linearly with input size
- Performance is reasonable but memory allocations could be optimized
- Good performance on M3 Pro processor

## TODO

Space complexity is currently O(n) because of the map. Can and should be optimized to O(1).
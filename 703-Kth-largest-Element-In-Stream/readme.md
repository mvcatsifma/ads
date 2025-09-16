## Kth Largest Element in Stream

https://leetcode.com/problems/kth-largest-element-in-a-stream/

## Approach 1: Naive Sorting

This implementation uses a **naive sorting approach** for maintaining the kth largest element in a stream:

```go
func (k *KthLargest) Add(val int) int {
    k.nums = append(k.nums, val)
    slices.Sort(k.nums)
    return k.nums[len(k.nums)-k.k]
}
```

**Strategy:**
1. Append new value to slice
2. Sort entire slice on every Add operation
3. Return kth largest element from sorted position

## Performance Characteristics

### Benchmark Results (Apple M3 Pro)
- **Operations/sec**: ~11,600 Add operations per second
- **Time per operation**: 86.24 μs average
- **Memory per operation**: 40 bytes
- **Allocations**: 0 per operation (slice growth amortized)

### Time Complexity
- **Add operation**: O(n log n) due to full sort
- **Space complexity**: O(n) for storing all elements

### Performance Analysis

**Strengths:**
- Simple implementation
- Zero allocations per operation (efficient slice growth)
- Predictable memory usage

**Weaknesses:**
- **Inefficient for streams**: Sorting entire array on each Add
- **Degrades with size**: Performance drops as stream grows
- **Overkill**: Only need kth largest, not full sort

## Production Suitability

❌ **Not recommended for production streams**

**Better alternatives:**
- **Min-heap of size k**: O(log k) per Add
- **Quickselect**: O(n) average case
- **Balanced BST**: O(log n) per Add

**Use case fit**: Acceptable only for small, infrequent streams where simplicity outweighs performance.

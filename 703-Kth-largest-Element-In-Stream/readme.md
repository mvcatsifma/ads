# Kth Largest Element in Stream

https://leetcode.com/problems/kth-largest-element-in-a-stream/

# Approach 1: Naive Sorting

This implementation uses a **naive sorting approach** for maintaining the kth largest element in a stream.

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

# Approach 2: Min-Heap

## Algorithm Overview

This implementation uses a **min-heap of size k** to efficiently maintain the kth largest element in a stream. The key insight is that in a min-heap of the k largest elements, the root (minimum) is always the kth largest overall.

**Strategy:**
1. Maintain a min-heap containing only the k largest elements seen so far
2. When adding a new element, push it and remove excess elements to keep heap size ≤ k
3. The heap root is always the kth largest element

## Runtime Efficiency

### Time Complexity
- **Constructor**: O(n log n) where n = initial array size
    - O(n log n) to build initial heap
    - O((n-k) log k) to trim to size k
- **Add operation**: O(log k)
    - O(log k) to push new element
    - O(log k) to pop if heap exceeds size k

### Performance Characteristics
**Excellent scalability:**
- Add operations are **independent of total stream size**
- Only depends on k (typically small: 1-1000)
- **Logarithmic growth**: Even k=1000 requires only ~10 heap operations

**Comparison with naive sorting:**
- **This implementation**: O(log k) per Add
- **Previous sorting**: O(n log n) per Add (where n grows with each Add)
- **Improvement**: Massive speedup for large streams

## Memory Efficiency

### Space Complexity
- **Heap storage**: O(k) - only stores k largest elements
- **Auxiliary space**: O(1) - no additional data structures

### Memory Advantages
**Bounded memory usage:**
- Heap never exceeds k elements regardless of stream size
- **Memory-efficient for large streams**: 1M elements with k=3 uses same memory as 10 elements with k=3
- **Garbage collection friendly**: Discards smaller elements immediately

**Previous implementation comparison:**
- **This**: O(k) memory
- **Previous**: O(n) memory where n = total elements added

## Production Suitability

✅ **Highly recommended for production**

**Strengths:**
- **Optimal time complexity** for the problem
- **Bounded memory usage** prevents memory leaks
- **Scales to millions of elements** efficiently
- **Standard library integration** using `container/heap`

**Use cases:**
- **Real-time analytics**: Top-k tracking in high-throughput systems
- **Streaming data**: Processing infinite data streams
- **Memory-constrained environments**: IoT devices, embedded systems

## Implementation Quality

**Well-designed:**
- Clean separation of concerns (heap vs KthLargest logic)
- Proper use of Go's heap interface
- Defensive `Peek()` method for safe root access

This implementation represents **optimal algorithmic choice** for the kth largest element problem.
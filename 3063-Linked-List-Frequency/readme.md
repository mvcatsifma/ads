## 3063. Linked List Frequency

https://leetcode.com/problems/linked-list-frequency/description/

## Initial approach: Fixed-Size Array with Full Sort Frequency Counter

Here's a summary of the initial solution and its inefficiencies:

1. Creates a large fixed-size slice (1,000,000 elements)
2. Uses indices to count frequencies directly
3. Sorts entire slice including zeros
4. Builds result list from sorted frequencies

**Key Inefficiencies:**
```go
// 1. Excessive Memory Allocation
freqs := make([]int, 10e5)  // Always allocates ~1M elements regardless of input size

// 2. Inefficient Loop Structure
for {                       // Uses break instead of for curr != nil
    freqs[head.Val]++
    if head.Next == nil {
        break
    }
    head = head.Next
}

// 3. Unnecessary Full Sort
sort.Ints(freqs)           // Sorts entire 1M slice, including many zeros

// 4. Inefficient Result Building
for i := len(freqs) - 1; i >= 0; i-- {  // Iterates through many unnecessary zeros
    freq := freqs[i]
    if freq == 0 {
        break
    }
    // ...
}
```

**Performance Issues:**
- Space Complexity: O(1) but with large constant (1M elements)
- Time Complexity: O(n log n) due to sort
- Memory: ~8MB per operation (from benchmarks)
- Many unnecessary iterations over zero values

This solution needs optimization primarily in:
1. Memory allocation (reduce fixed size)
2. Sorting efficiency (avoid sorting zeros)
3. Loop structure (cleaner traversal)

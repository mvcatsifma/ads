# Chapter 2: Sorting - Study Guide

## For Working Engineers

Sorting is everywhere in software. This chapter shows you when to care and which algorithm to pick.

## The Reality Check
**Your language probably has a good built-in sort.** Use it unless you have a specific reason not to. This chapter teaches you:
- What happens under the hood
- When built-in sort isn't optimal
- How to recognize which algorithm is being used
- Interview preparation (yes, they still ask)

## Section 2.1: Elementary Sorts
**Skim with purpose**

### Selection Sort, Insertion Sort, Shell Sort

#### What You Need to Know
- **Selection Sort**: O(n²), doesn't adapt to data, simple but useless
- **Insertion Sort**: O(n²) worst case, but O(n) for nearly sorted data
- **Shell Sort**: O(n^(3/2)) with right gaps, interesting but niche

#### Skip the Theory, Learn This
**When to use Insertion Sort:**
- Arrays with < 10 elements (many libraries switch to insertion sort for small subarrays)
- Nearly sorted data (adding few items to sorted list)
- When you need a stable, in-place sort for tiny datasets

**Shell Sort:**
- Skip unless you're implementing a sorting library
- Historical interest only for practitioners

#### Real-World Example
```go
// Insertion sort is great here:
func insertIntoSortedList(sorted []int, newItem int) []int {
    sorted = append(sorted, newItem)
    // Insertion sort on last element: O(n) because array is nearly sorted
}
```

## Section 2.2: Mergesort
**Essential - Read carefully**

### Why It Matters
- O(n log n) guaranteed (no worst case surprises)
- Stable (preserves equal elements' order)
- Used in Java's sort for objects
- Parallelizes beautifully

### Key Concepts
- Divide and conquer: Split, sort halves, merge
- Requires extra space: O(n)
- Stable: Important for multi-key sorting

### Practical Applications
- External sorting (data doesn't fit in memory)
- When you NEED guaranteed O(n log n)
- Parallel/distributed sorting
- When stability matters (sort by name, then by age, preserve name order for same age)

### Implementation Tips
- Switch to insertion sort for small subarrays (< 10 elements)
- Bottom-up mergesort avoids recursion overhead
- Can optimize merge: if max(left) <= min(right), already sorted

### When to Use
- Stability required
- Worst-case performance matters
- Linked lists (mergesort is perfect for linked lists - no extra space needed!)
- Parallel processing

## Section 2.3: Quicksort
**Essential - This is the workhorse**

### Why It Matters
- Average O(n log n), usually fastest in practice
- In-place (no extra space)
- Used in most standard libraries
- C's qsort, Python's sort (with improvements)

### Key Concepts
- Partition: Pick pivot, put smaller left, larger right
- Recursively sort partitions
- Pivot choice matters (random pivot prevents worst case)

### The Catch
- Worst case O(n²) if you're unlucky with pivots
- Not stable
- Can have bad cache performance

### Practical Improvements
1. **Three-way partitioning**: Essential for duplicate-heavy data
2. **Median-of-three**: Pick pivot as median of first, middle, last
3. **Cutoff to insertion sort**: Switch to insertion sort for small subarrays
4. **Randomization**: Shuffle array first or pick random pivot

### When to Use
- General-purpose sorting (it's usually the default)
- In-place required
- No duplicate-heavy data (use 3-way if you have)
- Average-case performance over worst-case guarantee

### When to Avoid
- Stability required → use mergesort
- Guaranteed O(n log n) required → use mergesort
- Lots of duplicates without 3-way partitioning → slow

## Section 2.4: Priority Queues
**Essential - Extremely practical**

### Why This Is Important
Priority queues are everywhere in real systems:
- Job scheduling (process highest priority first)
- Event-driven simulation
- Dijkstra's shortest path
- Huffman coding
- OS task schedulers

### Binary Heap Implementation
- Array-based binary tree
- Insert: O(log n)
- Remove max/min: O(log n)
- Peek: O(1)

### Practical Applications
1. **Task Schedulers**: Process high-priority tasks first
2. **Real-time Systems**: Event queue sorted by timestamp
3. **Graph Algorithms**: Dijkstra, Prim's MST
4. **Top-K Problems**: Keep k largest elements efficiently
5. **Median Finding**: Two heaps (max-heap for small half, min-heap for large half)

### Implementation Tips
- Array-based is simpler and faster than tree-based
- Parent at index k, children at 2k and 2k+1 (1-indexed)
- Or: Parent at k, children at 2k+1 and 2k+2 (0-indexed)
- Min-heap vs. max-heap: just flip the comparison

### Heapsort
- O(n log n) in-place, not stable
- Good worst-case, but slower than quicksort in practice
- Interesting but rarely the best choice

## Section 2.5: Applications
**Read for context**

### Key Takeaways
- **Stability matters** for multi-key sorting
- **Duplicate keys**: 3-way quicksort is huge win
- **Various types**: Just provide comparison function
- **Indirect sorts**: Sort indices, not data

## Practical Decision Tree

### Small Data (n < 50)
Use insertion sort or your language's built-in

### General Purpose (n > 50)
**Use Quicksort** (with improvements):
- 3-way partitioning
- Median-of-three pivot
- Cutoff to insertion sort

### Guaranteed Performance
**Use Mergesort**:
- External sorting
- Linked lists
- Parallel sorting

### Stability Required
**Use Mergesort** or **Timsort** (Python's hybrid)

### Priority Queue
**Use Binary Heap**:
- Not sorting, just need max/min efficiently
- Scheduling, event processing

### Already Have Built-in Sort
**Use it!** Modern standard libraries use:
- Java: Timsort (objects), Dual-pivot quicksort (primitives)
- Python: Timsort
- Go: Pattern-defeating quicksort
- C++: Introsort (quicksort with heapsort fallback)

## What to Skip
- Detailed mathematical proofs of complexity
- Every single sorting variant
- Historical algorithms (unless intellectually curious)
- Formal stability proofs

## What to Focus On
- Quicksort with improvements (this is what you'll implement/use)
- Mergesort for stability
- Priority queue/heap operations
- When to use which algorithm
- Understanding trade-offs

## Study Time Estimate
- Section 2.1: 1 hour (skim, understand insertion sort's niche)
- Section 2.2: 2-3 hours (understand mergesort, implement it)
- Section 2.3: 3-4 hours (understand quicksort, implement 3-way)
- Section 2.4: 3-4 hours (heap operations, implement priority queue)
- Section 2.5: 1 hour (applications, stability)

**Total: 10-13 hours for practical mastery**

## Exercises Worth Doing
1. Implement quicksort with 3-way partitioning
2. Implement binary heap priority queue
3. Compare performance: quicksort vs. mergesort vs. built-in
4. Solve LeetCode problems using priority queue

## Next Steps
Chapter 3 will show you how to search efficiently, building on sorting concepts.

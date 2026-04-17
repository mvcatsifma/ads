# Sorting Algorithm Performance Characteristics

Reference: Sedgewick "Algorithms", 4th Edition, Page 342

## Performance Comparison

| Algorithm          | Stable | In-place | Running Time         | Extra Space | Notes                                  |
|--------------------|--------|----------|----------------------|-------------|----------------------------------------|
| Selection Sort     | No     | Yes      | N²                   | 1           | N exchanges                            |
| Insertion Sort     | Yes    | Yes      | Between N and N²     | 1           | Depends on order of items              |
| Shell Sort         | No     | Yes      | N lg N               | 1           | Tight code, subquadratic               |
| Merge Sort         | Yes    | No       | N lg N               | N           | Guarantee                              |
| Merge Sort (BU)    | Yes    | No       | N lg N               | N           | Guarantee, no recursion                |
| Quick Sort         | No     | Yes      | N lg N               | lg N        | Probabilistic guarantee                |
| Quick Sort (3-way) | No     | Yes      | Between N and N lg N | lg N        | Improves quicksort when duplicate keys |
| Heap Sort          | No     | Yes      | N lg N               | 1           | Guarantee, in-place                    |

**Legend:**
- N = number of items
- lg N = log₂(N)
- Running time is number of compares (typical case or guarantee as noted)
- Extra space is proportional to given value (not including input array)
- In-place means uses ≤ c lg N extra memory

## Key Observations

### Time Complexity
- **Elementary sorts** (Selection, Insertion): O(N²) - suitable for small arrays
- **Shell sort**: Subquadratic but exact complexity depends on gap sequence
- **Efficient sorts** (Merge, Quick, Heap): O(N lg N) average case
- **Quick sort**: Fastest general-purpose sort in practice
  - Property: In-place partitioning with excellent cache locality
  - Evidence: 39% fewer compares than merge sort on average, minimal data movement
  - Can degrade to O(N²) without randomization/shuffle
- **3-way Quick sort**: O(N) for arrays with many duplicates

### Space Complexity
- **In-place sorts**: Selection, Insertion, Shell, Heap - O(1) extra space
- **Merge sort**: O(N) extra space for auxiliary array
- **Quick sort**: O(lg N) for recursion stack (average case)

### Stability
- **Stable**: Insertion Sort, Merge Sort (both variants)
- **Unstable**: Selection Sort, Shell Sort, Quick Sort (both variants), Heap Sort

## Algorithm Selection Guide

**Use Insertion Sort when:**
- Array size < 10-15 elements
- Array is nearly sorted
- Stability is required and array is small

**Use Merge Sort when:**
- Stability is required
- Guaranteed worst-case performance needed
- Extra space is available

**Use Quick Sort when:**
- Average case performance is priority (fastest general-purpose sort in practice)
- Space is limited (no auxiliary array)
- Stability not required
- Input is random or can be shuffled
- Note: If stability is important and space is available, merge sort might be the best choice

**Use 3-way Quick Sort when:**
- Many duplicate keys expected
- Same benefits as Quick Sort

**Use Heap Sort when:**
- Guaranteed N lg N needed
- Space is extremely limited (in-place)
- Stability not required

**Use Shell Sort when:**
- Simple implementation desired
- Small to medium-sized arrays
- Reasonable performance acceptable

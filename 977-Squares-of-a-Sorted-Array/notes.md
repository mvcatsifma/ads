# Notes

## Time Complexity
- Brute force: Square all elements, sort → O(n log n)
- Optimized: Two-pointer approach → O(n)

## Space Complexity
- Brute force: O(n) - for result array
- Optimized: O(n) - for result array only (O(1) if not counting output)

## Edge Cases
- [ ] Single element array
- [ ] All negative numbers
- [ ] All positive numbers
- [ ] Array containing zero
- [ ] All zeros
- [ ] Mixed negative and positive (where largest squares come from)

## Brute Force Baseline
Square each element, then sort the result array.

- Algorithm:
  1. Create result array
  2. Square each element
  3. Sort result array
- Time: O(n log n) - dominated by sorting
- Space: O(n) - for result array (or O(1) if not counting output)
- Why it works: Straightforward - squares then sorts
- Limitations: Not optimal - doesn't use the fact that input is already sorted

**Benchmark Results (Apple M3 Pro):**
```
BenchmarkSortedSquares-12    	17430738	        69.13 ns/op	     120 B/op	       4 allocs/op
```
- 69.13 ns per operation
- 120 bytes allocated per operation
- 4 allocations per operation

## Optimizations
**Key insight:** Input array is sorted, so negative numbers are on left, positive on right.
After squaring, the largest values come from the edges (most negative or most positive).

Think: Two-pointer approach
- Start from both ends
- Compare absolute values (or squared values)
- Build result array from back to front (largest to smallest)
- Time: O(n) - single pass
- Space: O(n) - for result array only

**Benchmark Results (Apple M3 Pro):**
```
BenchmarkSortedSquares-12    	72757412	        20.86 ns/op	      48 B/op	       1 allocs/op
```
- 20.86 ns per operation (3.3x faster than brute force)
- 48 bytes allocated per operation (60% less memory)
- 1 allocation per operation (75% fewer allocations)

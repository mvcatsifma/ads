# Notes

## Time Complexity
- TODO: Analyze time complexity
- Brute force: Square all elements, sort → O(n log n)
- Optimized: TODO

## Space Complexity
- TODO: Analyze space complexity
- Brute force: O(1) or O(n) depending on if we count output
- Optimized: TODO

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

## Optimizations
**Key insight:** Input array is sorted, so negative numbers are on left, positive on right.
After squaring, the largest values come from the edges (most negative or most positive).

Think: Two-pointer approach
- Start from both ends
- Compare absolute values (or squared values)
- Build result array from back to front (largest to smallest)
- Time: O(n) - single pass
- Space: O(n) - for result array only

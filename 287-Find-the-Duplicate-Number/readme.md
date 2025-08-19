## 287. Find the Duplicate Number

https://leetcode.com/problems/find-the-duplicate-number/

**Approach: Find Duplicate Number Using Binary Search**

## Description
This function finds the single duplicate number in an array containing `n+1` integers where each integer is between 1 and `n` (inclusive). It uses a binary search approach based on the pigeonhole principle.

**Algorithm Overview:**
- Performs binary search on the range [1, n] rather than on array indices
- For each candidate value `mid`, counts how many numbers in the array are ≤ `mid`
- If the count exceeds `mid`, the duplicate must be in the range [1, mid]
- Otherwise, the duplicate is in the range [mid+1, n]
- Continues until the search space narrows to a single number

**Time Complexity:** O(n log n) - binary search iterations × linear counting
**Space Complexity:** O(1) - only uses constant extra space

**Key Insight:** In an array with no duplicates, exactly `k` numbers would be ≤ `k`. When there's a duplicate, this count increases for all values ≥ the duplicate number.

The helper function `countLessOrEqual` efficiently counts elements in the array that are less than or equal to the given target value.
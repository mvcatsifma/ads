## 88. Merge Sorted Array

https://leetcode.com/problems/merge-sorted-array/description/

## Approach 1 - Merge Sorted Arrays (Naive Implementation)

Merges two sorted integer arrays into a single sorted array in-place. This implementation uses a simple append-and-sort approach rather than taking advantage of the pre-sorted nature of the input arrays.

**Algorithm:** Appends all elements from the second array to the first array, then applies bubble sort to the combined result.

**Time Complexity:** O((m+n)²) due to bubble sort  
**Space Complexity:** O(1) - modifies the first array in-place

**Use Case:** Suitable for small datasets or when simplicity is prioritized over performance. For larger datasets, consider the optimal O(m+n) merge approach that leverages the sorted property of both input arrays.

**Example:**
```go
nums1 := []int{1, 2, 3, 0, 0, 0} // has capacity for 6 elements
nums2 := []int{2, 5, 6}
merge(nums1, 3, nums2, 3)
// Result: nums1 = [1, 2, 2, 3, 5, 6]
```
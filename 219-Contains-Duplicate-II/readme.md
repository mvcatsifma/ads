## 219. Contains Duplicate II

https://leetcode.com/problems/contains-duplicate-ii/description/


## Approach 2 - sliding window with hash map optimization

## Description
Efficiently determine if an array contains duplicate values within a specified distance constraint using a sliding window approach with hash map optimization. This function checks whether any two identical elements exist within `k` positions of each other.

**Problem Statement:**
Given an integer array `nums` and an integer `k`, return `true` if there are two distinct indices `i` and `j` in the array such that `nums[i] == nums[j]` and the absolute difference between `i` and `j` is at most `k`.

**Examples:**
```
Input: nums = [1,2,3,1], k = 3
Output: true
Explanation: nums[0] and nums[3] are both 1, within distance k=3

Input: nums = [1,0,1,1], k = 1  
Output: true
Explanation: nums[2] and nums[3] are both 1, within distance k=1

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
Explanation: No duplicates exist within distance k=2
```

**Algorithm Approach:**
Uses a **sliding window technique** with a hash map to maintain a window of at most `k` elements. The map tracks elements currently within the valid distance range. As the window slides forward, elements outside the k-distance are removed, ensuring the map only contains elements that could form valid duplicate pairs.

**Key Optimizations:**
- Hash map provides O(1) duplicate detection
- Sliding window maintains exactly k+1 elements maximum
- Early termination when duplicate found
- Automatic cleanup of out-of-range elements

**Time Complexity:** O(n) - each element processed once with constant-time map operations  
**Space Complexity:** O(min(n,k)) - map size limited by window size or array length

**Use Cases:** Optimized duplicate detection, streaming data validation, memory-efficient proximity checking, and high-performance array analysis with distance constraints.

## Approach 1 -  Use a brute force nested loop

## Description
Determine if an array contains duplicate values within a specified distance constraint. This function checks whether any two identical elements exist within `k` positions of each other in the array.

**Problem Statement:**
Given an integer array `nums` and an integer `k`, return `true` if there are two distinct indices `i` and `j` in the array such that `nums[i] == nums[j]` and the absolute difference between `i` and `j` is at most `k`.

**Examples:**
```
Input: nums = [1,2,3,1], k = 3
Output: true
Explanation: nums[0] and nums[3] are both 1, and |0-3| = 3 ≤ k

Input: nums = [1,0,1,1], k = 1  
Output: true
Explanation: nums[2] and nums[3] are both 1, and |2-3| = 1 ≤ k

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
Explanation: No duplicates exist within distance k=2
```

**Algorithm Approach:**
Uses a **brute force nested loop** approach where for each element, it checks all elements within the next `k` positions for duplicates. The outer loop iterates through each element, while the inner loop examines the constrained window of potential duplicate locations.

**Key Characteristics:**
- Examines each element against its k-distance neighbors
- Early termination when duplicate found within range
- Bounds checking prevents array index overflow
- Simple and straightforward implementation

**Time Complexity:** O(n×k) - for each element, check up to k subsequent elements  
**Space Complexity:** O(1) - constant extra space usage

**Use Cases:** Data validation, proximity-based duplicate detection, sliding window constraint problems, and array analysis with distance limitations.
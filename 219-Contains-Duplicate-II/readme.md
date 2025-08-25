## 219. Contains Duplicate II

https://leetcode.com/problems/contains-duplicate-ii/description/

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
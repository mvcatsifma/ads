## 594. Longest Harmonious Subsequence

https://leetcode.com/problems/longest-harmonious-subsequence/description/

## Approach - Use a frequency counting strategy with hash map optimization

## Description
Find the length of the longest harmonious subsequence in an array, where a harmonious subsequence is defined as a subsequence in which the difference between the maximum and minimum values is exactly 1.

**Problem Statement:**
Given an integer array `nums`, return the length of its longest harmonious subsequence among all its possible subsequences. A subsequence maintains the relative order of elements but doesn't need to be contiguous.

**Harmonious Subsequence Rules:**
- Must contain at least two distinct values
- The difference between max and min values must be exactly 1
- Elements maintain their original relative order from the array

**Examples:**
```
Input: [1,3,2,2,5,2,3,7]
Output: 5
Explanation: Longest harmonious subsequence is [3,2,2,2,3] (contains 2 and 3)

Input: [1,2,3,4]  
Output: 2
Explanation: Multiple valid subsequences: [1,2], [2,3], or [3,4]

Input: [1,1,1,1]
Output: 0
Explanation: No harmonious subsequence possible (only one distinct value)
```

**Algorithm Approach:**
Uses a **frequency counting strategy** with hash map optimization. First pass counts occurrences of each number, then second pass examines each number to see if its consecutive neighbor (num+1) exists. The sum of frequencies for valid consecutive pairs represents potential harmonious subsequence lengths.

**Key Insights:**
- Only consecutive integer pairs can form harmonious subsequences
- Frequency counting eliminates need to track actual subsequence elements
- Checking only num+1 (not num-1) avoids counting pairs twice
- Maximum frequency sum across all valid pairs gives the answer

**Time Complexity:** O(n) - two linear passes through the data  
**Space Complexity:** O(n) - hash map stores unique elements and their frequencies

**Applications:** Data analysis, pattern recognition, sequence optimization, and problems involving constrained value ranges in ordered data.
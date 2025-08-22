## 3. Longest Substring Without Repeating Characters

https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

##Longest Substring Without Repeating Characters

## Description
Find the length of the longest contiguous substring within a given string that contains no duplicate characters. This classic string processing problem requires efficiently tracking character occurrences while maintaining optimal performance.

**Problem Statement:**
Given a string, return the length of the longest substring where all characters are unique.

**Examples:**
```
Input: "abcabcbb"
Output: 3
Explanation: "abc" is the longest substring without repeating characters

Input: "bbbbb" 
Output: 1
Explanation: "b" is the longest (all characters are the same)

Input: "pwwkew"
Output: 3  
Explanation: "wke" is the longest substring without repeating characters
```

**Algorithm Approach:**
Uses the **sliding window technique** with a hash map to efficiently track character occurrences. Two pointers maintain a dynamic window that expands when unique characters are found and contracts when duplicates are encountered.

**Key Insights:**
- Maintains a window of unique characters using left and right pointers
- Hash map provides O(1) character lookup and removal
- Each character is visited at most twice (once by each pointer)
- Avoids nested loops through intelligent window management

**Time Complexity:** O(n) - linear scan with each character processed at most twice  
**Space Complexity:** O(min(m,n)) - hash map size limited by charset size or string length

**Applications:** Text processing, data deduplication, pattern matching, and optimization problems requiring unique element tracking.
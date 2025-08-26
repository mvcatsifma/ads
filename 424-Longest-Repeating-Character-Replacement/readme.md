## 424. Longest Repeating Character Replacement

https://leetcode.com/problems/longest-repeating-character-replacement/description/

## Longest Repeating Character Replacement

Finds the length of the longest substring containing the same letter you can get after performing at most k character replacements.

**Algorithm:** Uses a sliding window approach with character frequency tracking. The window expands when valid (replacements ≤ k) and contracts when invalid, maintaining the maximum valid window size seen.

**Time Complexity:** O(n) where n is the string length  
**Space Complexity:** O(1) for the character frequency map (at most 26 unique characters)

**Example:**
```go
characterReplacement("AABABBA", 1) // returns 4
// Can replace one 'B' to get "AAAA" (substring "AABA" → "AAAA")
```
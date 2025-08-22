package p3

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters
// using a sliding window approach with a hash map to track character occurrences.
//
// Algorithm:
// 1. Use two pointers (l, r) to maintain a sliding window
// 2. Expand window by moving right pointer and adding characters to map
// 3. When duplicate found, shrink window from left until duplicate is removed
// 4. Track maximum window size encountered
//
// Time Complexity: O(n) - each character visited at most twice
// Space Complexity: O(min(m,n)) where m is charset size, n is string length
//
// Example:
//   lengthOfLongestSubstring("abcabcbb") returns 3 (substring "abc")
//   lengthOfLongestSubstring("bbbbb") returns 1 (substring "b")
//   lengthOfLongestSubstring("pwwkew") returns 3 (substring "wke")
//
// Parameters:
//   s: input string to analyze
//
// Returns: length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	answer := 0
	l := 0
	chars := make(map[uint8]bool)

	for r := 0; r < len(s); r++ {
		// Shrink window while duplicate exists
		for chars[s[r]] {
			delete(chars, s[l])
			l++
		}

		chars[s[r]] = true

		// Update max length
		if r-l+1 > answer {
			answer = r - l + 1
		}
	}

	return answer
}

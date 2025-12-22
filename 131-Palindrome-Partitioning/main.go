package p131

// partition returns all possible palindrome partitions of the input string.
// A palindrome partition divides the string into substrings where each substring is a palindrome.
// Example: partition("aab") returns [["a","a","b"], ["aa","b"]]
//
// Approach: Uses backtracking to explore all possible ways to split the string.
// For each position, tries all possible palindromic prefixes and recursively partitions the remainder.
// Time Complexity: O(n × 2^n) where n is string length (2^n partitions, O(n) to check palindrome).
// Space Complexity: O(n) for recursion stack + O(n × 2^n) for storing all partitions.
func partition(s string) [][]string {
	var res [][]string
	dfs(s, []string{}, &res)
	return res
}

// dfs performs depth-first search to find all palindrome partitions of the remaining string.
// Recursively tries each possible palindromic prefix and continues with the suffix.
//
// Parameters:
//   - s: remaining string to partition
//   - path: current partition being built
//   - res: pointer to result slice to collect all valid partitions
func dfs(s string, path []string, res *[][]string) {
	// Base case: no more characters to partition
	if len(s) == 0 {
		// Found a complete partition - copy path to avoid reference issues
		*res = append(*res, append([]string{}, path...))
		return
	}

	// Try all possible prefixes of remaining string
	for i := 1; i <= len(s); i++ {
		prefix := s[:i]
		// Only continue if current prefix is a palindrome
		if isPalindrome(prefix) {
			// Recursively partition the suffix with current prefix added to path
			dfs(s[i:], append(path, prefix), res)
		}
	}
}

// isPalindrome checks if a string reads the same forwards and backwards.
// Uses two pointers approach for O(n) time complexity.
//
// Returns true if the string is a palindrome, false otherwise.
func isPalindrome(s string) bool {
	lo, hi := 0, len(s)-1

	// Compare characters from both ends moving inward
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}

	return true
}

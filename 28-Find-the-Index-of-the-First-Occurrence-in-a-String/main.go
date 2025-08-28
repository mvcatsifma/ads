package p28

// strStr finds the first occurrence of a substring (needle) within a string (haystack).
//
// This function implements a naive string matching algorithm that performs character-by-character
// comparison. It uses O(1) auxiliary space complexity, making it memory efficient for large inputs.
//
// Parameters:
//   - haystack: the string to search within
//   - needle: the substring to search for
//
// Returns:
//   - The zero-based index of the first occurrence of needle in haystack
//   - 0 if needle is empty (following Go convention that empty string is found at position 0)
//   - -1 if needle is not found or if needle is longer than haystack
//
// Time Complexity: O(n*m) where n = len(haystack), m = len(needle)
// Space Complexity: O(1) - no additional memory allocation
//
// Example:
//   strStr("hello world", "wor") returns 6
//   strStr("aaaa", "aaa") returns 0
//   strStr("hello", "") returns 0
//   strStr("hello", "xyz") returns -1
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			// Check remaining characters
			j := 1
			for j < len(needle) && haystack[i+j] == needle[j] {
				j++
			}
			if j == len(needle) {
				return i
			}
		}
	}

	return -1
}

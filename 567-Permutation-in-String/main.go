package p567

// checkInclusion checks if s1 is a permutation of a substring of s2.
// Returns true if any contiguous substring of s2 has the same character frequencies as s1.
// Example: checkInclusion("ab", "eidbaooo") returns true because "ba" is a permutation of "ab".
func checkInclusion(s1 string, s2 string) bool {
	// Early exit: s1 can't be a permutation of any substring if it's longer than s2
	if len(s1) > len(s2) {
		return false
	}

	// Create frequency array for s1 (26 lowercase letters)
	s1Arr := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1Arr[s1[i]-'a']++ // Map character to index (a=0, b=1, ..., z=25)
	}

	// Sliding window: check every substring of s2 with length = len(s1)
	for i := 0; i <= len(s2)-len(s1); i++ {
		// Build frequency array for current substring of s2
		s2Arr := make([]int, 26)
		for j := 0; j < len(s1); j++ {
			s2Arr[s2[i+j]-'a']++
		}

		// Check if character frequencies match
		if matches(s1Arr, s2Arr) {
			return true
		}
	}

	return false
}

// matches compares two character frequency arrays.
// Returns true if all frequencies are identical, false otherwise.
func matches(arr1 []int, arr2 []int) bool {
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

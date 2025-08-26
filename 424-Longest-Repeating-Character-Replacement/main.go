package p424

// characterReplacement returns the length of the longest substring that can be made
// to contain only one unique character by replacing at most k characters.
// Uses a sliding window approach with character frequency tracking.
//
// Parameters:
//   - str: input string to analyze
//   - k: maximum number of character replacements allowed
//
// Returns the maximum length of valid substring.
//
// Example: characterReplacement("AABABBA", 1) returns 4 (substring "AABA")
func characterReplacement(str string, k int) int {
	var maxCount int
	var result int
	var left = 0

	counts := make(map[string]int)

	for right, strR := range str {
		counts[string(strR)]++
		maxCount = findMaxValue(counts)
		for (right-left+1)-maxCount > k {
			strL := string(str[left])
			counts[strL]--
			left++
		}

		if windowSize := right - left + 1; windowSize > result {
			result = windowSize
		}
	}

	return result
}

func findMaxValue(m map[string]int) int {
	var maxVal int
	for _, v := range m {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

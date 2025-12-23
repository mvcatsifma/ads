package p17

import "strings"

// buttons maps phone keypad digits to their corresponding letters.
// Represents the traditional T9 phone keypad layout where each digit (2-9) maps to 3-4 letters.
var buttons = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

// letterCombinations returns all possible letter combinations that a phone number could represent.
// Each digit (2-9) on a phone keypad corresponds to 3-4 letters, similar to old T9 texting.
// Example: letterCombinations("23") returns ["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// Approach: Uses backtracking to generate all possible combinations.
// For each digit position, tries all possible letters and recursively builds combinations.
// Time Complexity: O(3^n × 4^m) where n is digits with 3 letters, m is digits with 4 letters.
// Space Complexity: O(3^n × 4^m) for storing all combinations + O(n) for recursion stack.
func letterCombinations(digits string) []string {
	// Edge case: empty input returns empty result
	if len(digits) == 0 {
		return []string{}
	}

	var result []string

	// Recursive backtracking function
	var backtrack func(i int, curr []string)
	backtrack = func(i int, curr []string) {
		// Base case: built a complete combination
		if len(curr) == len(digits) {
			// Join letters into a single string and add to results
			result = append(result, strings.Join(curr, ""))
			return
		}

		// Get current digit and its corresponding letters
		digit := digits[i]
		button := buttons[string(digit)]

		// Try each possible letter for current digit
		for _, letter := range button {
			// Choose: add current letter to combination
			curr = append(curr, letter)

			// Explore: recursively process next digit
			backtrack(i+1, curr)

			// Unchoose: remove current letter (backtrack)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []string{})
	return result
}

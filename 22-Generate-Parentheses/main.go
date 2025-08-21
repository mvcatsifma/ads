package p22

// generateParenthesis generates all combinations of well-formed parentheses for n pairs.
// It uses backtracking to build valid parentheses strings by ensuring that at any point:
// - The number of opening parentheses doesn't exceed n
// - The number of closing parentheses doesn't exceed the number of opening parentheses
//
// Time Complexity: O(4^n / √n) - follows the nth Catalan number
// Space Complexity: O(4^n / √n) for output + O(n) for recursion stack
//
// Example:
//   generateParenthesis(2) returns ["(())", "()()"]
//   generateParenthesis(3) returns ["((()))", "(()())", "(())()", "()(())", "()()()"]
//
// Parameters:
//   n: number of parentheses pairs (must be non-negative)
//
// Returns:
//   []string: slice of all valid parentheses combinations
func generateParenthesis(n int) []string {
	var answer []string
	var backtracking func(s string, leftCount, rightCount int)
	backtracking = func(s string, leftCount int, rightCount int) {
		if len(s) == 2*n {
			answer = append(answer, s)
			return
		}

		if leftCount < n {
			backtracking(s+"(", leftCount+1, rightCount)
		}
		if rightCount < leftCount {
			backtracking(s+")", leftCount, rightCount+1)
		}

	}
	backtracking("", 0, 0)
	return answer
}

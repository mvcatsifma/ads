package p136

// singleNumber finds the single element that appears exactly once in an array
// where every other element appears exactly twice, using the XOR bit manipulation trick.
//
// Algorithm: XOR has the mathematical properties:
// - a ^ a = 0 (any number XOR itself equals zero)
// - a ^ 0 = a (any number XOR zero equals itself)
// - XOR is commutative and associative
//
// When we XOR all numbers together, the pairs cancel out (become 0),
// leaving only the single number that appears once.
//
// Example: [4, 1, 2, 1, 2] → 4 ^ 1 ^ 2 ^ 1 ^ 2 = 4 ^ (1^1) ^ (2^2) = 4 ^ 0 ^ 0 = 4
//
// Time complexity: O(n) - single pass through array
// Space complexity: O(1) - only uses one integer variable
//
// Constraints: Assumes exactly one number appears once, all others appear twice
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // XOR each number
	}
	return result
}

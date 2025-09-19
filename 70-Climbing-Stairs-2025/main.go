package p70

// climbStairs calculates the number of distinct ways to climb n stairs
// where you can climb either 1 or 2 steps at a time.
// This implements the Fibonacci sequence using O(1) space complexity
// by maintaining only the last two values instead of storing all previous results.
// Time complexity: O(n), Space complexity: O(1)
func climbStairs(n int) int {
	one, two := 1, 1
	for i := 0; i < n-1; i++ {
		temp := one
		one = one + two
		two = temp
	}

	return one
}

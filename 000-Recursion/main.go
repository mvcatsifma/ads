package recursion

// fibonacciDynamicTopdown returns the n-th Fibonacci number using
// top-down dynamic programming (memoized recursion). The caller must
// provide a memo slice of length at least n+1. Uncomputed entries in
// memo should be zero-initialized; memo[n] will be filled on demand.
// Time complexity: O(n). Space complexity: O(n) due to recursion depth.
func fibonacciDynamicTopdown(n int, memo []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] == 0 {
		memo[n] = fibonacciDynamicTopdown(n-1, memo) + fibonacciDynamicTopdown(n-2, memo)
	}

	return memo[n]
}

// fibonacciDynamicBottomUp returns the n-th Fibonacci number using a
// bottom-up dynamic programming approach. This implementation is
// entirely non-recursive and uses O(1) space by keeping only the last
// two computed values. Runs in O(n) time.
func fibonacciDynamicBottomUp(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1

	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return a + b // b now holds F(n)
}

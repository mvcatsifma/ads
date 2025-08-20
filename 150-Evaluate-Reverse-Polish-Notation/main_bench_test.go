package p50

import "testing"

func BenchmarkEvalRPN(b *testing.B) {
	// Test case with mixed operations - realistic RPN expression
	tokens := []string{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"}
	// This evaluates to: ((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)) = ((15 / 5) * 3) - 4 = 9 - 4 = 5

	b.ReportAllocs() // Enable memory allocation reporting
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := evalRPN(tokens)
		_ = result // Prevent compiler optimization
	}
}

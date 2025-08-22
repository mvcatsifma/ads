package p3

import "testing"

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	// Test with realistic mixed case: repeating and non-repeating patterns
	testString := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()"

	b.ReportAllocs() // Enable allocation reporting
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(testString)
	}
}

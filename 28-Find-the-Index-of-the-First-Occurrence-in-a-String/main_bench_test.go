package p28

import "testing"

func BenchmarkStrStr(b *testing.B) {
	haystack := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	needle := "xyz"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strStr(haystack, needle)
	}
}

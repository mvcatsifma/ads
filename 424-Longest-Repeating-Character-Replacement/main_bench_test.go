package p424

import "testing"

func BenchmarkCharacterReplacement(b *testing.B) {
	str := "AABABBAABCDEFGHIJKLMNOPQRSTUVWXYZAABABBA"
	k := 2

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		characterReplacement(str, k)
	}
}

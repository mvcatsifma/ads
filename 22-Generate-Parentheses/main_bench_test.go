package p22

import (
	"fmt"
	"testing"
)

func BenchmarkGenerateParenthesis(b *testing.B) {
	testCases := []int{1, 2, 3, 4}

	for _, n := range testCases {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				result := generateParenthesis(n)
				_ = result
			}
		})
	}
}

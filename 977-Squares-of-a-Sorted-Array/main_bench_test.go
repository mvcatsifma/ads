package p977

import "testing"

func BenchmarkSortedSquares(b *testing.B) {
	input := []int{-7, -3, 2, 3, 11}
	for i := 0; i < b.N; i++ {
		sortedSquares(input)
	}
}

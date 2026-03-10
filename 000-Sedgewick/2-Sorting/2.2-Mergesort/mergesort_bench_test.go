package mergesort

import "testing"

func BenchmarkMergeSort(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		MergeSort(arr)
	}
}

func BenchmarkMergeSortBU(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		MergeSortBU(arr)
	}
}

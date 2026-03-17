package quicksort

import "testing"

func BenchmarkQuickSort(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		QuickSort(arr)
	}
}

func BenchmarkQuickSort3Way(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		QuickSort3Way(arr)
	}
}

func BenchmarkQuickSort3Way_ManyDuplicates(b *testing.B) {
	input := []int{5, 2, 5, 1, 5, 2, 5, 1, 5, 2, 5, 1, 5}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		QuickSort3Way(arr)
	}
}

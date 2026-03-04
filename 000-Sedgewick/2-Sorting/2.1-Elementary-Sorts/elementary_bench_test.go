package elementary

import "testing"

func BenchmarkSelectionSort(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		SelectionSort(arr)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		InsertionSort(arr)
	}
}

func BenchmarkShellSort(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(input))
		copy(arr, input)
		ShellSort(arr)
	}
}

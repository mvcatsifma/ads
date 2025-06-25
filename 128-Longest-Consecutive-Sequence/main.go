package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	sort.Ints(nums)
	fmt.Printf("%d\n", nums)
}

// TODO: use map[int]bool in order to de-duplicate the input.
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	seqs := make([][]int, 0)
	seq := make([]int, 0)
	prev := 0
	for _, num := range nums {
		if num-prev == 1 {
			seq = append(seq, num)
		} else {
			seqs = append(seqs, seq)
			seq = make([]int, 0)
			seq = append(seq, num)
		}
		prev = num
	}
	seqs = append(seqs, seq)

	// TODO: time complexity of this loop is O(N),
	// replace with sort which has an average time complexity of O(n log n).
	var longest []int
	for _, seq := range seqs {
		if len(seq) > len(longest) {
			longest = seq
		}
	}

	return len(longest)
}

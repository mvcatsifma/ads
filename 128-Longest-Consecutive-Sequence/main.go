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

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	seqs := make([][]int, 0)
	seq := make([]int, 0)
	prev := nums[:1][0]
	seq = append(seq, prev)
	for _, num := range nums[1:] {
		if num-prev == 0 { // skip duplicate
			continue
		}
		if num-prev == 1 { // add to sequence
			seq = append(seq, num)
		} else { // start new sequence
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

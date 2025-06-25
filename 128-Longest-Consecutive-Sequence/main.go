package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 { // return immediately
		return 0
	}

	// Sort the input array
	sort.Ints(nums)

	seqs := make([][]int, 0) // list of sequences
	seq := make([]int, 0)    // the current sequence
	var prev int
	for i, num := range nums {
		if i == 0 { // first element in the input array is always added to the first sequence
			seq = append(seq, num)
			prev = num
			continue
		}
		diff := num - prev
		switch diff {
		case 0: // duplicate element
			continue
		case 1: // consecutive element
			seq = append(seq, num)
		default: // non-consecutive element, start a new sequence
			seqs = append(seqs, seq)
			seq = make([]int, 0)
			seq = append(seq, num)
		}
		prev = num
	}
	seqs = append(seqs, seq) // add the last sequence to the sequences array

	// Sort the accumulated sequences.
	sort.Slice(seqs, func(i, j int) bool {
		return len(seqs[i]) > len(seqs[j])
	})

	return len(seqs[0]) // return the length of the longest sequence
}

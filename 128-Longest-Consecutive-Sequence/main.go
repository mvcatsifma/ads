package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 { // return immediately
		return 0
	}

	seqs := make([][]int, 0) // list of sequences
	seq := make([]int, 0)    // the current sequence

	var result int
	var appendSeq = func(seq []int) {
		if len(seq) > result {
			result = len(seq)
		}
		seqs = append(seqs, seq)
	}

	// Sort the input array
	sort.Ints(nums)

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
			appendSeq(seq)
			seq = make([]int, 0)
			seq = append(seq, num)
		}
		prev = num
	}
	appendSeq(seq) // add the last sequence to the sequences array

	return result // return the length of the result sequence
}

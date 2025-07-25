package main

import "sort"

func frequenciesOfElements(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	freqs := make([]int, 10e5)
	for {
		freqs[head.Val]++
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	sort.Ints(freqs)

	dummy := &ListNode{}
	curr := dummy
	for i := len(freqs) - 1; i >= 0; i-- {
		freq := freqs[i]
		if freq == 0 {
			break
		}
		curr.Next = &ListNode{Val: freq}
		curr = curr.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

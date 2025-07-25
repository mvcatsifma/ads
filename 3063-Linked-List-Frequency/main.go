package p3063

import lib "leetcode/lib"

func frequenciesOfElements(head *lib.ListNode) *lib.ListNode {
	if head == nil {
		return nil
	}

	// Phase 1: O(n) time, O(k) space
	freqs := make(map[int]int)
	curr := head
	for curr != nil {
		freqs[curr.Val]++ // Count frequencies
		curr = curr.Next
	}

	// Phase 2: O(k) time, O(k) space
	dummy := &lib.ListNode{}
	curr = dummy
	for _, v := range freqs { // Random order is fine!
		curr.Next = &lib.ListNode{Val: v}
		curr = curr.Next
	}

	return dummy.Next
}

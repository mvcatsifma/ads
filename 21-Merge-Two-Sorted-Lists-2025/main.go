package main

import lib "leetcode-lib"

func mergeTwoLists(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	head := &lib.ListNode{}
	tail := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}

		tail = tail.Next
	}

	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}

	return head.Next
}

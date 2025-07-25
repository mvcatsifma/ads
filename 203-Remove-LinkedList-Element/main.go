package main

import lib "leetcode-lib"

func removeElements(head *lib.ListNode, val int) *lib.ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

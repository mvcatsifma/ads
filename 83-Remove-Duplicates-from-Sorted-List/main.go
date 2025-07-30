package p83

import (
	"leetcode/lib"
)

func deleteDuplicates(head *lib.ListNode) *lib.ListNode {
	values := make(map[int]bool)
	input := head
	var prev *lib.ListNode
	for head != nil {
		if ok := values[head.Val]; ok {
			prev.Next = head.Next
		} else {
			values[head.Val] = true
		}
		prev = head
		head = head.Next
	}

	return input
}

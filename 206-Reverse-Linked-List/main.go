package p206

import "leetcode/lib"

func reverseList(head *lib.ListNode) *lib.ListNode {
	var prev *lib.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

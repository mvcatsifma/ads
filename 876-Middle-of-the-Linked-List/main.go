package p876

import "leetcode/lib"

func middleNode(head *lib.ListNode) *lib.ListNode {
	var end = head
	var middle = head

	for end != nil && end.Next != nil {
		end = end.Next.Next
		middle = middle.Next
	}

	return middle
}

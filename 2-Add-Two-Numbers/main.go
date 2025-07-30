package p2

import (
	"leetcode/lib"
)

func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	dummy := &lib.ListNode{}
	curr := dummy
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		carry = sum / 10
		curr.Next = &lib.ListNode{Val: sum % 10}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return dummy.Next
}

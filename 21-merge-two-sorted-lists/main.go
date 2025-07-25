package p21

import "leetcode/lib"

func mergeTwoLists(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	result := &lib.ListNode{} // auxiliary linked list
	if l1 == nil && l2 == nil {
		return nil // nothing to do here
	}
	first := result // pointer to first node
	for {
		if l1 == nil { // l1 is exhausted
			result.Val = l2.Val
			result.Next = l2.Next
			break
		}
		if l2 == nil { // l2 is exhausted
			result.Val = l1.Val
			result.Next = l1.Next
			break
		}
		if l1.Val <= l2.Val { // merge next element from l1
			result.Val = l1.Val
			l1 = l1.Next
		} else { // merge next element from l2
			result.Val = l2.Val
			l2 = l2.Next
		}
		result.Next = &lib.ListNode{}
		result = result.Next // proceed cursor
	}
	return first
}

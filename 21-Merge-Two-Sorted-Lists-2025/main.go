package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) hasNext() bool {
	return l.Next != nil
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	head.Next = mergeTwoLists(l1, l2)

	return head
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	var next *ListNode
	for {
		next = head.Next
		head.Next = prev
		prev = head
		if next != nil {
			head = next
		} else {
			break
		}
	}
	return head
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var end = head
	var middle = head

	for end != nil && end.Next != nil {
		end = end.Next.Next
		middle = middle.Next
	}

	return middle
}

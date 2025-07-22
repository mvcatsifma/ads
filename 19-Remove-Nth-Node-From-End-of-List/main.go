package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	h1 := head
	h2 := head

	i := 0
	for i < n-1 && h2.Next != nil {
		h2 = h2.Next
		i++
	}

	var prev *ListNode
	for h2.Next != nil {
		prev = h1
		h1 = h1.Next
		h2 = h2.Next
	}

	if prev != nil {
		prev.Next = h1.Next
		return head
	}

	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

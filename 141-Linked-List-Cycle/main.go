package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	f := head
	s := head

	for s.Next != nil && f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

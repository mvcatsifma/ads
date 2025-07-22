package main

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	i := 0
	nodes := make(map[int]*ListNode)
	for {
		nodes[i] = head
		if head.Next == nil {
			break
		}
		head = head.Next
		i++
	}

	head0 := &ListNode{}
	current := head0
	mid := i / 2
	l := 0
	for m := i; m >= mid; m-- {
		current.Next = nodes[l]
		current.Next.Next = nodes[m]
		current = current.Next.Next
		l++
	}

	current.Next = nil
	head = head0.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

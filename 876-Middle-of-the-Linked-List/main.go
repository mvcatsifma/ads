package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// todo: refactor to not use a map for intermediate storage
func middleNode(head *ListNode) *ListNode {
	m := make(map[int]*ListNode)
	i := 0
	for head != nil {
		m[i] = head
		head = head.Next
		i++
	}
	index := i / 2
	return m[index]
}

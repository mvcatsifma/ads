package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	nodeToIdx := make(map[*ListNode]bool)
	for {
		nodeToIdx[head] = true
		if _, ok := nodeToIdx[head.Next]; ok {
			return true
		}
		head = head.Next
		if head == nil {
			break
		}
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

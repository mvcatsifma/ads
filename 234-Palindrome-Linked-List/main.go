package main

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	var nodes []*ListNode
	for {
		nodes = append(nodes, head)
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	i := 0
	j := len(nodes) - 1

	for i <= j {
		if nodes[i].Val != nodes[j].Val {
			return false
		}
		i++
		j--
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

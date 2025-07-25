package p83

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	values := make(map[int]bool)
	input := head
	var prev *ListNode
	for head != nil {
		if ok := values[head.Val]; ok {
			prev.Next = head.Next
		} else {
			values[head.Val] = true
		}
		prev = head
		head = head.Next
	}

	return input
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v", head)
		head = head.Next
	}
}

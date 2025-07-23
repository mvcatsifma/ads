package main

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head2 := slow.Next
	head2 = reverseLinkedList(head2)

	for {
		if head == nil || head2 == nil {
			break
		}
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}

	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

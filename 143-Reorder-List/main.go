package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	middle := midNode(head)
	newHead := middle.Next
	newHead = reverseList(newHead)
	middle.Next = nil

	c1 := head
	c2 := newHead
	var f1 *ListNode
	var f2 *ListNode

	for c1 != nil && c2 != nil {
		f1 = c1.Next
		f2 = c2.Next

		c1.Next = c2
		c2.Next = f1

		c1 = f1
		c2 = f2
	}
}

func reverseList(head *ListNode) *ListNode {
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

func midNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

//func reorderList(head *ListNode) {
//	if head == nil {
//		return
//	}
//
//	i := 0
//	nodes := make(map[int]*ListNode)
//	for {
//		nodes[i] = head
//		if head.Next == nil {
//			break
//		}
//		head = head.Next
//		i++
//	}
//
//	head0 := &ListNode{}
//	current := head0
//	mid := i / 2
//	l := 0
//	for m := i; m >= mid; m-- {
//		current.Next = nodes[l]
//		current.Next.Next = nodes[m]
//		current = current.Next.Next
//		l++
//	}
//
//	current.Next = nil
//	head = head0.Next
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

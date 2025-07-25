package p143

import "leetcode/lib"

func reorderList(head *lib.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	middle := midNode(head)
	newHead := middle.Next
	newHead = reverseList(newHead)
	middle.Next = nil

	c1 := head
	c2 := newHead
	var f1 *lib.ListNode
	var f2 *lib.ListNode

	for c1 != nil && c2 != nil {
		f1 = c1.Next
		f2 = c2.Next

		c1.Next = c2
		c2.Next = f1

		c1 = f1
		c2 = f2
	}
}

func reverseList(head *lib.ListNode) *lib.ListNode {
	var prev *lib.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func midNode(head *lib.ListNode) *lib.ListNode {
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

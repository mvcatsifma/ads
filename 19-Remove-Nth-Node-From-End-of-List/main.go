package main

import lib "leetcode-lib"

func removeNthFromEnd(head *lib.ListNode, n int) *lib.ListNode {
	dummy := &lib.ListNode{}
	dummy.Next = head
	length := 0
	first := head
	for first != nil {
		length++
		first = first.Next
	}
	i := 0
	first = dummy
	for i != length-n {
		i++
		first = first.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}

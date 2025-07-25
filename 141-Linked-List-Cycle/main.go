package main

import lib "leetcode-lib"

func hasCycle(head *lib.ListNode) bool {
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

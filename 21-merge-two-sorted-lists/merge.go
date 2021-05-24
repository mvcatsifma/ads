package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for x := l1; x != nil; x = x.Next {
		for y := l2; y != nil; y = y.Next {
			if y.hasNext() {
				if x.Val <= y.Next.Val {
					current := y
					next := y.Next
					newNode := &ListNode{
						Val:  x.Val,
						Next: next,
					}
					current.Next = newNode
					break
				}
			} else {

			}
			if y.Next != nil {
			} else {
				newNode := &ListNode{
					Val:  x.Val,
					Next: nil,
				}
				y.Next = newNode
				break
			}
		}
	}

	return l2
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) hasNext() bool {
	return l.Next != nil
}

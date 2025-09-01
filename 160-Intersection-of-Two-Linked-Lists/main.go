package p160

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[*ListNode]int)
	m[headA] = headA.Val
	for headA.Next != nil {
		headA = headA.Next
		m[headA] = headA.Val
	}

	if _, ok := m[headB]; ok {
		return headB
	}

	for headB.Next != nil {
		headB = headB.Next
		if _, ok := m[headB]; ok {
			return headB
		}
	}

	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

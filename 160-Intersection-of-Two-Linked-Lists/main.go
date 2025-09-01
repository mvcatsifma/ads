package p160

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodesInA := make(map[*ListNode]int)
	for headA != nil {
		nodesInA[headA] = headA.Val
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := nodesInA[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

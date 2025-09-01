package p160

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getLength(headA)
	lenB := getLength(headB)

	currA, currB := headA, headB

	if lenA > lenB {
		currA = headA
		for i := 0; i < lenA-lenB; i++ {
			currA = currA.Next
		}
		currB = headB
	} else {
		currB = headB
		for i := 0; i < lenB-lenA; i++ {
			currB = currB.Next
		}
		currA = headA
	}

	for currA != currB {
		currA = currA.Next
		currB = currB.Next
	}

	return currA
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

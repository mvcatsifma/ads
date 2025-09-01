package p160

// getIntersectionNode finds the node where two singly linked lists intersect.
// It returns the intersection node if the lists intersect, or nil if they don't.
//
// The function uses a two-phase approach:
// 1. Calculate the length of both lists
// 2. Align the starting positions by advancing the longer list's pointer
// 3. Traverse both lists simultaneously until finding the intersection
//
// Time complexity: O(m + n) where m and n are the lengths of the lists
// Space complexity: O(1)
//
// Parameters:
//   - headA: pointer to the head of the first linked list
//   - headB: pointer to the head of the second linked list
//
// Returns:
//   - *ListNode: the intersection node if lists intersect, nil otherwise
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

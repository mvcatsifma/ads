package p160

// getIntersectionNode finds the node where two singly linked lists intersect.
// It returns the intersection node if the lists intersect, or nil if they don't.
//
// The function uses the "two-pointer" approach to find the intersection node:
// 1. Initialize two pointers, pA and pB, to the heads of the two lists.
// 2. Traverse both lists simultaneously, moving one pointer at a time.
// 3. If a pointer reaches the end of its list, redirect it to the head of the other list.
// 4. The two pointers will eventually meet at the intersection node, or both will be nil if there is no intersection.
//
// Time complexity: O(m + n), where m and n are the lengths of the two lists.
// Space complexity: O(1), as the function uses only a constant amount of extra space.
//
// Parameters:
//   - headA: pointer to the head of the first linked list
//   - headB: pointer to the head of the second linked list
//
// Returns:
//   - *ListNode: the intersection node if lists intersect, nil otherwise
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var pA, pB = headA, headB

	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}

	return pA
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

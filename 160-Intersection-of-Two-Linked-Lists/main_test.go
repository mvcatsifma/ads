package p160

import (
	"testing"
)

// Test case structure
type args struct {
	intersectVal int
	listA        []int
	listB        []int
	skipA        int
	skipB        int
}

func Test_getIntersectionNode(t *testing.T) {
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Intersection at value 8",
			args: args{
				intersectVal: 8,
				listA:        []int{4, 1, 8, 4, 5},
				listB:        []int{5, 6, 1, 8, 4, 5},
				skipA:        2,
				skipB:        3,
			},
		},
		{
			name: "Intersection at value 2",
			args: args{
				intersectVal: 2,
				listA:        []int{1, 9, 1, 2, 4},
				listB:        []int{3, 2, 4},
				skipA:        3,
				skipB:        1,
			},
		},
		{
			name: "No intersection",
			args: args{
				intersectVal: 0,
				listA:        []int{2, 6, 4},
				listB:        []int{1, 5},
				skipA:        3,
				skipB:        2,
			},
		},
		{
			name: "Single node intersection",
			args: args{
				intersectVal: 1,
				listA:        []int{1},
				listB:        []int{1},
				skipA:        0,
				skipB:        0,
			},
		},
		{
			name: "One empty list",
			args: args{
				intersectVal: 0,
				listA:        []int{},
				listB:        []int{1, 2, 3},
				skipA:        0,
				skipB:        0,
			},
		},
		{
			name: "Both empty lists",
			args: args{
				intersectVal: 0,
				listA:        []int{},
				listB:        []int{},
				skipA:        0,
				skipB:        0,
			},
		},
		{
			name: "Intersection at head of both",
			args: args{
				intersectVal: 3,
				listA:        []int{3, 7, 8},
				listB:        []int{3, 7, 8},
				skipA:        0,
				skipB:        0,
			},
		},
		{
			name: "Different prefix lengths",
			args: args{
				intersectVal: 100,
				listA:        []int{1, 2, 100, 200},
				listB:        []int{99, 100, 200},
				skipA:        2,
				skipB:        1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, expected := createIntersectingLists(tt.args)
			result := getIntersectionNode(headA, headB)

			// Validate result
			if expected != result {
				t.Errorf("Expected intersection node %p, got %p", expected, result)
			}

			// Additional validation for non-nil results
			if tt.args.intersectVal != 0 {
				if result == nil {
					t.Errorf("Expected intersection node with value %d, got nil", tt.args.intersectVal)
				} else if result.Val != tt.args.intersectVal {
					t.Errorf("Expected intersection node value %d, got %d", tt.args.intersectVal, result.Val)
				}
			} else {
				if result != nil {
					t.Errorf("Expected no intersection (nil), got node with value %d", result.Val)
				}
			}
		})
	}
}

// Helper function to create linked list from slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}

	return head
}

// Helper function to get node at specific position
func getNodeAt(head *ListNode, pos int) *ListNode {
	current := head
	for i := 0; i < pos && current != nil; i++ {
		current = current.Next
	}
	return current
}

// Helper function to connect list to tail at specific position
func connectToTail(head, tail *ListNode) {
	if head == nil {
		return
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = tail
}

func createIntersectingLists(tc args) (*ListNode, *ListNode, *ListNode) {
	if tc.intersectVal == 0 {
		// No intersection
		return createList(tc.listA), createList(tc.listB), nil
	}

	// Create the common tail (intersection part) from listA
	intersectionStart := tc.skipA
	commonTail := createList(tc.listA[intersectionStart:])

	// Create list A
	var headA *ListNode
	if tc.skipA > 0 {
		headA = createList(tc.listA[:tc.skipA])
		connectToTail(headA, commonTail)
	} else {
		headA = commonTail
	}

	// Create list B - it has its own prefix + shares the common tail
	var headB *ListNode
	if tc.skipB > 0 {
		// listB should have its own unique prefix, then connect to common tail
		headB = createList(tc.listB[:tc.skipB])
		connectToTail(headB, commonTail)
	} else {
		headB = commonTail
	}

	// Expected intersection node is the start of common tail
	expectedIntersection := commonTail

	return headA, headB, expectedIntersection
}

// Create intersecting linked lists based on test case parameters
//func createIntersectingLists(tc args) (*ListNode, *ListNode, *ListNode) {
//	if tc.intersectVal == 0 {
//		// No intersection
//		return createList(tc.listA), createList(tc.listB), nil
//	}
//
//	// Create the common tail (intersection part)
//	intersectionStart := len(tc.listA) - tc.skipA
//	commonTail := createList(tc.listA[intersectionStart:])
//
//	// Create list A
//	var headA *ListNode
//	if tc.skipA > 0 {
//		headA = createList(tc.listA[:intersectionStart])
//		connectToTail(headA, commonTail)
//	} else {
//		headA = commonTail
//	}
//
//	// Create list B
//	var headB *ListNode
//	if tc.skipB > 0 {
//		headB = createList(tc.listB[:tc.skipB])
//		connectToTail(headB, commonTail)
//	} else {
//		headB = commonTail
//	}
//
//	// Expected intersection node
//	expectedIntersection := getNodeAt(headA, tc.skipA)
//
//	return headA, headB, expectedIntersection
//}

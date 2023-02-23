package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var output *ListNode
	//example1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: nil,
	//		},
	//	},
	//}
	//
	//output = deleteDuplicates(example1)
	//printList(output)
	//
	//example2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val: 3,
	//				Next: &ListNode{
	//					Val:  3,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//output = deleteDuplicates(example2)
	//printList(output)

	example3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	output = deleteDuplicates(example3)

	// todo: failure: head = [1,1,1], output = [1,1], expected [1]
	printList(output)
}

func deleteDuplicates(head *ListNode) *ListNode {
	values := make(map[int]bool)
	input := head
	var prev *ListNode
	for head != nil {
		if ok := values[head.Val]; ok {
			prev.Next = head.Next
		} else {
			values[head.Val] = true
		}
		prev = head
		head = head.Next
	}

	return input
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v", head)
		head = head.Next
	}
}

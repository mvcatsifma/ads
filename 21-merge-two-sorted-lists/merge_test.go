package main

import (
	"fmt"
	"testing"
)

// l1 = [1,2,4], l2 = [1,3,4]
var l1, l2 *ListNode

// l3 = [1], l4 =[]
var l3, l4 *ListNode

// l5 = [2], l6 = [1]
var l5, l6 *ListNode

func init() {
	l12 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l11 := &ListNode{
		Val:  2,
		Next: l12,
	}
	l1 = &ListNode{
		Val:  1,
		Next: l11,
	}

	l22 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l21 := &ListNode{
		Val:  3,
		Next: l22,
	}
	l2 = &ListNode{
		Val:  1,
		Next: l21,
	}

	l3 = &ListNode{
		Val: 1,
		Next: nil,
	}

	l5 = &ListNode{
		Val:  2,
		Next: nil,
	}

	l6 = &ListNode{
		Val:  1,
		Next: nil,
	}
}

func Test_Merge_1(t *testing.T) {
	l := mergeTwoLists(l1, l2)

	for x := l; x != nil; x = x.Next {
		fmt.Println(x.Val)
	}
}

func Test_Merge_2(t *testing.T) {
	l := mergeTwoLists(l3, nil)

	for x := l; x != nil; x = x.Next {
		fmt.Println(x.Val)
	}
}

func Test_Merge_3(t *testing.T) {
	l := mergeTwoLists(l5, l6)

	for x := l; x != nil; x = x.Next {
		fmt.Println(x.Val)
	}
}

func Test_Merge_4(t *testing.T) {
	l := mergeTwoLists(l6, l5)

	for x := l; x != nil; x = x.Next {
		fmt.Println(x.Val)
	}
}

func Test_Merge_5(t *testing.T) {
	l := mergeTwoLists(nil, nil)

	for x := l; x != nil; x = x.Next {
		fmt.Println(x.Val)
	}
}

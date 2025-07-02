package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "case 2",
			args: args{
				list1: []int{},
				list2: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodeList1 := buildLinkedList(tt.args.list1)
			nodeList2 := buildLinkedList(tt.args.list2)
			wantList := buildLinkedList(tt.want)
			if got := mergeTwoLists(nodeList1, nodeList2); !cmp.Equal(got, wantList) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

// buildLinkedList converts a slice of integers into a linked list
// and returns the head of the list
func buildLinkedList(nums []int) *ListNode {
	// Handle empty input
	if len(nums) == 0 {
		return nil
	}

	// Create head node
	head := &ListNode{Val: nums[0]}
	current := head

	// Iterate through remaining numbers
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

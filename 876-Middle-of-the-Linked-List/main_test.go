package main

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildLinkedList(tt.args.nums)
			if got := middleNode(head).Val; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BuildLinkedList converts a slice of integers into a linked list
// and returns the head of the list
func BuildLinkedList(nums []int) *ListNode {
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

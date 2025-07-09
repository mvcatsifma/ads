package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{},
				val:  1,
			},
			want: []int{},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{6, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		head := buildLinkedList(tt.args.nums)
		want := buildLinkedList(tt.want)
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(head, tt.args.val); !cmp.Equal(got, want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
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

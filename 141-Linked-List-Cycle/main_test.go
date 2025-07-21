package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head []int
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				head: []int{3, 2, 0, -4},
				pos:  1,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				head: []int{1, 2},
				pos:  0,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				head: []int{1},
				pos:  -1,
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				head: []int{},
				pos:  -1,
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				head: []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5},
				pos:  -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := buildLinkedList(tt.args.head, tt.args.pos)
			if got := hasCycle(list); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

// buildLinkedList converts a slice of integers into a linked list
// and returns the head of the list
func buildLinkedList(nums []int, pos int) *ListNode {
	// Handle empty input
	if len(nums) == 0 {
		return nil
	}

	// Create head node
	head := &ListNode{Val: nums[0]}
	current := head

	var tailNext *ListNode
	if pos == 0 {
		tailNext = head
	}

	// Iterate through remaining numbers
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
		if i == pos {
			tailNext = current
		}
	}

	if tailNext != nil {
		current.Next = tailNext
	}

	return head
}

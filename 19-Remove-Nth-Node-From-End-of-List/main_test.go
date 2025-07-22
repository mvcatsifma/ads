package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "case 3",
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
		{
			name: "case 3",
			args: args{
				head: []int{1, 2},
				n:    2,
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildLinkedList(tt.args.head)
			want := buildLinkedList(tt.want)
			if got := removeNthFromEnd(head, tt.args.n); !cmp.Equal(got, want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

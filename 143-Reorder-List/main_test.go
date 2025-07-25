package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	lib "leetcode-lib"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 4, 2, 3},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 5, 2, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := lib.CreateLinkedList(tt.args.nums)
			want := lib.CreateLinkedList(tt.want)
			reorderList(head)
			if !cmp.Equal(head, want) {
				t.Errorf("reorderList() = %v, want %v", head, tt.want)
			}
		})
	}
}

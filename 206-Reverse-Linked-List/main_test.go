package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	lib "leetcode-lib"
)

func Test_reverseList(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := lib.CreateLinkedList(tt.args.nums)
			want := lib.CreateLinkedList(tt.want)
			if got := reverseList(head); !cmp.Equal(got, want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

package p2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"leetcode/lib"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "case 2",
			args: args{
				l1: []int{0},
				l2: []int{0},
			},
			want: []int{0},
		},
		{
			name: "case 3",
			args: args{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "case 4",
			args: args{
				l1: []int{5, 6, 3, 0, 8},
				l2: []int{2, 4, 5},
			},
			want: []int{7, 0, 9, 0, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := lib.CreateLinkedList(tt.args.l1)
			l2 := lib.CreateLinkedList(tt.args.l2)
			want := lib.CreateLinkedList(tt.want)
			if got := addTwoNumbers(l1, l2); !cmp.Equal(got, want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

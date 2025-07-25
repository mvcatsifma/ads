package main

import (
	"testing"

	lib "leetcode-lib"
)

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
			name: "case 4",
			args: args{
				head: []int{},
				pos:  -1,
			},
			want: false,
		},
		{
			name: "case 5",
			args: args{
				head: []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5},
				pos:  -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := lib.CreateLinkedListFromNums(tt.args.head, tt.args.pos)
			if got := hasCycle(list); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

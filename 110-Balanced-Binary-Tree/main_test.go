package p110

import (
	"testing"

	"leetcode/lib"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				root: []int{3, 9, 20, -1, -1, 15, 7},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, 2, 2, 3, 3, -1, -1, 4, 4},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				root: []int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		root := lib.BuildTree(tt.args.root)

		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

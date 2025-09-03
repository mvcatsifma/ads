package p543

import (
	"testing"

	"leetcode/lib"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		root := lib.BuildTree(tt.args.root)

		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

package p235

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root []int
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:    2,
				q:    8,
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:    2,
				q:    4,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				root: []int{2, 1},
				p:    2,
				q:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			p := &lib.TreeNode{Val: tt.args.p}
			q := &lib.TreeNode{Val: tt.args.q}
			got := lowestCommonAncestor(root, p, q)
			if got.Val != tt.want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}

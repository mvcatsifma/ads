package p105

import (
	"math"
	"testing"

	"leetcode/lib"

	"github.com/google/go-cmp/cmp"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
		},
		{
			name: "case 2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := lib.BuildTree(tt.want)
			if got := buildTree(tt.args.preorder, tt.args.inorder); !cmp.Equal(got, want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

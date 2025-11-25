package p102

import (
	"math"
	"testing"

	"leetcode/lib"

	"github.com/google/go-cmp/cmp"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			name: "case 2",
			args: args{
				root: []int{1},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "case 3",
			args: args{
				root: []int{},
			},
			want: [][]int{},
		},
		{
			name: "case 4",
			args: args{
				root: []int{3, 9, 20, 5, 6, 15, 7},
			},
			want: [][]int{
				{3},
				{9, 20},
				{5, 6, 15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := levelOrder(root); !cmp.Equal(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

package p199

import (
	"math"
	"testing"

	"leetcode/lib"

	"github.com/google/go-cmp/cmp"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{1, 2, 3, math.MaxInt, 5, math.MaxInt, 4},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, 2, 3, 4, math.MaxInt, math.MaxInt, math.MaxInt, 5},
			},
			want: []int{1, 3, 4, 5},
		},
		{
			name: "case 3",
			args: args{
				root: []int{1, math.MaxInt, 3},
			},
			want: []int{1, 3},
		},
		{
			name: "case 4",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := rightSideView(root); !cmp.Equal(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightSideViewIterative(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{1, 2, 3, math.MaxInt, 5, math.MaxInt, 4},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, 2, 3, 4, math.MaxInt, math.MaxInt, math.MaxInt, 5},
			},
			want: []int{1, 3, 4, 5},
		},
		{
			name: "case 3",
			args: args{
				root: []int{1, math.MaxInt, 3},
			},
			want: []int{1, 3},
		},
		{
			name: "case 4",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := rightSideViewIterative(root); !cmp.Equal(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}

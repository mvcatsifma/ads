package p226

import (
	"testing"

	"leetcode/lib"

	"github.com/google/go-cmp/cmp"
)

func Test_invertTree(t *testing.T) {
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
				root: []int{4, 2, 7, 1, 3, 6, 9},
			},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "case 2",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
		{
			name: "case 3",
			args: args{
				root: []int{2, 1, 3},
			},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			want := lib.BuildTree(tt.want)
			if got := invertTree(root); !cmp.Equal(got, want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

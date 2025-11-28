package p000Trees

import (
	"math"
	"reflect"
	"testing"

	"leetcode/lib"
)

func Test_dfsPreorder(t *testing.T) {
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
				root: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := dfsPreorder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsInorder(t *testing.T) {
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
				root: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: []int{9, 3, 15, 20, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := dfsInorder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsInorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsPostorder(t *testing.T) {
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
				root: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: []int{9, 15, 7, 20, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := dfsPostorder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfsPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bfs(t *testing.T) {
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
				root: []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7},
			},
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := bfs(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

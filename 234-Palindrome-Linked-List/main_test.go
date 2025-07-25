package p234

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	lib "leetcode/lib"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 2, 1},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 2, 2, 2, 1},
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := lib.CreateLinkedList(tt.args.nums)
			if got := isPalindrome(head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseLinkedList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := lib.CreateLinkedList(tt.args.nums)
			want := lib.CreateLinkedList(tt.want)
			if got := reverseLinkedList(head); !cmp.Equal(got, want) {
				t.Errorf("reverseLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

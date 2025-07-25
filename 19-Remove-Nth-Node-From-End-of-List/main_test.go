package p19

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	lib "leetcode/lib"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "case 3",
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
		{
			name: "case 3",
			args: args{
				head: []int{1, 2},
				n:    2,
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := lib.CreateLinkedList(tt.args.head)
			want := lib.CreateLinkedList(tt.want)
			if got := removeNthFromEnd(head, tt.args.n); !cmp.Equal(got, want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

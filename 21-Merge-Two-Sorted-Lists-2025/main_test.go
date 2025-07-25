package p21_2025

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	lib "leetcode/lib"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "case 2",
			args: args{
				list1: []int{},
				list2: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodeList1 := lib.CreateLinkedList(tt.args.list1)
			nodeList2 := lib.CreateLinkedList(tt.args.list2)
			wantList := lib.CreateLinkedList(tt.want)
			if got := mergeTwoLists(nodeList1, nodeList2); !cmp.Equal(got, wantList) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

package p876

import (
	"reflect"
	"testing"

	lib "leetcode/lib"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := lib.CreateLinkedList(tt.args.nums)
			if got := middleNode(head).Val; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

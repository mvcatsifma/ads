package p27

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{3, 2, 2, 3}, val: 3}, want: []int{2, 2}},
		{args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, want: []int{0, 1, 3, 0, 4}},
		{args: args{nums: []int{0}, val: 2}, want: []int{0}},
		{args: args{nums: []int{2}, val: 2}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeElement(tt.args.nums, tt.args.val)
			if k != len(tt.want) {
				t.Errorf("removeElement() = %v, want %v", k, len(tt.want))
			}

			for i := 0; i < k; i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("removeElement() = %v, want %v", tt.args.nums[i], tt.want[i])
				}
			}
		})
	}
}

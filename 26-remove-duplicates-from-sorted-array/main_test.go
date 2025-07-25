package p26

import "testing"

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{1, 1, 2}}, want: []int{1, 2}},
		{args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.args.nums)
			if k != len(tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", k, len(tt.want))
			}

			for i := 0; i < k; i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums[i], tt.want[i])
				}
			}
		})
	}
}

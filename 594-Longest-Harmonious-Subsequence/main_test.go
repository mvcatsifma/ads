package p594

import "testing"

func Test_findLHS(t *testing.T) {
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
				nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1, 2, 2, 3, 4, 5, 1, 1, 1, 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.args.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}

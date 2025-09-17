package p1046

import "testing"

func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				stones: []int{2, 7, 4, 1, 8, 1},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				stones: []int{1},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				stones: []int{2, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

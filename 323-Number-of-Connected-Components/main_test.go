package p323

import "testing"

func Test_countComponents(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 5,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{3, 4},
				},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				n: 5,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{2, 3},
					{3, 4},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countComponents(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

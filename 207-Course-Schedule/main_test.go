package p207

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: true,
		},
		{
			name: "case 2 - cycle",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
		{
			name: "case 3 - complex cycle",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{1, 0},
					{2, 1},
					{0, 2},
				},
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{0, 1},
				},
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				numCourses: 5,
				prerequisites: [][]int{
					{1, 4},
					{2, 4},
					{3, 1},
					{3, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

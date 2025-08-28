package p28

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				haystack: "badbutsad",
				needle:   "sad",
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "case 4",
			args: args{
				haystack: "aaa",
				needle:   "aaaa",
			},
			want: -1,
		},
		{
			name: "case 5",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			name: "case 6",
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

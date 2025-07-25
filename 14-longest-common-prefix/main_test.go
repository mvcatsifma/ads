package p14

import "testing"

func Fuzz_longestCommonPrefix(f *testing.F) {
	f.Fuzz(func(t *testing.T, s1 string, s2 string, s3 string, s4 string) {
		longestCommonPrefix([]string{s1, s2, s3, s4})
	})
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "fl", args: args{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "empty", args: args{strs: []string{"dog", "racecar", "car"}}, want: ""},
		{name: "a", args: args{strs: []string{"ab", "a"}}, want: "a"},
		{name: "", args: args{strs: []string{"reflower", "flow", "flight"}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

package p49

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case 1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"},
			},
		},
		{
			name: "case 2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{
				{""},
			},
		},
		{
			name: "case 3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

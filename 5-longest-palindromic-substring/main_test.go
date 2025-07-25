package p5

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "Example 1",
			in:       "babad",
			expected: "bab",
		},
		{
			name:     "Example 2",
			in:       "cbbd",
			expected: "bb",
		},
		{
			name:     "Example 3",
			in:       "a",
			expected: "a",
		},
		{
			name:     "Example 4",
			in:       "ac",
			expected: "a",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := longestPalindrome(test.in)

			if actual != test.expected {
				t.Errorf("expected %s, actual %s", test.expected, actual)
			}
		})
	}
}

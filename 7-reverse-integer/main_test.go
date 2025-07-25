package p7

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name     string
		in       int
		expected int
	}{
		{
			name:     "Example 1",
			in:       123,
			expected: 321,
		},
		{
			name:     "Example 2",
			in:       -123,
			expected: -321,
		},
		{
			name:     "Example 3",
			in:       120,
			expected: 21,
		},
		{
			name:     "Example 4",
			in:       0,
			expected: 0,
		},
		{
			name:     "TestCase 12",
			in:       1534236469,
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := reverse(test.in)

			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

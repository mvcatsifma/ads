package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name     string
		in       int
		expected bool
	}{
		{
			name:     "Example 1",
			in:       121,
			expected: true,
		},
		{
			name:     "Example 2",
			in:       -121,
			expected: false,
		},
		{
			name:     "Example 3",
			in:       10,
			expected: false,
		},
		{
			name:     "Example 4",
			in:       -101,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isPalindrome2(test.in)

			if actual != test.expected {
				t.Errorf("expected: %t, got: %t\n", test.expected, actual)
			}
		})
	}
}

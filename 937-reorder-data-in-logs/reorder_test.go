package main

import (
	"reflect"
	"testing"
)

func Test_Reorder_1(t *testing.T) {
	reorderTests := []struct {
		name     string
		log      []string
		expected []string
	}{
		{
			name:     "Example 1",
			log:      []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			expected: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			name:     "Example 2",
			log:      []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			expected: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			name:     "Testcase 1",
			log:      []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"},
			expected: []string{"b aq cojj", "5 ba iedrz", "8 fj dzz k", "63 gu psub", "bb wsrd kp", "5r 446 6 3", "6s 87979 5", "3r 587 01", "jc 3480612", "r5 6316 71"},
		},
	}

	for _, tt := range reorderTests {
		t.Run(tt.name, func(t *testing.T) {
			var actual = reorderLogFiles(tt.log)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("exptected %#v actual %#v", tt.expected, actual)
			}
		})
	}
}

func TestParse(t *testing.T) {
	parseTests := []struct {
		log        string
		logType    LogType
		identifier string
		words      string
	}{
		{"dig1 8 1 5 1", Digit, "dig1", "8 1 5 1"},
		{"dig2 31 6", Digit, "dig2", "31 6"},
		{"let1 art can", Letter, "let1", "art can"},
		{"let3 art zero", Letter, "let3", "art zero"},
	}

	for _, tt := range parseTests {
		actual := Parse(tt.log)

		if actual.LogType != tt.logType {
			t.Errorf("expected a digit-log")
		}
		if actual.Identifier != tt.identifier {
			t.Errorf("expected %q, got %q", tt.identifier, actual.Identifier)
		}
		if actual.Words != tt.words {
			t.Errorf("expected %q got %q", tt.words, actual.Words)
		}
	}
}

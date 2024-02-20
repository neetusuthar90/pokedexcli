package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := CleanInput(cs.input)

		if len(actual) != len(cs.expected) {
			t.Errorf("The length are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not equal to %v", actualWord, expectedWord)
			}
		}
	}
}

package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this   has no left whitespaces     ",
			expected: []string{"this", "has", "no", "left", "whitespaces"},
		},
		{
			input:    "   this  has    no right whitespaces",
			expected: []string{"this", "has", "no", "right", "whitespaces"},
		},
		{
			input:    "this has    no whitespaces around the block   of  text",
			expected: []string{"this", "has", "no", "whitespaces", "around", "the", "block", "of", "text"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		if len(actual) != len(c.expected) {
			t.Errorf("Test case failed: \nCase: \n%s", c.input)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test

			if word != expectedWord {
				t.Errorf("Words did not match. %s != %s", word, expectedWord)
			}
		}
	}
}

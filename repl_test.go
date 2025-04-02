package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ZIYAD  ALMUBARAK ",
			expected: []string{"ziyad", "almubarak"},
		},
		{
			input:    "             OnE            tWo ",
			expected: []string{"one", "two"},
		},
		// add more cases here
	}
	for i, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			fmt.Println("ACTUAL", actual)
			t.Errorf("Test %v - '%v' FAIL: Length of expected '%v' and actual '%v' are not equal", i, c.input, len(c.expected), len(actual))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Test %v - FAIL: %v is not equal to %v", i, word, expectedWord)
				return
			}
		}
	}
}

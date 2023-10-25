package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_String_IsEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "string with windows newline is empty",
			input:    "\r\n",
			expected: true,
		},
		{
			name:     "string is zero",
			input:    "",
			expected: true,
		},
		{
			name:     "string with linux newline is empty",
			input:    "\n",
			expected: true,
		},
		{
			name:     "string with space is empty",
			input:    " ",
			expected: true,
		},
		{
			name:     "string with non-spaces is not empty",
			input:    "   a string   ",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.IsEmpty(tc.input)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

func Test_String_FindNonEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		ok       bool
		expected string
	}{
		{
			name:     "string with linux newline is empty",
			input:    []string{"\n", "something"},
			expected: "something",
			ok:       true,
		},
		{
			name:     "string with linux newline is empty",
			input:    []string{"\n"},
			expected: "",
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, ok := stringer.FindNonEmpty(tc.input...)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

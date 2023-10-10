package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Truncate_End(t *testing.T) {
	testCases := []struct {
		name     string
		length   int
		input    string
		expected string
		err      error
	}{
		{
			name:     "truncate end long string",
			length:   16,
			input:    "a test string that is too long",
			expected: "a test string...",
		},
		{
			name:     "truncate end short string",
			length:   16,
			input:    "a test string",
			expected: "a test string",
		},
		{
			name:     "truncate end equal length",
			length:   13,
			input:    "a test string",
			expected: "a test string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.TruncateEnd(tc.input, tc.length)
			assert.Equal(tt, string(tc.expected), string(actual))
		})
	}
}

func Test_Truncate_Begin(t *testing.T) {
	testCases := []struct {
		name     string
		length   int
		input    string
		expected string
		err      error
	}{
		{
			name:     "truncate end long string",
			length:   16,
			input:    "a test string that is too long",
			expected: "...t is too long",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.TruncateBegin(tc.input, tc.length)
			assert.Equal(tt, string(tc.expected), string(actual))
		})
	}
}

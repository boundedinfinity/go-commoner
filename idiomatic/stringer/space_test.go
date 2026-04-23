package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_stringer_compact_space(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "A__B",
			input:    "A  B",
			expected: "A B",
		},
		{
			name:     "_A__B_",
			input:    " A  B ",
			expected: " A B ",
		},
		{
			name:     "____A__B__C____",
			input:    "    A  B  C    ",
			expected: "    A B C    ",
		},
		{
			name:     "<empty>",
			input:    "",
			expected: "",
		},
		{
			name:     "___",
			input:    "   ",
			expected: " ",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.CompactSpace(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_FirstNonSpace(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "case 1",
			input:    "  x",
			expected: 2,
		},
		{
			name:     "case 2",
			input:    "",
			expected: 0,
		},
		{
			name:     "case 3",
			input:    "\n",
			expected: 0,
		},
		{
			name:     "case 4",
			input:    "\t",
			expected: 0,
		},
		{
			name:     "case 5",
			input:    "  ",
			expected: 0,
		},
		{
			name:     "case 6",
			input:    "    x  ",
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.FirstNonSpace(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

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

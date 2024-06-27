package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_TabToSpace(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		tabWidth int
		expected string
	}{
		{
			name:     "case 1",
			input:    "x",
			tabWidth: 4,
			expected: "x",
		},
		{
			name:     "case 2",
			input:    "\tx",
			tabWidth: 4,
			expected: "    x",
		},
		{
			name:     "case 3",
			input:    "\t\tx",
			tabWidth: 4,
			expected: "        x",
		},
		{
			name:     "case 4",
			input:    "\tx\t",
			tabWidth: 4,
			expected: "    x    ",
		},
		{
			name:     "case 5",
			input:    "x",
			tabWidth: 2,
			expected: "x",
		},
		{
			name:     "case 6",
			input:    "  x",
			tabWidth: 2,
			expected: "  x",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.TabToSpace(tc.input, tc.tabWidth)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Stringer_Transform(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "trim mulitple spaces",
			input:    "   a string   ",
			expected: "a string",
		},
		{
			name:     "trim single space",
			input:    " a string ",
			expected: "a string",
		},
		{
			name:     "trim no space",
			input:    "a string",
			expected: "a string",
		},
		{
			name:     "trim tabs space",
			input:    "\ta string\t",
			expected: "a string",
		},
		{
			name:     "trim linux new line space",
			input:    "\na string\n",
			expected: "a string",
		},
		{
			name:     "trim widnows new line space",
			input:    "\r\na string\r\n",
			expected: "a string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.TrimSpace(tc.input)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_NoneOf(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		list     []string
		expected bool
	}{
		{
			name:     "case 1",
			input:    "z",
			list:     []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "case 2",
			input:    "a",
			list:     []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "case 3",
			input:    "A",
			list:     []string{"a", "b", "c"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.NoneOf(tc.input, tc.list...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_NoneOfIgnoreCase(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		list     []string
		expected bool
	}{
		{
			name:     "case 1",
			input:    "a",
			list:     []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "case 2",
			input:    "z",
			list:     []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "case 3",
			input:    "A",
			list:     []string{"a", "b", "c"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.NoneOfIgnoreCase(tc.input, tc.list...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

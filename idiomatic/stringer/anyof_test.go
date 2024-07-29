package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_AnyOf(t *testing.T) {
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
			expected: true,
		},
		{
			name:     "case 2",
			input:    "z",
			list:     []string{"a", "b", "c"},
			expected: false,
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
			actual := stringer.AnyOf(tc.input, tc.list...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_AnyOfIgnoreCase(t *testing.T) {
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
			expected: true,
		},
		{
			name:     "case 2",
			input:    "z",
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
			actual := stringer.AnyOfIgnoreCase(tc.input, tc.list...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

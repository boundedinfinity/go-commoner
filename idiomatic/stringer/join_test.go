package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Join_string(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		sep      string
		expected string
	}{
		{
			name:     "case 1",
			input:    []string{"a", "b", "c"},
			sep:      ", ",
			expected: "a, b, c",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.Join(tc.sep, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Join_int(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		sep      string
		expected string
	}{
		{
			name:     "case 1",
			input:    []int{1, 2, 3},
			sep:      ", ",
			expected: "1, 2, 3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := stringer.Join(tc.sep, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

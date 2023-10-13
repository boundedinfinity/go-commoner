package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	testCases := []struct {
		name     string
		match    string
		list     []string
		expected bool
	}{
		{
			name:     "contains a",
			match:    "a",
			list:     []string{"a", "a", "a", "a"},
			expected: true,
		},
		{
			name:     "does not contain a",
			match:    "a",
			list:     []string{"b", "c", "d", "e"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Contains(tc.match, tc.list...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

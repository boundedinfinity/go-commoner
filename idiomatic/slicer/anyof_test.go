package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_AnyOf(t *testing.T) {
	testCases := []struct {
		name     string
		match    string
		input    []string
		expected bool
	}{
		{
			name:     "contains a",
			match:    "a",
			input:    []string{"b", "c", "b", "a"},
			expected: true,
		},
		{
			name:     "does not contain a",
			match:    "a",
			input:    []string{"b", "c", "d", "e"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.AnyOf(tc.match, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

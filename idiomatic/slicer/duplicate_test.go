package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Times(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		count    int
		expected []string
	}{
		{
			name:     "duplicates",
			input:    []string{"a"},
			count:    2,
			expected: []string{"a", "a"},
		},
		{
			name:     "duplicates",
			input:    []string{"a", "b"},
			count:    2,
			expected: []string{"a", "b", "a", "b"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Duplicate(tc.count, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

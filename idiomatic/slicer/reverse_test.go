package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "group by a",
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "empty input",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "nil input",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Reverse(tc.input...)
			assert.ElementsMatch(t, tc.expected, actual, tc.name)
		})
	}
}

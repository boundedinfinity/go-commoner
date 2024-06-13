package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Flatten(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "flatten 1",
			input:    [][]int{{1}, {2}},
			expected: []int{1, 2},
		},
		{
			name:     "flatten 2",
			input:    [][]int{{1, 2}, {3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "flatten 3",
			input:    [][]int{{1, 2}, {}},
			expected: []int{1, 2},
		},
		{
			name:     "flatten 4",
			input:    [][]int{{1, 2}, nil},
			expected: []int{1, 2},
		},
		{
			name:     "flatten 4",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "flatten 5",
			input:    [][]int{nil},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Flatten(tc.input...)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

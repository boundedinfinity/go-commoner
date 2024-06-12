package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		size     int
		expected [][]int
	}{
		{
			name:     "Even chunks",
			input:    []int{1, 2, 3, 4},
			size:     2,
			expected: [][]int{{1, 2}, {3, 4}},
		},
		{
			name:     "Uneven chunks",
			input:    []int{1, 2, 3, 4, 5},
			size:     2,
			expected: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:     "size of 1",
			input:    []int{1, 2, 3, 4, 5},
			size:     1,
			expected: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name:     "size of 0",
			input:    []int{1, 2, 3, 4, 5},
			size:     0,
			expected: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name:     "size of -1",
			input:    []int{1, 2, 3, 4, 5},
			size:     -1,
			expected: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Chunk(tc.size, tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

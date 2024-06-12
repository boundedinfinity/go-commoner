package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Prepend(t *testing.T) {
	testCases := []struct {
		name     string
		input1   []int
		input2   []int
		expected []int
	}{
		{
			name:     "case 1",
			input1:   []int{2, 3},
			input2:   []int{1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "case 2",
			input1:   []int{2, 3},
			input2:   []int{},
			expected: []int{2, 3},
		},
		{
			name:     "case 3",
			input1:   []int{},
			input2:   []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "case 4",
			input1:   []int{},
			input2:   []int{},
			expected: []int{},
		},
		{
			name:     "case 5",
			input1:   nil,
			input2:   []int{},
			expected: []int{},
		},
		{
			name:     "case 6",
			input1:   []int{},
			input2:   nil,
			expected: []int{},
		},
		{
			name:     "case 7",
			input1:   nil,
			input2:   nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Prepend(tc.input1, tc.input2...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

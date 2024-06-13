package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Sort(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "case 1",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "case 2",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "case 3",
			input:    nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Sort(tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

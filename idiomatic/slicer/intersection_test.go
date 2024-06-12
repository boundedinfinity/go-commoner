package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Intersection(t *testing.T) {
	testCases := []struct {
		name     string
		input1   []string
		input2   []string
		expected []string
	}{
		{
			name:     "no  duplicates",
			input1:   []string{"a", "a", "a", "a"},
			input2:   []string{"a", "a", "a", "a"},
			expected: []string{"a"},
		},
		{
			name:     "no duplicates",
			input1:   []string{"a", "b", "c", "d"},
			input2:   []string{"c", "d", "e", "f"},
			expected: []string{"c", "d"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Intersection(tc.input1, tc.input2)
			assert.ElementsMatch(t, tc.expected, actual, tc.name)
		})
	}
}

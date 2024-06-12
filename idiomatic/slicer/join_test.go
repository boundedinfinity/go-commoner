package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Join(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		sep      string
		expected string
	}{
		{
			name:     "join 1",
			input:    []int{1, 2, 3},
			sep:      ",",
			expected: "1,2,3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Join(tc.sep, tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

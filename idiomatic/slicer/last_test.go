package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Last(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
		ok       bool
	}{
		{
			name:     "case 1",
			input:    []int{1, 2, 3},
			expected: 3,
			ok:       true,
		},
		{
			name:     "case 1",
			input:    []int{},
			expected: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, ok := slicer.Last(tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
			assert.Equal(t, tc.ok, ok, tc.name)
		})
	}
}

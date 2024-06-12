package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Tail(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
		ok       bool
	}{
		{
			name:     "case 1",
			input:    []int{1, 2, 3},
			expected: []int{2, 3},
			ok:       true,
		},
		{
			name:     "case 2",
			input:    []int{},
			expected: []int{},
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, ok := slicer.Tail(tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
			assert.Equal(t, tc.ok, ok, tc.name)
		})
	}
}

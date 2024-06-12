package slicer_test

import (
	"fmt"
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
			actual := slicer.Times(tc.count, tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

func Test_TimesFn(t *testing.T) {
	fn := func(i, j int, s string) string { return fmt.Sprintf("%v:%v:%v", i, j, s) }

	testCases := []struct {
		name     string
		input    []string
		count    int
		fn       func(int, int, string) string
		expected []string
	}{
		{
			name:     "duplicates",
			input:    []string{"a"},
			count:    2,
			fn:       fn,
			expected: []string{"0:0:a", "1:0:a"},
		},
		{
			name:     "duplicates",
			input:    []string{"a", "b"},
			count:    2,
			fn:       fn,
			expected: []string{"0:0:a", "0:1:b", "1:0:a", "1:1:b"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.TimesFn(tc.fn, tc.count, tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

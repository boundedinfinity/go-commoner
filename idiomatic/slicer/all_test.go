package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_All(t *testing.T) {
	testCases := []struct {
		name     string
		match    string
		list     []string
		expected bool
	}{
		{
			name:     "a to all a's",
			match:    "a",
			list:     []string{"a", "a", "a", "a"},
			expected: true,
		},
		{
			name:     "a to all a's and b's",
			match:    "a",
			list:     []string{"a", "b", "a", "a"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.All(tc.match, tc.list...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_AllFn(t *testing.T) {
	testCases := []struct {
		name     string
		match    func(int, string) bool
		list     []string
		expected bool
	}{
		{
			name:     "a to all a's",
			match:    func(_ int, elem string) bool { return elem == "a" },
			list:     []string{"a", "a", "a", "a"},
			expected: true,
		},
		{
			name:     "a to all a's and b's",
			match:    func(_ int, elem string) bool { return elem == "a" },
			list:     []string{"a", "b", "a", "a"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.AllFn(tc.match, tc.list...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

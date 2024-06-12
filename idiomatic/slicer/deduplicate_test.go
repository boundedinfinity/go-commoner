package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Deduplicate(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "with  duplicates",
			input:    []string{"a", "a", "a", "a"},
			expected: []string{"a"},
		},
		{
			name:     "no duplicates",
			input:    []string{"a", "b", "c", "d"},
			expected: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Deduplicate(tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

func Test_DeduplicateFn(t *testing.T) {
	type Thing struct {
		K string
		V int
	}

	testCases := []struct {
		name     string
		input    []Thing
		expected []Thing
	}{
		{
			name: "with  duplicates",
			input: []Thing{
				{K: "x", V: 100},
				{K: "y", V: 200},
				{K: "z", V: 300},
				{K: "x", V: 100},
				{K: "y", V: 200},
				{K: "z", V: 300},
			},
			expected: []Thing{
				{K: "x", V: 100},
				{K: "y", V: 200},
				{K: "z", V: 300},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			fn := func(_ int, elem Thing) string { return elem.K }
			actual := slicer.DeduplicateFn(fn, tc.input...)
			assert.Equal(t, tc.expected, actual, tc.name)
		})
	}
}

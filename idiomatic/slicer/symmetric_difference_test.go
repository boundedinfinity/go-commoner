package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_SymmetricDifference(t *testing.T) {
	testCases := []struct {
		name     string
		input1   []int
		input2   []int
		expected []int
	}{
		{
			name:     "with some overlap",
			input1:   []int{1, 2, 3},
			input2:   []int{3, 4, 5},
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "with no overlap",
			input1:   []int{1, 2, 3},
			input2:   []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "with all overlap",
			input1:   []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.SymmetricDifference(tc.input1, tc.input2)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func Test_SymmetricDifferenceFn(t *testing.T) {
	type Thing struct {
		K string
		V int
	}

	testCases := []struct {
		name     string
		input1   []Thing
		input2   []Thing
		expected []Thing
	}{
		{
			name: "with some overlap",
			input1: []Thing{
				{K: "a", V: 100},
				{K: "b", V: 200},
				{K: "c", V: 300},
			},
			input2: []Thing{
				{K: "c", V: 300},
				{K: "d", V: 400},
				{K: "e", V: 500},
			},
			expected: []Thing{
				{K: "a", V: 100},
				{K: "b", V: 200},
				{K: "d", V: 400},
				{K: "e", V: 500},
			},
		},
		{
			name: "with no overlap",
			input1: []Thing{
				{K: "a", V: 100},
				{K: "b", V: 200},
				{K: "c", V: 300},
			},
			input2: []Thing{
				{K: "d", V: 400},
				{K: "e", V: 500},
				{K: "f", V: 600},
			},
			expected: []Thing{
				{K: "a", V: 100},
				{K: "b", V: 200},
				{K: "c", V: 300},
				{K: "d", V: 400},
				{K: "e", V: 500},
				{K: "f", V: 600},
			},
		},
		{
			name: "with all overlap",
			input1: []Thing{
				{K: "a", V: 100},
				{K: "b", V: 200},
				{K: "c", V: 300},
			},
			input2: []Thing{
				{K: "a", V: 100},
				{K: "b", V: 200},
				{K: "c", V: 300},
			},
			expected: []Thing{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			fn := func(_ int, elem Thing) string { return elem.K }
			actual := slicer.SymmetricDifferenceFn(fn, tc.input1, tc.input2)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

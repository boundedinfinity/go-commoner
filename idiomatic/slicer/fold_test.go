package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Fold_SameType(t *testing.T) {
	fn := func(_ int, v int, a int) int { return a + v }

	testCases := []struct {
		name     string
		input    []int
		initial  int
		expected int
		fn       func(int, int, int) int
	}{
		{
			name:     "case 1",
			input:    []int{1, 2, 3, 4, 5},
			initial:  0,
			expected: 15,
			fn:       fn,
		},
		{
			name:     "case 2",
			input:    []int{},
			initial:  0,
			expected: 0,
			fn:       fn,
		},
		{
			name:     "case 3",
			input:    []int{1, 2, 3, 4, 5},
			initial:  10,
			expected: 25,
			fn:       fn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Fold(tc.fn, tc.initial, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_Fold_DiffType(t *testing.T) {
	fn := func(_ int, v int, a byte) int {
		if v == 0 {
			return int(a)
		} else {
			return v
		}
	}

	testCases := []struct {
		name     string
		input    []byte
		initial  int
		expected int
		fn       func(int, int, byte) int
	}{
		{
			name:     "case 1",
			input:    []byte{1, 2, 3, 4, 5},
			initial:  0,
			expected: 1,
			fn:       fn,
		},
		{
			name:     "case 2",
			input:    []byte{},
			initial:  0,
			expected: 0,
			fn:       fn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Fold(tc.fn, tc.initial, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func Test_FoldRight_DiffType(t *testing.T) {
	fn := func(_ int, v int, a byte) int {
		if v == 0 {
			return int(a)
		} else {
			return v
		}
	}

	testCases := []struct {
		name     string
		input    []byte
		initial  int
		expected int
		fn       func(int, int, byte) int
	}{
		{
			name:     "case 1",
			input:    []byte{1, 2, 3, 4, 5},
			initial:  0,
			expected: 5,
			fn:       fn,
		},
		{
			name:     "case 2",
			input:    []byte{},
			initial:  0,
			expected: 0,
			fn:       fn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.FoldRight(tc.fn, tc.initial, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

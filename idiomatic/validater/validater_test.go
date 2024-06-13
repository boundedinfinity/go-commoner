package validater_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	testCases := []struct {
		name      string
		input     any
		validator string
		expected  [][]int
	}{
		// {
		// 	name:     "Even chunks",
		// 	input:    []int{1, 2, 3, 4},
		// 	size:     2,
		// 	expected: [][]int{{1, 2}, {3, 4}},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := slicer.Chunk(tc.size, tc.input...)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

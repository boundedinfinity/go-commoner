package slices_test

import (
	"testing"

	"github.com/boundedinfinity/commons/slices"
	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}, {7}}
	input := []int{1, 2, 3, 4, 5, 6, 7}
	actual := slices.Chunk(input, 2)

	assert.Equal(t, expected, actual)
}

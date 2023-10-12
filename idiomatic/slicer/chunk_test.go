package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}, {7}}
	input := []int{1, 2, 3, 4, 5, 6, 7}
	actual := slicer.Chunk(2, input...)

	assert.Equal(t, expected, actual)
}

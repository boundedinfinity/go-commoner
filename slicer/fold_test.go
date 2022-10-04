package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_FoldLeft_sameType(t *testing.T) {
	expected := 15
	input := []int{1, 2, 3, 4, 5}

	actual := slicer.FoldLeft(0, input, func(v int, a int) int {
		return a + v
	})

	assert.Equal(t, expected, actual)
}

func Test_FoldLeft_diffType(t *testing.T) {
	expected := 1
	input := []byte{1, 2, 3, 4, 5}

	actual := slicer.FoldLeft(0, input, func(v int, a byte) int {
		if v == 0 {
			return int(a)
		} else {
			return v
		}
	})

	assert.Equal(t, expected, actual)
}

func Test_FoldRight_diffType(t *testing.T) {
	expected := 5
	input := []byte{1, 2, 3, 4, 5}

	actual := slicer.FoldRight(0, input, func(v int, a byte) int {
		if v == 0 {
			return int(a)
		} else {
			return v
		}
	})

	assert.Equal(t, expected, actual)
}

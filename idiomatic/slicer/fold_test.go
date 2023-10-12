package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_FoldLeft_sameType(t *testing.T) {
	expected := 15
	input := []int{1, 2, 3, 4, 5}
	fn := func(v int, a int) int { return a + v }

	actual := slicer.FoldLeft(0, fn, input...)
	assert.Equal(t, expected, actual)
}

func Test_FoldLeft_diffType(t *testing.T) {
	expected := 1
	input := []byte{1, 2, 3, 4, 5}
	fn := func(v int, a byte) int {
		if v == 0 {
			return int(a)
		} else {
			return v
		}
	}

	actual := slicer.FoldLeft(0, fn, input...)
	assert.Equal(t, expected, actual)
}

func Test_FoldRight_diffType(t *testing.T) {
	expected := 5
	input := []byte{1, 2, 3, 4, 5}
	fn := func(v int, a byte) int {
		if v == 0 {
			return int(a)
		} else {
			return v
		}
	}

	actual := slicer.FoldRight(0, fn, input...)
	assert.Equal(t, expected, actual)
}

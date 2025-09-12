package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_DupCount(t *testing.T) {
	expected := []slicer.CountDuplicatesResult[int]{
		{Count: 2, Item: 1},
		{Count: 2, Item: 2},
		{Count: 2, Item: 3},
	}
	input := []int{1, 1, 2, 2, 3, 3}
	actual := slicer.CountDuplicates(input...)

	assert.ElementsMatch(t, expected, actual)
}

func Test_DupCountFn(t *testing.T) {
	type Thing struct {
		K string
		V int
	}

	thing1 := Thing{K: "x", V: 100}
	thing2 := Thing{K: "y", V: 200}
	thing3 := Thing{K: "z", V: 300}

	expected := []slicer.CountDuplicatesResult[Thing]{
		{Item: thing1, Count: 2},
		{Item: thing2, Count: 2},
		{Item: thing3, Count: 2},
	}

	input := []Thing{thing1, thing2, thing3, thing1, thing2, thing3}

	actual := slicer.CountDuplicatesBy(func(elem Thing) string {
		return elem.K
	}, input...)

	assert.ElementsMatch(t, expected, actual)
}

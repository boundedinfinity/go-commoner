package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Uniq(t *testing.T) {
	expectedStr := []string{"a", "b", "c"}
	actualStr := slicer.Uniq("a", "a", "b", "b", "c", "c")
	assert.ElementsMatch(t, expectedStr, actualStr)

	expectedInt := []int{1, 2, 3}
	actualInt := slicer.Uniq(1, 1, 2, 2, 3, 3)
	assert.ElementsMatch(t, expectedInt, actualInt)
}

func Test_UniqFn(t *testing.T) {
	type thing struct {
		S string
		T string
	}

	a := thing{"A", "X"}
	b := thing{"B", "X"}
	c := thing{"C", "X"}

	expectedStr := []thing{a, b, c}
	actualStr := slicer.UniqFn(func(_ int, x thing) string { return x.S }, a, a, b, b, c, c)
	assert.ElementsMatch(t, expectedStr, actualStr)

}

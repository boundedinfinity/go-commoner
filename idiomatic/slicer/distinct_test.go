package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Distinct_String(t *testing.T) {
	expected := []string{"a", "b", "c"}
	input := []string{"a", "a", "b", "b", "c", "c"}
	actual := slicer.Distinct(input)

	assert.ElementsMatch(t, expected, actual)
}

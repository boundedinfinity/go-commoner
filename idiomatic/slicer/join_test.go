package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Join(t *testing.T) {
	expected := "1,2,3"
	input := []int{1, 2, 3}
	actual := slicer.Join(",", input...)

	assert.Equal(t, expected, actual)
}

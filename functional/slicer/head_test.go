package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/functional/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Head(t *testing.T) {
	expected := 1
	actual := slicer.Head(1, 2, 3)

	assert.True(t, actual.Defined())
	assert.Equal(t, expected, actual.Get())
}

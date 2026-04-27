package mather_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/stretchr/testify/assert"
)

func Test_Max(t *testing.T) {
	expected := 5

	input := []int{1, 2, 3, 4, 5}
	actual := mather.Max(input...)

	assert.Equal(t, expected, actual)
}

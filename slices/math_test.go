package slices_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/slices"
	"github.com/stretchr/testify/assert"
)

func Test_Min(t *testing.T) {
	expected := 1

	input := []int{1, 2, 3, 4, 5}
	actual := slices.Min(input)

	assert.Equal(t, expected, actual)
}

func Test_Max(t *testing.T) {
	expected := 5

	input := []int{1, 2, 3, 4, 5}
	actual := slices.Max(input)

	assert.Equal(t, expected, actual)
}

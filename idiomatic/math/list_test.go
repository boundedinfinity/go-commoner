package math_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/math"
	"github.com/stretchr/testify/assert"
)

func Test_Min(t *testing.T) {
	expected := 1

	input := []int{1, 2, 3, 4, 5}
	actual := math.MinOf(input...)

	assert.Equal(t, expected, actual)
}

func Test_Max(t *testing.T) {
	expected := 5

	input := []int{1, 2, 3, 4, 5}
	actual := math.MaxOf(input...)

	assert.Equal(t, expected, actual)
}

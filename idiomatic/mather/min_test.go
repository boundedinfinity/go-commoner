package mather_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/stretchr/testify/assert"
)

func Test_Min(t *testing.T) {
	expected := 1

	input := []int{1, 2, 3, 4, 5}
	actual, _ := mather.MinOk(input...)

	assert.Equal(t, expected, actual)
}

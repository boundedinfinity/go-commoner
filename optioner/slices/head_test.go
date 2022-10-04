package slices_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/optioner/slices"
	"github.com/stretchr/testify/assert"
)

func Test_HeadO(t *testing.T) {
	expected := 1
	input := []int{1, 2, 3}
	actual := slices.Head(input)

	assert.True(t, actual.Defined())
	assert.Equal(t, expected, actual.Get())
}

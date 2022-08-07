package slices_test

import (
	"testing"

	"github.com/boundedinfinity/commons/slices"
	"github.com/stretchr/testify/assert"
)

func Test_Join(t *testing.T) {
	expected := "1,2,3"
	input := []int{1, 2, 3}
	actual := slices.Join(input, ",")

	assert.Equal(t, expected, actual)
}

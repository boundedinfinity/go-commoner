package slices_test

import (
	"testing"

	"github.com/boundedinfinity/commons/slices"
	"github.com/stretchr/testify/assert"
)

func Test_Flatten(t *testing.T) {
	as := []int{1, 2, 3}
	bs := []int{4, 5, 6}
	actual := slices.Flatten(as, bs)
	expected := []int{1, 2, 3, 4, 5, 6}

	assert.Equal(t, expected, actual)
}

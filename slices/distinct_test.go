package slices_test

import (
	"testing"

	"github.com/boundedinfinity/commons/slices"
	"github.com/stretchr/testify/assert"
)

func Test_Distinct_String(t *testing.T) {
	expected := []string{"a", "b", "c"}
	input := []string{"a", "a", "b", "b", "c", "c"}
	actual := slices.Distinct(input)

	assert.ElementsMatch(t, expected, actual)
}

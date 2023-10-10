package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	input := "a b c d"
	expected := "d c b a"
	actual := stringer.Reverse(input)
	assert.Equal(t, expected, actual)
}

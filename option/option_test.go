package option_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/option"
	"github.com/stretchr/testify/assert"
)

func Test_OrFirst(t *testing.T) {
	input := option.Some("s")
	or1 := option.None[string]()
	or2 := option.Some("x")
	or3 := option.Some("y")
	or4 := option.Some("z")

	actual := input.OrFirst(or1, or2, or3, or4)
	expected := option.Some("x")

	assert.Equal(t, expected.Defined(), actual.Defined())
	assert.Equal(t, expected.Get(), actual.Get())
}

func Test_OrLast(t *testing.T) {
	input := option.Some("s")
	or1 := option.None[string]()
	or2 := option.Some("x")
	or3 := option.Some("y")
	or4 := option.Some("z")

	actual := input.OrLast(or1, or2, or3, or4)
	expected := option.Some("z")

	assert.Equal(t, expected.Defined(), actual.Defined())
	assert.Equal(t, expected.Get(), actual.Get())
}

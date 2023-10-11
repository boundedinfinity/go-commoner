package optioner_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_Some_OrFirst(t *testing.T) {
	input := optioner.Some("s")
	or1 := optioner.None[string]()
	or2 := optioner.Some("x")
	or3 := optioner.Some("y")
	or4 := optioner.Some("z")

	actual := optioner.FirstOf(input, or1, or2, or3, or4)
	expected := optioner.Some("s")

	assert.Equal(t, expected.Defined(), actual.Defined())
	assert.Equal(t, expected.Get(), actual.Get())
}

func Test_None_OrFirst(t *testing.T) {
	input := optioner.None[string]()
	or1 := optioner.None[string]()
	or2 := optioner.Some("x")
	or3 := optioner.Some("y")
	or4 := optioner.Some("z")

	actual := optioner.FirstOf(input, or1, or2, or3, or4)
	expected := optioner.Some("x")

	assert.Equal(t, expected.Defined(), actual.Defined())
	assert.Equal(t, expected.Get(), actual.Get())
}

func Test_Some_OrLast(t *testing.T) {
	input := optioner.Some("s")
	or1 := optioner.None[string]()
	or2 := optioner.Some("x")
	or3 := optioner.Some("y")
	or4 := optioner.Some("z")

	actual := optioner.LastOf(input, or1, or2, or3, or4)
	expected := optioner.Some("z")

	assert.Equal(t, expected.Defined(), actual.Defined())
	assert.Equal(t, expected.Get(), actual.Get())
}

func Test_None_OrLast(t *testing.T) {
	input := optioner.Some("s")
	or1 := optioner.None[string]()
	or2 := optioner.None[string]()
	or3 := optioner.None[string]()
	or4 := optioner.None[string]()

	actual := optioner.LastOf(input, or1, or2, or3, or4)
	expected := optioner.Some("s")

	assert.Equal(t, expected.Defined(), actual.Defined())
	assert.Equal(t, expected.Get(), actual.Get())
}

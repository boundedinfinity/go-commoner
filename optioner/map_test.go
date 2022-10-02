package optioner_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_FlatMap_Some(t *testing.T) {
	input := optioner.Some(1)
	fn2 := optioner.ToFlatMapFunc(func(a int) string {
		return fmt.Sprintf("%v", a)
	})

	actual := optioner.FlatMap(input, fn2)
	expected := optioner.Some("1")

	assert.Equal(t, actual.Empty(), expected.Empty())
	assert.Equal(t, actual.Defined(), expected.Defined())
	assert.Equal(t, actual.Get(), expected.Get())
}

func Test_FlatMap_None(t *testing.T) {
	input := optioner.None[int]()
	fn := func(a int) optioner.Option[string] {
		return optioner.Some(fmt.Sprintf("%v", a))
	}
	actual := optioner.FlatMap(input, fn)
	expected := optioner.None[string]()

	assert.Equal(t, actual.Empty(), expected.Empty())
	assert.Equal(t, actual.Defined(), expected.Defined())
}

package option_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/go-commoner/option"
	"github.com/stretchr/testify/assert"
)

func Test_FlatMap_Some(t *testing.T) {
	input := option.Some(1)
	fn2 := option.ToFlatMapFunc(func(a int) string {
		return fmt.Sprintf("%v", a)
	})

	actual := option.FlatMap(input, fn2)
	expected := option.Some("1")

	assert.Equal(t, actual.Empty(), expected.Empty())
	assert.Equal(t, actual.Defined(), expected.Defined())
	assert.Equal(t, actual.Get(), expected.Get())
}

func Test_FlatMap_None(t *testing.T) {
	input := option.None[int]()
	fn := func(a int) option.Option[string] {
		return option.Some(fmt.Sprintf("%v", a))
	}
	actual := option.FlatMap(input, fn)
	expected := option.None[string]()

	assert.Equal(t, actual.Empty(), expected.Empty())
	assert.Equal(t, actual.Defined(), expected.Defined())
}

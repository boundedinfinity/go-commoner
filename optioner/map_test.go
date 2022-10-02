package optioner_test

import (
	"fmt"
	"testing"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_FlatMap(t *testing.T) {
	tcs := []struct {
		name     string
		input    o.Option[int]
		expected o.Option[string]
		err      error
	}{
		{
			name:     "Some",
			input:    o.Some(1),
			expected: o.Some("1"),
		},
		{
			name:     "None",
			input:    o.None[int](),
			expected: o.None[string](),
		},
	}

	fn := func(a int) o.Option[string] {
		return o.Some(fmt.Sprintf("%v", a))
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := o.FlatMap(tc.input, fn)
			assert.Equal(t, actual.Empty(), tc.expected.Empty())
			assert.Equal(t, actual.Defined(), tc.expected.Defined())
			assert.Equal(t, actual.Get(), tc.expected.Get())
		})
	}
}

func Test_FlatMap_Chain(t *testing.T) {
	a := 1

	fn1 := func(a int) o.Option[float32] {
		f := float32(a)
		return o.Some(f)
	}

	fn2 := func(a float32) o.Option[string] {
		return o.Some(fmt.Sprintf("%v", a))
	}

	o1 := o.FlatMap(o.Some(a), fn1)
	o2 := o.FlatMap(o1, fn2)
	fmt.Print(o2)

	// input := option.Some(1)
	// fn2 := option.ToFlatMapFunc(func(a int) string {
	// 	return fmt.Sprintf("%v", a)
	// })

	// actual := option.FlatMap(input, fn2)
	// expected := option.Some("1")

	// assert.Equal(t, actual.Empty(), expected.Empty())
	// assert.Equal(t, actual.Defined(), expected.Defined())
	// assert.Equal(t, actual.Get(), expected.Get())
}

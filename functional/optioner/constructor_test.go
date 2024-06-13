package optioner_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_Option_constructors(t *testing.T) {
	var intVal int = 1

	tcs := []struct {
		name                 string
		actual               optioner.Option[int]
		expectedDefined      bool
		expectedGet          int
		expectedOrElseInput  int
		expectedOrElseOutput int
	}{
		{
			name:                 "None constructor",
			actual:               optioner.None[int](),
			expectedGet:          0,
			expectedDefined:      false,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 2,
		},
		{
			name:                 "Some OfOk with true",
			actual:               optioner.OfOk(1, true),
			expectedGet:          1,
			expectedDefined:      true,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 1,
		},
		{
			name:                 "Some OfOk with false",
			actual:               optioner.OfOk(1, false),
			expectedGet:          0,
			expectedDefined:      false,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 2,
		},
		{
			name:                 "Some OfFn with true return",
			actual:               optioner.OfFn(1, func(i int) bool { return true }),
			expectedGet:          1,
			expectedDefined:      true,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 1,
		},
		{
			name:                 "Some OfFn with false",
			actual:               optioner.OfFn(1, func(i int) bool { return false }),
			expectedGet:          0,
			expectedDefined:      false,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 2,
		},
		{
			name:                 "Some OfPtr with value",
			actual:               optioner.OfPtr(&intVal),
			expectedGet:          intVal,
			expectedDefined:      true,
			expectedOrElseInput:  2,
			expectedOrElseOutput: intVal,
		},
		{
			name:                 "Some OfPtr with nil",
			actual:               optioner.OfPtr[int](nil),
			expectedGet:          0,
			expectedDefined:      false,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 2,
		},
		{
			name:                 "Some constructor",
			actual:               optioner.Some(1),
			expectedGet:          1,
			expectedDefined:      true,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 1,
		},
		{
			name:                 "Some OfZero",
			actual:               optioner.OfZero(0),
			expectedGet:          0,
			expectedDefined:      false,
			expectedOrElseInput:  2,
			expectedOrElseOutput: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Equal(tt, tc.expectedDefined, tc.actual.Defined())
			assert.Equal(tt, !tc.expectedDefined, tc.actual.Empty())
			assert.Equal(tt, tc.expectedGet, tc.actual.Get())
			assert.Equal(tt,
				tc.expectedOrElseOutput,
				tc.actual.OrElse(tc.expectedOrElseInput),
				tc.name,
			)
		})
	}
}

func Test_Some_with_boolean(t *testing.T) {
	actual := optioner.Some(false)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), false)
	assert.Equal(t, actual.OrElse(true), false)
}

func Test_None_with_string(t *testing.T) {
	actual := optioner.None[string]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), "")
	assert.Equal(t, actual.OrElse("x"), "x")
}

func Test_None_with_int(t *testing.T) {
	actual := optioner.None[int]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), 0)
	assert.Equal(t, actual.OrElse(1), 1)
}

func Test_None_with_bool(t *testing.T) {
	actual := optioner.None[bool]()

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), false)
	assert.Equal(t, actual.OrElse(true), true)
}

func Test_Option_with_nil_string(t *testing.T) {
	actual := optioner.OfPtr[string](nil)

	assert.Equal(t, actual.Empty(), true)
	assert.Equal(t, actual.Defined(), false)
	assert.Equal(t, actual.Get(), "")
	assert.Equal(t, actual.OrElse("x"), "x")
}

func Test_Option_with_string(t *testing.T) {
	v := "s"
	actual := optioner.OfPtr(&v)

	assert.Equal(t, actual.Empty(), false)
	assert.Equal(t, actual.Defined(), true)
	assert.Equal(t, actual.Get(), "s")
	assert.Equal(t, actual.OrElse("x"), "s")
}

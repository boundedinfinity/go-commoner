package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Reduce_Int(t *testing.T) {
	expected := 15

	input := []int{1, 2, 3, 4, 5}

	actual := slicer.Reduce(func(_ int, v int, a int) int {
		return a + v
	}, 0, input...)

	assert.Equal(t, expected, actual)
}

func Test_Reduce_Struct(t *testing.T) {
	type Thing struct {
		a string
		b []string
	}

	expected := Thing{
		a: "a_val",
		b: []string{"b3", "b2", "b1"},
	}

	input := []Thing{
		{a: "a_val"},
		{b: []string{"b1"}},
		{b: []string{"b2"}},
		{b: []string{"b3"}},
		{},
	}

	fn := func(_ int, v Thing, a Thing) Thing {
		if v.a != "" {
			a.a = v.a
		}

		if len(v.b) > 0 {
			a.b = append(a.b, v.b...)
		}

		return a
	}

	actual := slicer.Reduce(fn, Thing{b: []string{}}, input...)
	assert.EqualValues(t, expected, actual)
}

package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Reduce_Int(t *testing.T) {
	expected := 15

	input := []int{1, 2, 3, 4, 5}

	actual := slicer.Reduce(0, input, func(v int, a int) int {
		return a + v
	})

	assert.Equal(t, expected, actual)
}

func Test_Reduce_Struct(t *testing.T) {
	type Thing struct {
		a string
		b []string
	}

	expected := Thing{
		a: "a_val",
		b: []string{"b1", "b2", "b3"},
	}

	input := []Thing{
		{a: "a_val"},
		{b: []string{"b1"}},
		{b: []string{"b2"}},
		{b: []string{"b3"}},
		{},
	}

	actual := slicer.Reduce(Thing{b: []string{}}, input, func(v Thing, a Thing) Thing {
		if v.a != "" {
			a.a = v.a
		}

		if len(v.b) > 0 {
			a.b = append(a.b, v.b...)
		}

		return a
	})

	assert.Equal(t, expected, actual)
}

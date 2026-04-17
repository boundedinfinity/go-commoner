package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Each(t *testing.T) {
	type Type1 struct {
		thing string
	}

	testCases := []struct {
		name     string
		input    []Type1
		expected []Type1
		ok       bool
	}{
		{
			name:     "should find",
			input:    []Type1{{thing: "a"}, {thing: "b"}},
			expected: []Type1{{thing: "A"}, {thing: "B"}},
			ok:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := []Type1{}

			slicer.Each(func(t1 Type1) {
				actual = append(actual, Type1{
					thing: stringer.ToCapital(t1.thing),
				})
			}, tc.input...)

			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

// func Test_EachErr_NoErr(t *testing.T) {
// 	type Type1 struct {
// 		thing string
// 	}

// 	type Type2 struct {
// 		thing string
// 	}

// 	expected := []Type2{{thing: "a"}, {thing: "b"}}
// 	input := []Type1{{thing: "a"}, {thing: "b"}}
// 	actual, err := slicer.MapErr(input, func(t1 Type1) (Type2, error) {
// 		return Type2{thing: t1.thing}, nil
// 	})

// 	assert.ElementsMatch(t, expected, actual)
// 	assert.Nil(t, err)
// }

// func Test_EachErr_WithErr(t *testing.T) {
// 	type Type1 struct {
// 		thing string
// 	}

// 	type Type2 struct {
// 		thing string
// 	}

// 	expected := []Type2{}
// 	input := []Type1{{thing: "a"}, {thing: "b"}}
// 	actual, err := slicer.MapErr(input, func(t1 Type1) (Type2, error) {
// 		return Type2{}, errors.New("map error")
// 	})

// 	assert.ElementsMatch(t, expected, actual)
// 	assert.NotNil(t, err)
// }

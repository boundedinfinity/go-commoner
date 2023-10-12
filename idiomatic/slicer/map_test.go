package slicer_test

import (
	"errors"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	type Type1 struct {
		thing string
	}

	type Type2 struct {
		thing string
	}

	expected := []Type2{{thing: "a"}, {thing: "b"}}
	input := []Type1{{thing: "a"}, {thing: "b"}}
	actual := slicer.Map(func(t1 Type1) Type2 {
		return Type2{thing: t1.thing}
	}, input...)

	assert.ElementsMatch(t, expected, actual)
}

func Test_MapErr_NoErr(t *testing.T) {
	type Type1 struct {
		thing string
	}

	type Type2 struct {
		thing string
	}

	expected := []Type2{{thing: "a"}, {thing: "b"}}
	input := []Type1{{thing: "a"}, {thing: "b"}}
	actual, err := slicer.MapErr(func(t1 Type1) (Type2, error) {
		return Type2{thing: t1.thing}, nil
	}, input...)

	assert.ElementsMatch(t, expected, actual)
	assert.Nil(t, err)
}

func Test_MapErr_WithErr(t *testing.T) {
	type Type1 struct {
		thing string
	}

	type Type2 struct {
		thing string
	}

	expected := []Type2{}
	input := []Type1{{thing: "a"}, {thing: "b"}}
	actual, err := slicer.MapErr(func(t1 Type1) (Type2, error) {
		return Type2{}, errors.New("map error")
	}, input...)

	assert.ElementsMatch(t, expected, actual)
	assert.NotNil(t, err)
}

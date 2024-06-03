package slicer_test

import (
	"errors"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	type Thing1 struct {
		aItem string
	}

	type Thing2 struct {
		bItem string
	}

	expected := []Thing2{{bItem: "a"}, {bItem: "b"}}
	input := []Thing1{{aItem: "a"}, {aItem: "b"}}

	actual := slicer.Map(func(t1 Thing1) Thing2 {
		return Thing2{bItem: t1.aItem}
	}, input...)

	assert.ElementsMatch(t, expected, actual)
}

func Test_MapErr_NoErr(t *testing.T) {
	type Thing1 struct {
		aItem string
	}

	type Thing2 struct {
		bItem string
	}

	expected := []Thing2{{bItem: "a"}, {bItem: "b"}}
	input := []Thing1{{aItem: "a"}, {aItem: "b"}}
	actual, err := slicer.MapErr(func(t1 Thing1) (Thing2, error) {
		return Thing2{bItem: t1.aItem}, nil
	}, input...)

	assert.ElementsMatch(t, expected, actual)
	assert.Nil(t, err)
}

func Test_MapErr_WithErr(t *testing.T) {
	type Thing1 struct {
		aItem string
	}

	type Thing2 struct {
		bItem string
	}

	expected := []Thing2{}
	input := []Thing1{{aItem: "a"}, {aItem: "b"}}
	actual, err := slicer.MapErr(func(t1 Thing1) (Thing2, error) {
		return Thing2{}, errors.New("map error")
	}, input...)

	assert.ElementsMatch(t, expected, actual)
	assert.NotNil(t, err)
}

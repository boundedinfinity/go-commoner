package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Dedup(t *testing.T) {
	expected := []int{1, 2, 3}
	input := []int{1, 1, 2, 2, 3, 3}
	actual := slicer.Dedup(input)

	assert.Equal(t, expected, actual)
}

func Test_DedupFn(t *testing.T) {
	type Thing struct {
		K string
		V int
	}

	expected := []Thing{
		{K: "x", V: 100},
		{K: "y", V: 200},
		{K: "z", V: 300},
	}

	input := []Thing{
		{K: "x", V: 100},
		{K: "y", V: 200},
		{K: "z", V: 300},
		{K: "x", V: 100},
		{K: "y", V: 200},
		{K: "z", V: 300},
	}

	actual := slicer.DedupFn(input, func(item Thing) string {
		return item.K
	})

	assert.Equal(t, expected, actual)
}

package pather_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/stretchr/testify/assert"
)

func Test_SwapExt_1(t *testing.T) {
	expected := "/some/path/file.txt"
	input := "/some/path/file.json"

	actual := pather.SwapExt(input, "json", "txt")

	assert.Equal(t, expected, actual)
}

func Test_SwapExt_2(t *testing.T) {
	expected := "/some/path/file.txt"
	input := "/some/path/file.json"

	actual := pather.SwapExt(input, ".json", ".txt")

	assert.Equal(t, expected, actual)
}

func Test_SwapExt_3(t *testing.T) {
	expected := "/some/path/file.txt"
	input := "/some/path/file.txt"

	actual := pather.SwapExt(input, ".json", ".txt")

	assert.Equal(t, expected, actual)
}

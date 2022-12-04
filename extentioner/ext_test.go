package extentioner_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/extentioner"
	"github.com/stretchr/testify/assert"
)

func Test_Ext_1(t *testing.T) {
	input := "/some/path/file.json"
	expected := ".json"
	actual := extentioner.Ext(input)

	assert.Equal(t, expected, actual)
}

func Test_ExtOnly_1(t *testing.T) {
	input := "/some/path/file.json"
	expected := "json"
	actual := extentioner.ExtOnly(input)

	assert.Equal(t, expected, actual)
}

func Test_Normalize_1(t *testing.T) {
	expected := ".txt"
	input := "txt"
	actual := extentioner.Normalize(input)

	assert.Equal(t, expected, actual)
}

func Test_Normalize_2(t *testing.T) {
	expected := ".txt"
	input := ".txt"
	actual := extentioner.Normalize(input)

	assert.Equal(t, expected, actual)
}

func Test_Strip_1(t *testing.T) {
	expected := "/some/path/file"
	input := "/some/path/file.json"
	actual := extentioner.Strip(input)

	assert.Equal(t, expected, actual)
}

func Test_Swap_1(t *testing.T) {
	input := "/some/path/file.json"
	expected := "/some/path/file.txt"
	actual := extentioner.Swap(input, "json", "txt")

	assert.Equal(t, expected, actual)
}

func Test_Swap_2(t *testing.T) {
	input := "/some/path/file.json"
	expected := "/some/path/file.txt"
	actual := extentioner.Swap(input, ".json", ".txt")

	assert.Equal(t, expected, actual)
}

func Test_Swap_3(t *testing.T) {
	input := "/some/path/file.txt"
	expected := "/some/path/file.txt"
	actual := extentioner.Swap(input, ".json", ".txt")

	assert.Equal(t, expected, actual)
}

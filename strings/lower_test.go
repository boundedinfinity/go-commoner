package strings_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/strings"
	"github.com/stretchr/testify/assert"
)

func Test_Lowercase_Normal_String(t *testing.T) {
	input := "StrinG"
	expected := "string"
	actual := strings.Lowercase(input)

	assert.Equal(t, expected, actual)
}

func Test_Lowercase_Typed_String(t *testing.T) {
	type MyString string
	var input MyString = "StrinG"
	var expected string = "string"
	actual := strings.Lowercase(input)

	assert.Equal(t, expected, actual)
}

func Test_LowercaseFirst_Normal_String(t *testing.T) {
	input := "STRING"
	expected := "sTRING"
	actual := strings.LowercaseFirst(input)

	assert.Equal(t, expected, actual)
}

func Test_LowercaseFirst_Typed_String(t *testing.T) {
	type MyString string
	var input MyString = "STRING"
	var expected string = "sTRING"
	actual := strings.LowercaseFirst(input)

	assert.Equal(t, expected, actual)
}

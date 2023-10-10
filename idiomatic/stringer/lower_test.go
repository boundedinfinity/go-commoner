package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Lowercase_Normal_String(t *testing.T) {
	input := "StrinG"
	expected := "string"
	actual := stringer.Lowercase(input)

	assert.Equal(t, expected, actual)
}

func Test_Lowercase_Typed_String(t *testing.T) {
	type MyString string
	var input MyString = "StrinG"
	var expected string = "string"
	actual := stringer.Lowercase(input)

	assert.Equal(t, expected, actual)
}

func Test_LowercaseFirst_Normal_String(t *testing.T) {
	input := "STRING"
	expected := "sTRING"
	actual := stringer.LowercaseFirst(input)

	assert.Equal(t, expected, actual)
}

func Test_LowercaseFirst_Typed_String(t *testing.T) {
	type MyString string
	var input MyString = "STRING"
	var expected string = "sTRING"
	actual := stringer.LowercaseFirst(input)

	assert.Equal(t, expected, actual)
}

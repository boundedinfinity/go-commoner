package strings_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/strings"
	"github.com/stretchr/testify/assert"
)

func Test_Uppercase_Normal_String(t *testing.T) {
	input := "string"
	expected := "STRING"
	actual := strings.Uppercase(input)

	assert.Equal(t, expected, actual)
}

func Test_Uppercase_Typed_String(t *testing.T) {
	type MyString string
	var input MyString = "string"
	var expected string = "STRING"
	actual := strings.Uppercase(input)

	assert.Equal(t, expected, actual)
}

func Test_UppercaseFirst_Normal_String(t *testing.T) {
	input := "string"
	expected := "String"
	actual := strings.UppercaseFirst(input)

	assert.Equal(t, expected, actual)
}

func Test_UppercaseFirst_Typed_String(t *testing.T) {
	type MyString string
	var input MyString = "string"
	var expected string = "String"
	actual := strings.UppercaseFirst(input)

	assert.Equal(t, expected, actual)
}

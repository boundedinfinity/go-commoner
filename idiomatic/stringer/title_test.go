package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_Title_Normal_String(t *testing.T) {
	input := "string"
	expected := "STRING"
	actual := stringer.Uppercase(input)

	assert.Equal(t, expected, actual)
}

func Test_Title_Typed_String(t *testing.T) {
	type MyString string
	var input MyString = "string"
	var expected string = "STRING"
	actual := stringer.Uppercase(input)

	assert.Equal(t, expected, actual)
}

package strings_test

import (
	"testing"

	"github.com/boundedinfinity/commons/strings"
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
	var expected MyString = "STRING"
	actual := strings.Uppercase(input)

	assert.Equal(t, expected, actual)
}

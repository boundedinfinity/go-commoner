package strings_test

import (
	"testing"

	"github.com/boundedinfinity/commons/strings"
	"github.com/stretchr/testify/assert"
)

func Test_String_IsBlank(t *testing.T) {
	assert.Equal(t, strings.IsBlank(""), true)

	type MyString string
	var ms MyString = ""
	assert.Equal(t, strings.IsBlank(ms), true)
}

func Test_String_IsEmpty(t *testing.T) {
	assert.Equal(t, strings.IsEmpty(""), true)

	type MyString string
	var ms MyString = ""
	assert.Equal(t, strings.IsBlank(ms), true)
}

func Test_String_Abbreviate(t *testing.T) {
	input := "a test string that is too long"
	expected := "a test string..."
	actual := strings.Abbreviate(input, 16)
	assert.Equal(t, actual, expected)
}

func Test_String_Abbreviate_shorter(t *testing.T) {
	input := "a test string"
	expected := "a test string"
	actual := strings.Abbreviate(input, 16)
	assert.Equal(t, actual, expected)
}

func Test_String_Abbreviate_equal_length(t *testing.T) {
	input := "a test string"
	expected := "a test ..."
	actual := strings.Abbreviate(input, len(input)-3)
	assert.Equal(t, actual, expected)
}

package strings_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/strings"
	"github.com/stretchr/testify/assert"
)

func Test_String_Trim_one_space(t *testing.T) {
	input := " a string "
	expected := "a string"
	actual := strings.TrimSpace(input)
	assert.Equal(t, expected, actual)
}

func Test_String_Trim_multiple_space(t *testing.T) {
	input := "   a string   "
	expected := "a string"
	actual := strings.TrimSpace(input)
	assert.Equal(t, expected, actual)
}

func Test_String_Trim_other_characters(t *testing.T) {
	input := "!^a string^!"
	expected := "a string"
	actual := strings.Trim(input, "!^")
	assert.Equal(t, expected, actual)
}

func Test_String_TrimFn_one_space(t *testing.T) {
	input := " a string "
	expected := "a string"
	actual := strings.TrimFunc(input, func(x rune) bool { return x == ' ' })
	assert.Equal(t, expected, actual)
}

func Test_String_TrimFn_multiple_space(t *testing.T) {
	input := "   a string   "
	expected := "a string"
	actual := strings.TrimFunc(input, func(x rune) bool { return x == ' ' })
	assert.Equal(t, expected, actual)
}

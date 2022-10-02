package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_String_Trim_one_space(t *testing.T) {
	input := " a string "
	expected := "a string"
	actual := stringer.TrimSpace(input)
	assert.Equal(t, expected, actual)
}

func Test_String_Trim_multiple_space(t *testing.T) {
	input := "   a string   "
	expected := "a string"
	actual := stringer.TrimSpace(input)
	assert.Equal(t, expected, actual)
}

func Test_String_Trim_other_characters(t *testing.T) {
	input := "!^a string^!"
	expected := "a string"
	actual := stringer.Trim(input, "!^")
	assert.Equal(t, expected, actual)
}

func Test_String_TrimFn_one_space(t *testing.T) {
	input := " a string "
	expected := "a string"
	actual := stringer.TrimFunc(input, func(x rune) bool { return x == ' ' })
	assert.Equal(t, expected, actual)
}

func Test_String_TrimFn_multiple_space(t *testing.T) {
	input := "   a string   "
	expected := "a string"
	actual := stringer.TrimFunc(input, func(x rune) bool { return x == ' ' })
	assert.Equal(t, expected, actual)
}

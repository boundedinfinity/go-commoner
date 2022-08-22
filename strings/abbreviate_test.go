package strings_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/strings"
	"github.com/stretchr/testify/assert"
)

func Test_String_Abbreviate(t *testing.T) {
	input := "a test string that is too long"
	expected := "a test string..."
	actual := strings.Abbreviate(input, 16)
	assert.Equal(t, expected, actual)
}

func Test_String_Abbreviate_shorter(t *testing.T) {
	input := "a test string"
	expected := "a test string"
	actual := strings.Abbreviate(input, 16)
	assert.Equal(t, expected, actual)
}

func Test_String_Abbreviate_equal_length(t *testing.T) {
	input := "a test string"
	expected := "a test ..."
	actual := strings.Abbreviate(input, len(input)-3)
	assert.Equal(t, expected, actual)
}

func Test_String_AbbreviateLeft(t *testing.T) {
	input := "a test string that is too long"
	expected := "...t is too long"
	actual := strings.AbbreviateLeft(input, 16)
	assert.Equal(t, expected, actual)
}

package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_String_Split(t *testing.T) {
	input := "a b c d"
	expected := []string{"a", "b", "c", "d"}
	actual := stringer.Split(input, " ")
	assert.Equal(t, expected, actual)
}

func Test_String_Split_empty(t *testing.T) {
	type MyString string
	var input MyString = "a b c d"
	expected := []string{"a", "b", "c", "d"}
	actual := stringer.Split(input, " ")
	assert.Equal(t, expected, actual)
}

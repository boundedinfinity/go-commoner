package golang_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/stringer/golang"
	"github.com/stretchr/testify/assert"
)

func Test_Name_normal(t *testing.T) {
	input := "name"
	expected := "Name"
	actual := golang.Name(input)

	assert.Equal(t, expected, actual)
}

func Test_Name_keyword(t *testing.T) {
	input := "if"
	expected := "_if"
	actual := golang.Name(input)

	assert.Equal(t, expected, actual)
}

func Test_Name_starts_with_number(t *testing.T) {
	input := "9name"
	expected := "_9name"
	actual := golang.Name(input)

	assert.Equal(t, expected, actual)
}

func Test_Name_with_space(t *testing.T) {
	input := "na me"
	expected := "Na_me"
	actual := golang.Name(input)

	assert.Equal(t, expected, actual)
}

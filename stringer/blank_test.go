package stringer_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/stretchr/testify/assert"
)

func Test_String_IsBlank(t *testing.T) {
	assert.Equal(t, stringer.IsBlank(""), true)

	type MyString string
	var ms MyString = ""
	assert.Equal(t, stringer.IsBlank(ms), true)
}

func Test_String_IsEmpty(t *testing.T) {
	assert.Equal(t, stringer.IsEmpty(""), true)

	type MyString string
	var ms MyString = ""
	assert.Equal(t, stringer.IsBlank(ms), true)
}
